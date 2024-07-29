package v1

import (
	"fmt"
	"net/http"

	"github.com/amirhnajafiz/letsgo/internal/monitoring/logging"
	"github.com/amirhnajafiz/letsgo/internal/monitoring/metrics"
	"github.com/amirhnajafiz/letsgo/internal/storage"
	"github.com/amirhnajafiz/letsgo/pkg/converter"
	"github.com/amirhnajafiz/letsgo/pkg/measure"
)

// Handler is a struct that implements application handlers for user
// requests.
type Handler struct {
	Storage storage.Storage
	Metrics metrics.Metrics
	Logr    logging.Logr
}

func (h Handler) HandleGetRequests(w http.ResponseWriter, r *http.Request) {
	// check if method is GET
	if r.Method != http.MethodGet {
		h.Logr.Warning(fmt.Sprintf("unsupported method: %s\n", r.Method))

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
	bytes, err := converter.StructToJSON(objectResponse{
		Key:    key,
		Values: obj.(string),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		h.Logr.Error(err, "failed to marshal object to JSON")

		return
	}

	// return the response
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
}

func (h Handler) HandlePostRequests(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != http.MethodPost {
		h.Logr.Warning(fmt.Sprintf("unsupported method: %s\n", r.Method))

		return
	}

	// get object key and value from query params
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")

	// store in storage
	h.Storage.Store(key, value)

	// update metrics
	h.Metrics.IncObjects()
	h.Metrics.ObsMemory(measure.GetObjectSize(key) + measure.GetObjectSize(value))

	// return the response
	w.WriteHeader(http.StatusOK)
}

func (h Handler) HandleDeleteRequests(w http.ResponseWriter, r *http.Request) {
	// check if method is DELETE
	if r.Method != http.MethodDelete {
		h.Logr.Warning(fmt.Sprintf("unsupported method: %s\n", r.Method))

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

	// delete object from the storage
	h.Storage.Delete(key)

	// update metrics
	h.Metrics.DecObjects()
	h.Metrics.ObsMemory(-1 * (measure.GetObjectSize(key) + measure.GetObjectSize(obj)))

	// return the response
	w.WriteHeader(http.StatusOK)
}
