# metric

The `metric` package provides a provider-agnostic interface for sending application metrics.

## Initialization

```go
package main

import "github.com/louvri/gowd/metric"

func main() {
    // Default: reads DD_AGENT_HOST from env, uses port 8125
    client := metric.Default("my-namespace", "my-service")
    defer client.Close()

    // Or with custom host and port
    client := metric.New("localhost", "my-namespace", "my-service", 8125, true)
    defer client.Close()
}
```

## Supported Metric Types

| Method | Description |
|---|---|
| `Count` / `CountError` | Absolute count values |
| `Increment` / `IncrementError` | Increment by 1 |
| `Decrement` / `DecrementError` | Decrement by 1 |
| `Gauge` / `GaugeError` | Point-in-time values |
| `Histogram` / `HistogramError` | Value distributions |
| `Timing` / `TimingError` | Duration measurements |

## Example

```go
client.Increment("data-creation", []string{"source:s3"})
client.IncrementError("data-creation", []string{"source:s3"})
```

All metrics are prefixed with the `serviceName` provided during initialization. `Error` variants additionally append `.error` to the metric name.
