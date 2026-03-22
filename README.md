# gowd

A Go library that provides a simple, unified interface for application monitoring and metrics. Currently supports [Datadog](https://www.datadoghq.com/).

## Installation

```
go get github.com/louvri/gowd
```

## Usage

### Initialization

Using default configuration (reads `DD_AGENT_HOST` from environment, port 8125):

```go
import "github.com/louvri/gowd/metric"

client := metric.Default("my-namespace", "my-service")
defer client.Close()
```

Or with custom configuration:

```go
client := metric.New("localhost", "my-namespace", "my-service", 8125, true)
defer client.Close()
```

### Sending Metrics

```go
// Count
client.Count("records-processed", 10, []string{"source:api"})

// Increment / Decrement
client.Increment("request", []string{"endpoint:users"})
client.Decrement("active-connections", []string{"region:us"})

// Gauge
client.Gauge("queue-depth", 42.0, []string{"queue:jobs"})

// Histogram
client.Histogram("payload-size", 1024.5, []string{"type:upload"})

// Timing
client.Timing("query-duration", queryDuration, []string{"db:primary"})
```

Every metric method has an `Error` variant that appends `.error` to the metric name:

```go
client.IncrementError("request", []string{"endpoint:users"})
// sends: my-service.request.error
```

### Enabling / Disabling

Metrics can be toggled at runtime without replacing the client:

```go
client.DisableMetric() // stops sending metrics
client.EnableMetric()  // resumes sending metrics
```

## License

[MIT](LICENSE)
