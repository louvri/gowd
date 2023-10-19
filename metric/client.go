package metric

import (
	"github.com/louvri/gowd/metric/datadog"
	"time"
)

type ClientInterface interface {
	DisableMetric()
	EnableMetric()
	Count(blockName string, value int64, tags []string)
	CountError(blockName string, value int64, tags []string)
	Increment(blockName string, tags []string)
	IncrementError(blockName string, tags []string)
	Decrement(blockName string, tags []string)
	DecrementError(blockName string, tags []string)
	Timing(blockName string, value time.Duration, tags []string)
	TimingError(blockName string, value time.Duration, tags []string)
}

// New returns a new Client object.
// serviceName will be used as prefix at each metric.
func New(host, namespace, serviceName string, port int, enabled bool) ClientInterface {
	return datadog.New(host, namespace, serviceName, port, enabled)
}

// Default returns a new Client object with default config.
// serviceName will be used as prefix at each metric.
func Default(namespace, serviceName string) ClientInterface {
	return datadog.Default(namespace, serviceName)
}
