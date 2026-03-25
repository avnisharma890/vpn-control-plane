package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "vpn_requests_total",
			Help: "Total number of API requests",
		},
		[]string{"method", "endpoint"},
	)

	ErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "vpn_errors_total",
			Help: "Total number of errors",
		},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "vpn_request_duration_seconds",
			Help:    "Request latency",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint"},
	)
)

func Init() {
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(ErrorCounter)
	prometheus.MustRegister(RequestDuration)
}