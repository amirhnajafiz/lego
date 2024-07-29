package v1

import (
	"fmt"
	"net/http"
	"time"
)

// requestInfoLog prints the request data and result after the request is
// handled.
func (h Handler) requestInfoLog(r *http.Request) {
	requestDuration := time.Since(r.Context().Value("timestamp").(time.Time)).Microseconds()

	messages := []string{
		fmt.Sprintf(" [%s]", r.Method),
		fmt.Sprintf(" path: %s", r.URL.Path),
		fmt.Sprintf(" status: %d", r.Response.StatusCode),
		fmt.Sprintf(" duration: %d ms", requestDuration),
	}

	h.Logr.Info(messages...)
}
