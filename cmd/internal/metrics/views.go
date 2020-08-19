package metrics

import (
	"log"

	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/plugin/runmetrics"
	"go.opencensus.io/stats/view"
	"ocExample/cmd/internal/settings"
)

func RegisterExporters(exporters ...view.Exporter) {
	for _, exporter := range exporters {
		view.RegisterExporter(exporter)
	}
}

func SetupViews(config settings.MetricsViewSettings) (err error) {
	if config.GRPCClient {
		log.Print("registering grpc client views")
		err = view.Register(ocgrpc.DefaultClientViews...)
	}
	if config.HTTPClient && err == nil {
		log.Print("registering http client views")
		err = view.Register(HTTPClientViews...)
	}
	if config.HTTPServer && err == nil {
		log.Print("registering http server views")
		err = view.Register(HTTPServerViews...)
	}
	if config.RuntimeMetrics && err == nil {
		log.Print("registering runtime metrics")
		err = runmetrics.Enable(runmetrics.RunMetricOptions{
			EnableCPU:    true,
			EnableMemory: true,
		})
	}
	return err
}

var HTTPServerViews = []*view.View{
	ochttp.ServerRequestCountView,
	ochttp.ServerRequestBytesView,
	ochttp.ServerResponseBytesView,
	ochttp.ServerLatencyView,
	ochttp.ServerRequestCountByMethod,
	ochttp.ServerResponseCountByStatusCode,
}

var HTTPClientViews = []*view.View{
	ochttp.ClientCompletedCount,
	ochttp.ClientSentBytesDistribution,
	ochttp.ClientReceivedBytesDistribution,
	ochttp.ClientRoundtripLatencyDistribution,
}