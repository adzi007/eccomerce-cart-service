package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HttpDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "route", "status"})

	HttpRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"method", "route", "status"})
)

func RegisterMetrics() {
	prometheus.MustRegister(HttpDuration, HttpRequests)
}
