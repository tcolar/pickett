// Automatically generated by MockGen. DO NOT EDIT!
// Source: io/io.go

package io

import (
	io "io"
	time "time"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of IOHelper interface
type MockIOHelper struct {
	ctrl     *gomock.Controller
	recorder *_MockIOHelperRecorder
}

// Recorder for MockIOHelper (not exported)
type _MockIOHelperRecorder struct {
	mock *MockIOHelper
}

func NewMockIOHelper(ctrl *gomock.Controller) *MockIOHelper {
	mock := &MockIOHelper{ctrl: ctrl}
	mock.recorder = &_MockIOHelperRecorder{mock}
	return mock
}

func (_m *MockIOHelper) EXPECT() *_MockIOHelperRecorder {
	return _m.recorder
}

func (_m *MockIOHelper) Debug(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Debug", _s...)
}

func (_mr *_MockIOHelperRecorder) Debug(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Debug", _s...)
}

func (_m *MockIOHelper) OpenDockerfileRelative(dir string) (io.Reader, error) {
	ret := _m.ctrl.Call(_m, "OpenDockerfileRelative", dir)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIOHelperRecorder) OpenDockerfileRelative(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenDockerfileRelative", arg0)
}

func (_m *MockIOHelper) DirectoryRelative(dir string) string {
	ret := _m.ctrl.Call(_m, "DirectoryRelative", dir)
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIOHelperRecorder) DirectoryRelative(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DirectoryRelative", arg0)
}

func (_m *MockIOHelper) Fatalf(_param0 string, _param1 ...interface{}) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Fatalf", _s...)
}

func (_mr *_MockIOHelperRecorder) Fatalf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Fatalf", _s...)
}

func (_m *MockIOHelper) CheckFatal(_param0 error, _param1 string, _param2 ...interface{}) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "CheckFatal", _s...)
}

func (_mr *_MockIOHelperRecorder) CheckFatal(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CheckFatal", _s...)
}

func (_m *MockIOHelper) ConfigReader() io.Reader {
	ret := _m.ctrl.Call(_m, "ConfigReader")
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

func (_mr *_MockIOHelperRecorder) ConfigReader() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigReader")
}

func (_m *MockIOHelper) ConfigFile() string {
	ret := _m.ctrl.Call(_m, "ConfigFile")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIOHelperRecorder) ConfigFile() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigFile")
}

func (_m *MockIOHelper) LastTimeInDirRelative(dir string) (time.Time, error) {
	ret := _m.ctrl.Call(_m, "LastTimeInDirRelative", dir)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIOHelperRecorder) LastTimeInDirRelative(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LastTimeInDirRelative", arg0)
}

// Mock of DockerCli interface
type MockDockerCli struct {
	ctrl     *gomock.Controller
	recorder *_MockDockerCliRecorder
}

// Recorder for MockDockerCli (not exported)
type _MockDockerCliRecorder struct {
	mock *MockDockerCli
}

func NewMockDockerCli(ctrl *gomock.Controller) *MockDockerCli {
	mock := &MockDockerCli{ctrl: ctrl}
	mock.recorder = &_MockDockerCliRecorder{mock}
	return mock
}

func (_m *MockDockerCli) EXPECT() *_MockDockerCliRecorder {
	return _m.recorder
}

func (_m *MockDockerCli) CmdRun(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdRun", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdRun(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdRun", arg0...)
}

func (_m *MockDockerCli) CmdPs(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdPs", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdPs(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdPs", arg0...)
}

func (_m *MockDockerCli) CmdTag(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdTag", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdTag(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdTag", arg0...)
}

func (_m *MockDockerCli) CmdCommit(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdCommit", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdCommit(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdCommit", arg0...)
}

func (_m *MockDockerCli) CmdInspect(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdInspect", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdInspect(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdInspect", arg0...)
}

func (_m *MockDockerCli) CmdBuild(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdBuild", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdBuild(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdBuild", arg0...)
}

func (_m *MockDockerCli) CmdCp(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdCp", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdCp(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdCp", arg0...)
}

func (_m *MockDockerCli) CmdWait(_param0 ...string) error {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CmdWait", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockDockerCliRecorder) CmdWait(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CmdWait", arg0...)
}

func (_m *MockDockerCli) Stdout() string {
	ret := _m.ctrl.Call(_m, "Stdout")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDockerCliRecorder) Stdout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stdout")
}

func (_m *MockDockerCli) LastLineOfStdout() string {
	ret := _m.ctrl.Call(_m, "LastLineOfStdout")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDockerCliRecorder) LastLineOfStdout() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LastLineOfStdout")
}

func (_m *MockDockerCli) Stderr() string {
	ret := _m.ctrl.Call(_m, "Stderr")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDockerCliRecorder) Stderr() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stderr")
}

func (_m *MockDockerCli) EmptyOutput() bool {
	ret := _m.ctrl.Call(_m, "EmptyOutput")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockDockerCliRecorder) EmptyOutput() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EmptyOutput")
}

func (_m *MockDockerCli) DecodeInspect(_param0 ...string) (Inspected, error) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DecodeInspect", _s...)
	ret0, _ := ret[0].(Inspected)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockDockerCliRecorder) DecodeInspect(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeInspect", arg0...)
}

func (_m *MockDockerCli) DumpErrOutput() {
	_m.ctrl.Call(_m, "DumpErrOutput")
}

func (_mr *_MockDockerCliRecorder) DumpErrOutput() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DumpErrOutput")
}

// Mock of Inspected interface
type MockInspected struct {
	ctrl     *gomock.Controller
	recorder *_MockInspectedRecorder
}

// Recorder for MockInspected (not exported)
type _MockInspectedRecorder struct {
	mock *MockInspected
}

func NewMockInspected(ctrl *gomock.Controller) *MockInspected {
	mock := &MockInspected{ctrl: ctrl}
	mock.recorder = &_MockInspectedRecorder{mock}
	return mock
}

func (_m *MockInspected) EXPECT() *_MockInspectedRecorder {
	return _m.recorder
}

func (_m *MockInspected) CreatedTime() time.Time {
	ret := _m.ctrl.Call(_m, "CreatedTime")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

func (_mr *_MockInspectedRecorder) CreatedTime() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreatedTime")
}
