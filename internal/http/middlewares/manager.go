package middlewares

import (
	"github.com/amirhnajafiz/letsgo/internal/monitoring/logging"
	"github.com/amirhnajafiz/letsgo/internal/monitoring/metrics"
)

// MiddlewaresManager contains middleware methods that are being used
// after the http router and before handlers.
type MiddlewaresManager struct {
	Metrics metrics.Metrics
	Logr    logging.Logr
}
