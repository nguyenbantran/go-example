package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {

	// prometheus exporter
	if false {
		startPrometheusServer()
	}

	// expvar exporter
	if false {
		monitoringSomeExpvar()
	}

	// telemetry
	if true {
		startTelemetry()
	}
}

func startPrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(":8888", nil)
	}()
}
