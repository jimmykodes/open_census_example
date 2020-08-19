package settings

import "github.com/netflix/go-env"

func NewSettings() (*Settings, error) {
	settings := &Settings{}
	_, err := env.UnmarshalFromEnviron(settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

type MetricsViewSettings struct {
	GRPCClient     bool `env:"METRICS_VIEW_GRPC_CLIENT"`
	HTTPClient     bool `env:"METRICS_VIEW_HTTP_CLIENT"`
	HTTPServer     bool `env:"METRICS_VIEW_HTTP_SERVER"`
	RuntimeMetrics bool `env:"METRICS_VIEW_RUNTIME_METRICS"`
}

type MetricsSettings struct {
	ServerAddr          string `env:"METRICS_SERVER_ADDR"`
}

type Settings struct {
	ListenAddr string `env:"LISTEN_ADDR,default=:80"`
	MetricsViewSettings
	MetricsSettings
}
