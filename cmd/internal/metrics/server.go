package metrics

import (
	"net/http"

	"contrib.go.opencensus.io/exporter/prometheus"
)

func Server(addr string, exporter *prometheus.Exporter) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/metrics", exporter)
	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
