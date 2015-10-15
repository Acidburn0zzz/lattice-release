// This file was generated by counterfeiter
package fake_noaa_consumer

import (
	"sync"

	"github.com/cloudfoundry-incubator/lattice/ltc/app_examiner"
	"github.com/cloudfoundry/sonde-go/events"
)

type FakeNoaaConsumer struct {
	GetContainerMetricsStub        func(string, string) ([]*events.ContainerMetric, error)
	getContainerMetricsMutex       sync.RWMutex
	getContainerMetricsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getContainerMetricsReturns struct {
		result1 []*events.ContainerMetric
		result2 error
	}
}

func (fake *FakeNoaaConsumer) GetContainerMetrics(arg1 string, arg2 string) ([]*events.ContainerMetric, error) {
	fake.getContainerMetricsMutex.Lock()
	fake.getContainerMetricsArgsForCall = append(fake.getContainerMetricsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.getContainerMetricsMutex.Unlock()
	if fake.GetContainerMetricsStub != nil {
		return fake.GetContainerMetricsStub(arg1, arg2)
	} else {
		return fake.getContainerMetricsReturns.result1, fake.getContainerMetricsReturns.result2
	}
}

func (fake *FakeNoaaConsumer) GetContainerMetricsCallCount() int {
	fake.getContainerMetricsMutex.RLock()
	defer fake.getContainerMetricsMutex.RUnlock()
	return len(fake.getContainerMetricsArgsForCall)
}

func (fake *FakeNoaaConsumer) GetContainerMetricsArgsForCall(i int) (string, string) {
	fake.getContainerMetricsMutex.RLock()
	defer fake.getContainerMetricsMutex.RUnlock()
	return fake.getContainerMetricsArgsForCall[i].arg1, fake.getContainerMetricsArgsForCall[i].arg2
}

func (fake *FakeNoaaConsumer) GetContainerMetricsReturns(result1 []*events.ContainerMetric, result2 error) {
	fake.GetContainerMetricsStub = nil
	fake.getContainerMetricsReturns = struct {
		result1 []*events.ContainerMetric
		result2 error
	}{result1, result2}
}

var _ app_examiner.NoaaConsumer = new(FakeNoaaConsumer)