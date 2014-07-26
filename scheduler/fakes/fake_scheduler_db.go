// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/builds"
	"github.com/concourse/atc/config"
	"github.com/concourse/atc/scheduler"
)

type FakeSchedulerDB struct {
	CreateBuildWithInputsStub        func(job string, inputs builds.VersionedResources) (builds.Build, error)
	createBuildWithInputsMutex       sync.RWMutex
	createBuildWithInputsArgsForCall []struct {
		job    string
		inputs builds.VersionedResources
	}
	createBuildWithInputsReturns struct {
		result1 builds.Build
		result2 error
	}
	GetLatestInputVersionsStub        func([]config.Input) (builds.VersionedResources, error)
	getLatestInputVersionsMutex       sync.RWMutex
	getLatestInputVersionsArgsForCall []struct {
		arg1 []config.Input
	}
	getLatestInputVersionsReturns struct {
		result1 builds.VersionedResources
		result2 error
	}
	GetBuildForInputsStub        func(job string, inputs builds.VersionedResources) (builds.Build, error)
	getBuildForInputsMutex       sync.RWMutex
	getBuildForInputsArgsForCall []struct {
		job    string
		inputs builds.VersionedResources
	}
	getBuildForInputsReturns struct {
		result1 builds.Build
		result2 error
	}
	GetNextPendingBuildStub        func(job string) (builds.Build, builds.VersionedResources, error)
	getNextPendingBuildMutex       sync.RWMutex
	getNextPendingBuildArgsForCall []struct {
		job string
	}
	getNextPendingBuildReturns struct {
		result1 builds.Build
		result2 builds.VersionedResources
		result3 error
	}
}

func (fake *FakeSchedulerDB) CreateBuildWithInputs(job string, inputs builds.VersionedResources) (builds.Build, error) {
	fake.createBuildWithInputsMutex.Lock()
	defer fake.createBuildWithInputsMutex.Unlock()
	fake.createBuildWithInputsArgsForCall = append(fake.createBuildWithInputsArgsForCall, struct {
		job    string
		inputs builds.VersionedResources
	}{job, inputs})
	if fake.CreateBuildWithInputsStub != nil {
		return fake.CreateBuildWithInputsStub(job, inputs)
	} else {
		return fake.createBuildWithInputsReturns.result1, fake.createBuildWithInputsReturns.result2
	}
}

func (fake *FakeSchedulerDB) CreateBuildWithInputsCallCount() int {
	fake.createBuildWithInputsMutex.RLock()
	defer fake.createBuildWithInputsMutex.RUnlock()
	return len(fake.createBuildWithInputsArgsForCall)
}

func (fake *FakeSchedulerDB) CreateBuildWithInputsArgsForCall(i int) (string, builds.VersionedResources) {
	fake.createBuildWithInputsMutex.RLock()
	defer fake.createBuildWithInputsMutex.RUnlock()
	return fake.createBuildWithInputsArgsForCall[i].job, fake.createBuildWithInputsArgsForCall[i].inputs
}

