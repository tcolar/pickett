package pickett

import (
	"fmt"
	"path/filepath"
	"time"

	docker_utils "github.com/docker/docker/utils"

	"github.com/igneous-systems/pickett/io"
)

// input for a policy to make a decision.  The actual object that could be stopped or started
// if the policy chooses to do so is r.
type policyInput struct {
	hasStarted       bool
	containerName    string
	containerStarted time.Time
	isRunning        bool
	r                runner
}

type stopPolicy int
type startPolicy int

const (
	NEVER stopPolicy = iota
	FRESH
	ALWAYS

	DONT startPolicy = iota
	RESTART
	CONTINUE
)

type policy struct {
	startIfNonExistant bool
	rebuildIfOOD       bool
	start              startPolicy
	stop               stopPolicy
}

//defaultPolicy returns a sensibly initialized policy object.
func defaultPolicy() policy {
	return policy{
		startIfNonExistant: true,
		rebuildIfOOD:       true,
		start:              RESTART,
		stop:               FRESH,
	}
}

//formContainerKey is a helper for forming the keyname in etcd that corresponds
//to a particular network's container.
func formContainerKey(r runner) string {
	return filepath.Join(io.PICKETT_KEYSPACE, CONTAINERS, r.name())
}

//start runs the runner in its policyInput and records the docker container name into etcd.
//note that this is the lowest level code that knows about the options to docker and etcd.
//this code is the actual implementation of start.
func (p *policyInput) start(teeOutput bool, image string, links map[string]string, cli io.DockerCli, etcd io.EtcdClient) error {
	args := []string{}
	if !teeOutput {
		args = append(args, "-d")
	} else {
		args = append(args, "-i", "-t")
	}
	for k, v := range links {
		args = append(args, "--link", fmt.Sprintf("%s:%s", k, v))
	}
	for k, v := range p.r.exposed() {
		args = append(args, "-p", fmt.Sprintf("%s:%d:%d", k, v, v))
	}
	args = append(append(args, image), p.r.entryPoint()...)
	err := cli.CmdRun(teeOutput, args...)
	if err != nil {
		return err
	}
	if !teeOutput {
		id := cli.LastLineOfStdout()
		insp, err := cli.DecodeInspect(id)
		if err != nil {
			return err
		}
		if _, err = etcd.Put(formContainerKey(p.r), insp.ContainerName()); err != nil {
			return err
		}
		p.containerName = insp.ContainerName()
	}
	return nil
}

// stop stops the runner in its policyInput removes the container from etcd.  This is the actual
// implementation of stop.
func (p *policyInput) stop(cli io.DockerCli, etcd io.EtcdClient) error {
	if err := cli.CmdStop(p.containerName); err != nil {
		return err
	}
	if _, err := etcd.Del(formContainerKey(p.r)); err != nil {
		return err
	}
	return nil
}

const (
	CONTAINERS = "containers"
)

func (p stopPolicy) String() string {
	switch p {
	case ALWAYS:
		return "ALWAYS"
	case NEVER:
		return "NEVER"
	case FRESH:
		return "FRESH"
	}
	panic("unknown stop policy")
}

func (p startPolicy) String() string {
	switch p {
	case CONTINUE:
		return "CONTINUE"
	case DONT:
		return "DONT"
	case RESTART:
		return "RESTART"
	}
	panic("unknown start policy")
}

func (p policy) String() string {
	init := ""
	if !p.startIfNonExistant {
		init = "DONT INIT "
	}
	stop := fmt.Sprintf("STOP[%s]", p.stop)
	rebuild := ""
	if !p.rebuildIfOOD {
		rebuild = "DONT REBUILD "
	}
	start := fmt.Sprintf("START[%s]", p.start)
	return fmt.Sprintf("%s%s %s%s", init, stop, rebuild, start)
}

