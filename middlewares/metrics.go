package middlewares

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	MetricsUrl string = "/metrics"
)

func MetricsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == MetricsUrl {
			promhttp.Handler().ServeHTTP(w, req)
			return
		}
		next.ServeHTTP(w, req)
	})
}
