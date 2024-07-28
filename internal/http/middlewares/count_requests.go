package middlewares

import (
	"net/http"
)

func (m MiddlewaresManager) CountRequestsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Metrics.IncRequests()

		next.ServeHTTP(w, r)
	})
}
