package main

import (
	"go.opencensus.io/plugin/ochttp"
	"log"
	"net/http"
	"ocExample/cmd/internal/metrics"
	"ocExample/cmd/internal/settings"

	"contrib.go.opencensus.io/exporter/prometheus"
)

func main() {
	appSettings, err := settings.NewSettings()
	if err != nil {
		log.Fatal(err)
	}

	prometheusExporter, err := prometheus.NewExporter(prometheus.Options{})
	if err != nil {
		log.Fatal(err)
	}
	metrics.RegisterExporters(prometheusExporter)
	err = metrics.SetupViews(appSettings.MetricsViewSettings)
	if err != nil {
		log.Fatal(err)
	}

	metricsServer := metrics.Server(appSettings.MetricsSettings.ServerAddr, prometheusExporter)
	go metricsServer.ListenAndServe()

	indexHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("success"))
	})
	server := http.NewServeMux()
	server.Handle("/", ochttp.WithRouteTag(indexHandler, "/"))

	_ = http.ListenAndServe(appSettings.ListenAddr, &ochttp.Handler{Handler: server})
}
