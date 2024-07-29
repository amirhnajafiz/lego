package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

// requestInfoLog prints the request data and result after the request is handled.
func (m MiddlewaresManager) requestInfoLog(r *http.Request, timestamp time.Time) {
	requestDuration := time.Since(timestamp).Microseconds()

	messages := []string{
		fmt.Sprintf(" [%s]", r.Method),
		fmt.Sprintf(" path: %s", r.URL.Path),
		fmt.Sprintf(" status: %d", r.Response.StatusCode),
		fmt.Sprintf(" duration: %d ms", requestDuration),
	}

	m.Logr.Info(messages...)
}

func (m MiddlewaresManager) LogPerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := time.Now()

		defer m.requestInfoLog(r, ts)

		next.ServeHTTP(w, r)
	})
}
