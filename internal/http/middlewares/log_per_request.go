package middlewares

import (
	"fmt"
	"net/http"
)

func (m MiddlewaresManager) LogPerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.Logr.Info(fmt.Sprintf("[%s] %s\n", r.Method, r.URL.Path))

		next.ServeHTTP(w, r)
	})
}
