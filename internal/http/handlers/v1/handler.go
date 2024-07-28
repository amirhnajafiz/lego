package responses

import (
	"github.com/amirhnajafiz/letsgo/internal/metrics"
	"github.com/amirhnajafiz/letsgo/internal/storage"
)

type Handler struct {
	storage storage.Storage
	metrics *metrics.Metrics
}
