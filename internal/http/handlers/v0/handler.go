package v0

import (
	"net/http"

	"github.com/amirhnajafiz/letsgo/internal/monitoring/metrics"
	"github.com/amirhnajafiz/letsgo/pkg/converter"
)

// Handler is a struct that implements application handlers for getting
// systems status.
type Handler struct {
	Metrics metrics.Metrics
}

func (h Handler) ReturnHealthStatus(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h Handler) ReturnMetrics(w http.ResponseWriter, _ *http.Request) {
	// get metrics
	exportedMetrics := h.Metrics.Export()

	// convert to JSON object
	bytes, err := converter.StructToJSON(exportedMetrics)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// return the response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}
