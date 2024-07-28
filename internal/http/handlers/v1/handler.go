package responses

import (
	"net/http"

	"github.com/amirhnajafiz/letsgo/internal/metrics"
	"github.com/amirhnajafiz/letsgo/internal/storage"
	"github.com/amirhnajafiz/letsgo/pkg/converter"
)

// Handler is a struct that implements application handlers for user
// requests.
type Handler struct {
	Storage storage.Storage
	Metrics metrics.Metrics
}

func (h Handler) HandleGetRequests(w http.ResponseWriter, r *http.Request) {
	// check if method is GET
	if r.Method != http.MethodGet {
		return
	}

	// get object key from query params
	key := r.URL.Query().Get("key")

	// fetch object from storage
	obj := h.Storage.Receive(key)
	if obj == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	// convert to JSON object
	bytes, err := converter.StructToJSON(obj)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// return the response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func (h Handler) HandlePostRequests(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != http.MethodPost {
		return
	}
}

func (h Handler) HandleDeleteRequests(w http.ResponseWriter, r *http.Request) {
	// check if method is DELETE
	if r.Method != http.MethodDelete {
		return
	}
}
