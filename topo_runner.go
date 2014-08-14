package pickett

import (
	"fmt"
	"time"

	"github.com/igneous-systems/pickett/io"
)

//nodeOrName represents a labelled entity. It can be either a tag that must be in the local
//docker cache or a node that is part of our dependency graph.
type nodeOrName struct {
	name      string
	isNode    bool
	node      node
	instances int
}

//topo runner is single node in a topology
type topoRunner struct {
	n             string
	runIn         nodeOrName
	entry         []string
	consumes      []runner
	policy        policy
	expose        map[io.Port][]io.PortBinding
	containerName string
}

func (n *topoRunner) name() string {
	return n.n
}

func (n *topoRunner) exposed() map[io.Port][]io.PortBinding {
	return n.expose
}

func (n *topoRunner) entryPoint() []string {
	return n.entry
}

//in returns a single node that is our inbound edge, the container we run in.
func (n *topoRunner) in() []node {
	result := []node{}
	if n.runIn.isNode {
		result = append(result, n.runIn.node)
	}
	return result
}

//imageName returns the image name needed to run this network.
func (n *topoRunner) imageName() string {
	return n.runIn.name
}

// run actually does the work to launch this network ,including launching all the networks
// that this one depends on (consumes).  Note that behavior of starting or stopping
// particular dependent services is controllled through the policy apparatus.
func (n *topoRunner) run(teeOutput bool, conf *Config, topoName string, instance int) (*policyInput, error) {

	links := make(map[string]string)
	for _, r := range n.consumes {
		conf.helper.Debug("launching %s because %s consumes it (only launching one instance)", r.name(), n.name())
		input, err := r.run(false, conf, topoName, 0)
		if err != nil {
			return nil, err
		}
		links[input.containerName] = input.r.name()
	}

	in, err := createPolicyInput(n, topoName, instance, conf)
	if err != nil {
		return nil, err
	}
	n.containerName = in.containerName //for use in destroy
	return in, n.policy.appyPolicy(teeOutput, in, topoName, instance, links, conf)
}

// imageIsOutOfDate delegates to the image if it is a node, otherwise false.
func (n *topoRunner) imageIsOutOfDate(conf *Config) (bool, error) {
	if !n.runIn.isNode {
		conf.helper.Debug("'%s' can't be out of date, image '%s' is not buildable\n", n.name(), n.runIn.name)
		return false, nil
	}
	return n.runIn.node.isOutOfDate(conf)
}

// we build the image if indeed that is possible
func (n *topoRunner) imageBuild(conf *Config) error {
	if !n.runIn.isNode {
		fmt.Printf("[pickett WARNING] '%s' can't be built, image '%s' is not buildable\n", n.name(), n.runIn.name)
		return nil
	}
	return n.runIn.node.build(conf)
}

type outcomeProxyBuilder struct {
	net        *topoRunner
	inputName  string
	repository string
	tagname    string
}

func (o *outcomeProxyBuilder) ood(conf *Config) (time.Time, bool, error) {
	ood, err := o.net.imageIsOutOfDate(conf)
	if ood || err != nil {
		return time.Time{}, ood, err
	}

	info, err := conf.cli.InspectImage(o.tag())
	if err != nil {
		//ignoring this because we are assuming it means does not exist
		return time.Time{}, true, nil
	}
	return info.CreatedTime(), false, nil
}

func (o *outcomeProxyBuilder) build(conf *Config) (time.Time, error) {
	err := o.net.imageBuild(conf)
	if err != nil {
		return time.Time{}, err
	}
	conf.helper.Debug("using run node %s to build", o.net.name())

	//this is starting to look dodgier and dodgier
	in, err := o.net.run(true, conf, "pickett-build", 0)
	if err != nil {
		return time.Time{}, err
	}
	_, err = conf.cli.CmdCommit(in.containerName, &io.TagInfo{o.repository, o.tagname})
	if err != nil {
		return time.Time{}, err
	}
	insp, err := conf.cli.InspectImage(o.tag())
	if err != nil {
		return time.Time{}, err
	}
	return insp.CreatedTime(), nil
}

func (o *outcomeProxyBuilder) in() []node {
	result := []node{}
	if o.net.runIn.isNode {
		return append(result, o.net.runIn.node)
	}
	return result
}

func (o *outcomeProxyBuilder) tag() string {
	return o.repository + ":" + o.tagname
}
