# monitoring metric

Initiate at your code with:
```
package main

import (
    ...
    "github.com/louvri/gowd/metric"
    ...
)

func main() {
	...
	metricClient := metric.Default("my-namespace", "my-service")
	...
}
```
or you can specify the Datadog agent host detail with:
```
    ...
    metricClient := metric.New("localhost", "my-namespace", "my-service", 8125, true)
    ...
```

Send the metrics with:
```
    ...
    metricClient.Increment("data-creation". []string{"source-01"})
    ...
    metricClient.IncrementError("data-creation". []string{"source-01"})
    ...
```