//applyPolicy takes a given policy and starts or stops containers as appropriate. teeOutput is
//really a proxy for "the user requested this be started".
func (p policy) appyPolicy(teeOutput bool, in *policyInput, links map[string]string, conf *Config) error {

	//STEP 0: is image OOD?
	ood, err := in.r.imageIsOutOfDate(conf)
	if err != nil {
		return err
	}

	//STEP1: is existing at all? All codepaths inside this branch return.
	if !in.hasStarted {
		if !p.startIfNonExistant {
			fmt.Printf("[pickett] policy %s is not starting service %s", p, in.r.name())
			return nil
		}
		if p.rebuildIfOOD && ood {
			conf.helper.Debug("policy %s, rebuilding out of date image for '%s'", in.r.name())
			if err := in.r.imageBuild(conf); err != nil {
				return err
			}
		}
		conf.helper.Debug("policy %s, initial start of %s", p, in.r.name())
		return in.start(teeOutput, in.r.imageName(), links, conf.cli, conf.etcd)
	}

	//STEP2: stop?
	if in.isRunning && ood && p.stop == FRESH {
		conf.helper.Debug("policy %s, stopping %s (because its out of date)", p, in.r.name)
		err = in.stop(conf.cli, conf.etcd)
		if err != nil {
			return err
		}
		in.isRunning = false
	} else if in.isRunning && p.stop == ALWAYS {
		conf.helper.Debug("policy %s, stopping %s because policy is ALWAYS stop", p, in.r.name)
		err = in.stop(conf.cli, conf.etcd)
		if err != nil {
			return err
		}
		in.isRunning = false
	}
	//STEP3: start?
	if !in.isRunning {
		if ood && p.rebuildIfOOD {
			conf.helper.Debug("policy %s, rebuilding out of date image for '%s'", in.r.name())
			if err := in.r.imageBuild(conf); err != nil {
				return err
			}
		}

		var img string
		startIt := false
		if p.start == CONTINUE {
			//this is the nasty case, need to commit the container and then continue
			//execution from where it was
			if err := conf.cli.CmdCommit(in.containerName); err != nil {
				return err
			}
			img = conf.cli.LastLineOfStdout()
			conf.helper.Debug("policy %s, continuing %s from image %s", p, in.r.name, img)
			startIt = true
		} else if p.start == RESTART {
			img = in.r.imageName()
			conf.helper.Debug("policy %s, continuing %s restarting from image %s", p, in.r.name, img)
			startIt = true
		}
		if !startIt {
			if err := in.start(teeOutput, img, links, conf.cli, conf.etcd); err != nil {
				return err
			}
		} else {
			conf.helper.Debug("policy %s, not starting %s", p, in.r.name())
		}
	} else if teeOutput {
		fmt.Printf("[pickett] policy %s, ignoring %s which is already running", p, in.r.name)
	}
	return nil
}

//createPolicyInput does the work of interrogating etcd and if necessary docker to figure
//out the state of services.  It returns a policyInput suitable for applying policy to.
func createPolicyInput(r runner, conf *Config) (*policyInput, error) {
	value, present, err := conf.etcd.Get(formContainerKey(r))
	if err != nil {
		return nil, err
	}
	result := &policyInput{
		hasStarted:    present,
		containerName: value,
		r:             r,
	}
	// XXX this logic with the else clauses seems error prone
	if present {
		insp, err := conf.cli.DecodeInspect(value)
		if err != nil {
			status, ok := err.(*docker_utils.StatusError)
			if ok && status.StatusCode == 1 {
				conf.helper.Debug("ignoring docker container %s that is AWOL, probably was manually killed...", value)
				if _, err := conf.etcd.Del(formContainerKey(r)); err != nil {
					return nil, err
				}
				result.isRunning = false
			} else {
				return nil, err
			}
		} else {
			result.isRunning = insp.Running()
			result.containerStarted = insp.CreatedTime()
		}
	}
	return result, nil
}