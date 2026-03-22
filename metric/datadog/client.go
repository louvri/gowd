package datadog

import (
	"fmt"
	"os"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
)

type Client struct {
	client      *statsd.Client
	enabled     bool
	serviceName string
}

// New returns a new Client object using datadog.
// serviceName will be used as prefix at each metric.
func New(host, namespace, serviceName string, port int, enabled bool) *Client {
	if !enabled {
		return &Client{}
	}
	addr := fmt.Sprintf("%s:%d", host, port)
	opts := []statsd.Option{}
	if namespace != "" {
		opts = append(opts, statsd.WithNamespace(namespace))
	}
	client, err := statsd.New(addr, opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gowd: error starting datadog client: %s\n", err)
		return &Client{}
	}
	return &Client{
		client:      client,
		enabled:     enabled,
		serviceName: serviceName,
	}
}

// Default returns a new Client object using datadog with default config.
// serviceName will be used as prefix at each metric.
func Default(namespace, serviceName string) *Client {
	return New(os.Getenv("DD_AGENT_HOST"), namespace, serviceName, 8125, true)
}

func (c *Client) metricName(blockName string) string {
	return fmt.Sprintf("%s.%s", c.serviceName, blockName)
}

func (c *Client) errorMetricName(blockName string) string {
	return fmt.Sprintf("%s.%s.error", c.serviceName, blockName)
}

// Close flushes buffered metrics and closes the underlying client.
func (c *Client) Close() error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

// DisableMetric stops sending metrics to datadog.
func (c *Client) DisableMetric() {
	c.enabled = false
}

// EnableMetric resumes sending metrics to datadog.
func (c *Client) EnableMetric() {
	c.enabled = true
}

// Count sends a count metric with serviceName as prefix.
func (c *Client) Count(blockName string, value int64, tags []string) {
	if c.enabled {
		_ = c.client.Count(c.metricName(blockName), value, tags, 1)
	}
}

// CountError sends a count metric with serviceName as prefix and .error suffix.
func (c *Client) CountError(blockName string, value int64, tags []string) {
	if c.enabled {
		_ = c.client.Count(c.errorMetricName(blockName), value, tags, 1)
	}
}

// Increment sends an increment metric with serviceName as prefix.
func (c *Client) Increment(blockName string, tags []string) {
	if c.enabled {
		_ = c.client.Incr(c.metricName(blockName), tags, 1)
	}
}

// IncrementError sends an increment metric with serviceName as prefix and .error suffix.
func (c *Client) IncrementError(blockName string, tags []string) {
	if c.enabled {
		_ = c.client.Incr(c.errorMetricName(blockName), tags, 1)
	}
}

// Decrement sends a decrement metric with serviceName as prefix.
func (c *Client) Decrement(blockName string, tags []string) {
	if c.enabled {
		_ = c.client.Decr(c.metricName(blockName), tags, 1)
	}
}

// DecrementError sends a decrement metric with serviceName as prefix and .error suffix.
func (c *Client) DecrementError(blockName string, tags []string) {
	if c.enabled {
		_ = c.client.Decr(c.errorMetricName(blockName), tags, 1)
	}
}

// Gauge sends a gauge metric with serviceName as prefix.
func (c *Client) Gauge(blockName string, value float64, tags []string) {
	if c.enabled {
		_ = c.client.Gauge(c.metricName(blockName), value, tags, 1)
	}
}

// GaugeError sends a gauge metric with serviceName as prefix and .error suffix.
func (c *Client) GaugeError(blockName string, value float64, tags []string) {
	if c.enabled {
		_ = c.client.Gauge(c.errorMetricName(blockName), value, tags, 1)
	}
}

// Histogram sends a histogram metric with serviceName as prefix.
func (c *Client) Histogram(blockName string, value float64, tags []string) {
	if c.enabled {
		_ = c.client.Histogram(c.metricName(blockName), value, tags, 1)
	}
}

// HistogramError sends a histogram metric with serviceName as prefix and .error suffix.
func (c *Client) HistogramError(blockName string, value float64, tags []string) {
	if c.enabled {
		_ = c.client.Histogram(c.errorMetricName(blockName), value, tags, 1)
	}
}

// Timing sends a timing metric with serviceName as prefix.
func (c *Client) Timing(blockName string, value time.Duration, tags []string) {
	if c.enabled {
		_ = c.client.Timing(c.metricName(blockName), value, tags, 1)
	}
}

// TimingError sends a timing metric with serviceName as prefix and .error suffix.
func (c *Client) TimingError(blockName string, value time.Duration, tags []string) {
	if c.enabled {
		_ = c.client.Timing(c.errorMetricName(blockName), value, tags, 1)
	}
}
