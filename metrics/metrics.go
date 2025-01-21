package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

/*
HTTP Metrics:

user_api_http_requests_total - Total number of HTTP requests
user_api_http_request_duration_seconds - Request duration histogram
user_api_http_requests_in_flight - Current number of requests being processed

Business Metrics:

user_api_users_total - Total number of users in the system
user_api_user_operations_total - Total number of user operations (create/update/delete)
*/

type Metrics struct {
	// HTTP metrics
	HttpRequestsTotal    *prometheus.CounterVec
	HttpRequestDuration  *prometheus.HistogramVec
	HttpRequestsInFlight *prometheus.GaugeVec

	// Business metrics
	UsersTotal          prometheus.Gauge
	UserOperationsTotal *prometheus.CounterVec
}

func NewMetrics(namespace string) *Metrics {
	m := &Metrics{}

	// HTTP metrics
	m.HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	m.HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request duration in seconds",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	m.HttpRequestsInFlight = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: namespace,
			Name:      "http_requests_in_flight",
			Help:      "Current number of HTTP requests in flight",
		},
		[]string{"method", "endpoint"},
	)

	// Business metrics
	m.UsersTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      "users_total",
		Help:      "Total number of users in the system",
	})

	m.UserOperationsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "user_operations_total",
			Help:      "Total number of user operations",
		},
		[]string{"operation", "status"},
	)

	return m
}
