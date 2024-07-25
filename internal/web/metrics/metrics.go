package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	WebsocketConns prometheus.Gauge
	ReqDuration    *prometheus.HistogramVec // request duration for new game
}

func NewMetrics(req prometheus.Registerer) *Metrics {
	m := &Metrics{
		WebsocketConns: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "websocket_connections",
			Help:      "total number of active websocket connections",
		}),
		ReqDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "app",
			Name:      "request_game_duration",
			Help:      "request duration when creating new game and requesting form",
			Buckets:   prometheus.LinearBuckets(0.05, 0.05, 20),
		}, []string{"method"}),
	}
	req.MustRegister(m.WebsocketConns, m.ReqDuration)
	return m
}
