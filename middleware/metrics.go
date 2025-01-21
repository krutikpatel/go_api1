package middleware

import (
	"api1/metrics"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware(m *metrics.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		method := c.Request.Method

		// Track in-flight requests
		m.HttpRequestsInFlight.WithLabelValues(method, path).Inc()
		defer m.HttpRequestsInFlight.WithLabelValues(method, path).Dec()

		// Process request
		c.Next()

		// Record duration
		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		// Update metrics
		m.HttpRequestsTotal.WithLabelValues(
			method,
			path,
			string(status),
		).Inc()

		m.HttpRequestDuration.WithLabelValues(
			method,
			path,
		).Observe(duration)
	}
}
