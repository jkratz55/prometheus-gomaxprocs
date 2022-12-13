package metrics

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func init() {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	goMaxProcs := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_maxprocs",
		Help: "GOMAXPROCS value of the Go runtime",
	})
	if err := prometheus.Register(goMaxProcs); err != nil {
		logger.Printf("Failed to register go_maxprocs gauge")
	}
	goMaxProcs.Set(float64(runtime.GOMAXPROCS(0)))

	cpus := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "go_cpus",
		Help: "The logical CPUs available to the Goroutine",
	})
	if err := prometheus.Register(cpus); err != nil {
		logger.Printf("Failed to register go_cpus gauge")
	}
	cpus.Set(float64(runtime.NumCPU()))

	go func() {
		// Because GOMAXPROCS can be changed via calls to runtime.GOMAXPROCS
		// with a value > 0 we need to periodically poll in the event the
		// value was changed programmatically.
		ticker := time.NewTicker(time.Minute * 1)
		for {
			select {
			case <-ticker.C:
				goMaxProcs.Set(float64(runtime.GOMAXPROCS(0)))
			}
		}
	}()
}
