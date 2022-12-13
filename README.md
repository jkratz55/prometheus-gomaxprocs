# Prometheus GOMAXPROCS

A simple package that automatically configures and registers prometheus gauges for GOMAXPROCS. The motivation for this library was I had a need to observe and alert on CPU utilization of pods in Kubernetes by utilization in relation to GOMAXPROCS value rather than requested CPU. 

## Usage

To use this library simply import `"github.com/jkratz55/prometheus-gomaxprocs/metrics"` as a side effect. It will automatically register prometheus gauge for GOMAXPROCS value.

```go
package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	_ "github.com/jkratz55/prometheus-gomaxprocs/metrics"
)

func main() {
	server := &http.Server{
		Addr:    ":8082",
		Handler: promhttp.Handler(),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
```
