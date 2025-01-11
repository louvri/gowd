package datadog

import (
	"fmt"
	"github.com/DataDog/datadog-go/v5/statsd"
	"log"
	"os"
	"time"
)

type Client struct {
	Client      *statsd.Client
	Enabled     bool
	ServiceName string
}

// New returns a new Client object using datadog.
// serviceName will be used as prefix at each metric.
func New(host, namespace, serviceName string, port int, enabled bool) *Client {
	if enabled {
		url := fmt.Sprintf("%s:%d", host, port)

		var statsDClient *statsd.Client
		var err error
		if namespace == "" {
			statsDClient, err = statsd.New(url)
		} else {
			statsDClient, err = statsd.New(url, statsd.WithNamespace(namespace))
		}
		if err != nil {
			log.Fatalf("error starting datadog client: %s\n", err)
		}
		return &Client{
			Client:      statsDClient,
			Enabled:     enabled,
			ServiceName: serviceName,
		}
	} else {
		return &Client{}
	}
}

// Default returns a new Client object using datadog with default config.
// serviceName will be used as prefix at each metric.
func Default(namespace, serviceName string) *Client {
	return New(os.Getenv("DD_AGENT_HOST"), namespace, serviceName, 8125, true)
}

// DisableMetric set monitor to be disabled, it will stop sending metrics to datadog
func (c *Client) DisableMetric() {
	c.Enabled = false
}

// EnableMetric set monitor to be enabled, it will start sending metrics to datadog if previously disabled
func (c *Client) EnableMetric() {
	c.Enabled = true
}

// Count send count metric with specified block name, with serviceName as prefix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) Count(blockName string, value int64, tags []string) {
	if c.Enabled {
		_ = c.Client.Count(fmt.Sprintf("%s.%s", c.ServiceName, blockName), value, tags, 1)
	}
}

// CountError send count metric with specified block name, with serviceName as prefix and .error as suffix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) CountError(blockName string, value int64, tags []string) {
	if c.Enabled {
		_ = c.Client.Count(fmt.Sprintf("%s.%s.error", c.ServiceName, blockName), value, tags, 1)
	}
}

// Increment send increment metric with specified block name, with serviceName as prefix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) Increment(blockName string, tags []string) {
	if c.Enabled {
		_ = c.Client.Incr(fmt.Sprintf("%s.%s", c.ServiceName, blockName), tags, 1)
	}
}

// IncrementError send increment metric with specified block name, with serviceName as prefix and .error as suffix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) IncrementError(blockName string, tags []string) {
	if c.Enabled {
		_ = c.Client.Incr(fmt.Sprintf("%s.%s.error", c.ServiceName, blockName), tags, 1)
	}
}

// Decrement send decrement metric with specified metric name, with serviceName as prefix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) Decrement(blockName string, tags []string) {
	if c.Enabled {
		_ = c.Client.Decr(fmt.Sprintf("%s.%s", c.ServiceName, blockName), tags, 1)
	}
}

// DecrementError send decrement metric with specified metric name, with serviceName as prefix and .error as suffix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) DecrementError(blockName string, tags []string) {
	if c.Enabled {
		_ = c.Client.Decr(fmt.Sprintf("%s.%s.error", c.ServiceName, blockName), tags, 1)
	}
}

// Timing send time metric with specified metric name, with serviceName as prefix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) Timing(blockName string, value time.Duration, tags []string) {
	if c.Enabled {
		_ = c.Client.Timing(fmt.Sprintf("%s.%s", c.ServiceName, blockName), value, tags, 1)
	}
}

// TimingError send time metric with specified metric name, with serviceName as prefix and .error as suffix.
// the metric will only be sent if the monitor's enabled.
func (c *Client) TimingError(blockName string, value time.Duration, tags []string) {
	if c.Enabled {
		_ = c.Client.Timing(fmt.Sprintf("%s.%s.error", c.ServiceName, blockName), value, tags, 1)
	}
}
