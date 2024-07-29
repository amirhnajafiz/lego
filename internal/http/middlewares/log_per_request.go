package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

// loggingResponseWriter is used to access the request's response status
// after the request is handled.
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (m MiddlewaresManager) LogPerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ts := time.Now()
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		defer func() {
			m.Logr.Info(fmt.Sprintf("[%s] %s status: %d latency: %d ms", r.Method, r.URL.Path, lrw.statusCode, time.Since(ts).Microseconds()))
		}()

		next.ServeHTTP(lrw, r)
	})
}
