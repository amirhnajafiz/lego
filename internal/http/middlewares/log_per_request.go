package middlewares

import (
	"context"
	"net/http"
	"time"
)

// contextKey type is used for setting a value on user request context
type contextKey string

func (m MiddlewaresManager) LogPerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), contextKey("timestamp"), time.Now()))

		next.ServeHTTP(w, r)
	})
}
