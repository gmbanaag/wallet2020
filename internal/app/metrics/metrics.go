package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics handles prometheus metrics
type Metrics struct {
}

// PrometheusHandler returns the prometheus handler
func (m *Metrics) PrometheusHandler() http.Handler {
	return promhttp.Handler()
}
