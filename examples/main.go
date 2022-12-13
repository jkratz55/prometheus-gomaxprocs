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
