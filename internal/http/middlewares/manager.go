package middlewares

import "github.com/amirhnajafiz/letsgo/internal/metrics"

// MiddlewaresManager contains middleware methods that are being used
// after the http router and before handlers.
type MiddlewaresManager struct {
	Metrics *metrics.Metrics
}