func (fake *FakeSchedulerDB) CreateBuildWithInputsReturns(result1 builds.Build, result2 error) {
	fake.CreateBuildWithInputsStub = nil
	fake.createBuildWithInputsReturns = struct {
		result1 builds.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetLatestInputVersions(arg1 []config.Input) (builds.VersionedResources, error) {
	fake.getLatestInputVersionsMutex.Lock()
	defer fake.getLatestInputVersionsMutex.Unlock()
	fake.getLatestInputVersionsArgsForCall = append(fake.getLatestInputVersionsArgsForCall, struct {
		arg1 []config.Input
	}{arg1})
	if fake.GetLatestInputVersionsStub != nil {
		return fake.GetLatestInputVersionsStub(arg1)
	} else {
		return fake.getLatestInputVersionsReturns.result1, fake.getLatestInputVersionsReturns.result2
	}
}

func (fake *FakeSchedulerDB) GetLatestInputVersionsCallCount() int {
	fake.getLatestInputVersionsMutex.RLock()
	defer fake.getLatestInputVersionsMutex.RUnlock()
	return len(fake.getLatestInputVersionsArgsForCall)
}

func (fake *FakeSchedulerDB) GetLatestInputVersionsArgsForCall(i int) []config.Input {
	fake.getLatestInputVersionsMutex.RLock()
	defer fake.getLatestInputVersionsMutex.RUnlock()
	return fake.getLatestInputVersionsArgsForCall[i].arg1
}

func (fake *FakeSchedulerDB) GetLatestInputVersionsReturns(result1 builds.VersionedResources, result2 error) {
	fake.GetLatestInputVersionsStub = nil
	fake.getLatestInputVersionsReturns = struct {
		result1 builds.VersionedResources
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetBuildForInputs(job string, inputs builds.VersionedResources) (builds.Build, error) {
	fake.getBuildForInputsMutex.Lock()
	defer fake.getBuildForInputsMutex.Unlock()
	fake.getBuildForInputsArgsForCall = append(fake.getBuildForInputsArgsForCall, struct {
		job    string
		inputs builds.VersionedResources
	}{job, inputs})
	if fake.GetBuildForInputsStub != nil {
		return fake.GetBuildForInputsStub(job, inputs)
	} else {
		return fake.getBuildForInputsReturns.result1, fake.getBuildForInputsReturns.result2
	}
}

func (fake *FakeSchedulerDB) GetBuildForInputsCallCount() int {
	fake.getBuildForInputsMutex.RLock()
	defer fake.getBuildForInputsMutex.RUnlock()
	return len(fake.getBuildForInputsArgsForCall)
}

func (fake *FakeSchedulerDB) GetBuildForInputsArgsForCall(i int) (string, builds.VersionedResources) {
	fake.getBuildForInputsMutex.RLock()
	defer fake.getBuildForInputsMutex.RUnlock()
	return fake.getBuildForInputsArgsForCall[i].job, fake.getBuildForInputsArgsForCall[i].inputs
}

func (fake *FakeSchedulerDB) GetBuildForInputsReturns(result1 builds.Build, result2 error) {
	fake.GetBuildForInputsStub = nil
	fake.getBuildForInputsReturns = struct {
		result1 builds.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeSchedulerDB) GetNextPendingBuild(job string) (builds.Build, builds.VersionedResources, error) {
	fake.getNextPendingBuildMutex.Lock()
	defer fake.getNextPendingBuildMutex.Unlock()
	fake.getNextPendingBuildArgsForCall = append(fake.getNextPendingBuildArgsForCall, struct {
		job string
	}{job})
	if fake.GetNextPendingBuildStub != nil {
		return fake.GetNextPendingBuildStub(job)
	} else {
		return fake.getNextPendingBuildReturns.result1, fake.getNextPendingBuildReturns.result2, fake.getNextPendingBuildReturns.result3
	}
}

func (fake *FakeSchedulerDB) GetNextPendingBuildCallCount() int {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return len(fake.getNextPendingBuildArgsForCall)
}

func (fake *FakeSchedulerDB) GetNextPendingBuildArgsForCall(i int) string {
	fake.getNextPendingBuildMutex.RLock()
	defer fake.getNextPendingBuildMutex.RUnlock()
	return fake.getNextPendingBuildArgsForCall[i].job
}

func (fake *FakeSchedulerDB) GetNextPendingBuildReturns(result1 builds.Build, result2 builds.VersionedResources, result3 error) {
	fake.GetNextPendingBuildStub = nil
	fake.getNextPendingBuildReturns = struct {
		result1 builds.Build
		result2 builds.VersionedResources
		result3 error
	}{result1, result2, result3}
}

var _ scheduler.SchedulerDB = new(FakeSchedulerDB)
