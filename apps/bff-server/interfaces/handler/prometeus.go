package handler

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewPrometeus Prometeus
func NewPrometeus(reg *prometheus.Registry) http.Handler {
	return mergeMetricsHandler(
		promhttp.Handler(),
		promhttp.HandlerFor(reg, promhttp.HandlerOpts{}),
	)
}

func mergeMetricsHandler(handlers ...http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, h := range handlers {
			h.ServeHTTP(w, r)
		}
	})
}
