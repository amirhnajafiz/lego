package http

import (
	"net/http"

	v0 "github.com/amirhnajafiz/letsgo/internal/http/handlers/v0"
	v1 "github.com/amirhnajafiz/letsgo/internal/http/handlers/v1"
	"github.com/amirhnajafiz/letsgo/internal/http/middlewares"
	"github.com/amirhnajafiz/letsgo/internal/monitoring/logging"
	"github.com/amirhnajafiz/letsgo/internal/monitoring/metrics"
	"github.com/amirhnajafiz/letsgo/internal/storage"
)

// Bootstrap function returns a http server instance
// with all requirements to be executed in the main process.
func Bootstrap() *http.ServeMux {
	// create a new net/http router
	router := http.NewServeMux()

	// create system's requirements
	metricsInstance := metrics.NewMetrics()
	storageInstance := storage.NewStorage()
	logrInstance := logging.NewLogr()

	// create middleware and handlers instances
	middleware := middlewares.MiddlewaresManager{
		Metrics: metricsInstance,
		Logr:    logrInstance,
	}
	v0Handler := v0.Handler{
		Metrics: metricsInstance,
	}
	v1Handler := v1.Handler{
		Metrics: metricsInstance,
		Storage: storageInstance,
		Logr:    logrInstance,
	}

	// setup router endpoints
	router.HandleFunc("/healthz", v0Handler.ReturnHealthStatus)
	router.HandleFunc("/metrics", v0Handler.ReturnMetrics)

	router.Handle("/v1/get", middleware.LogPerRequest(middleware.CountRequestsMiddleware(http.HandlerFunc(v1Handler.HandleGetRequests))))
	router.Handle("/v1/new", middleware.LogPerRequest(middleware.CountRequestsMiddleware(http.HandlerFunc(v1Handler.HandlePostRequests))))
	router.Handle("/v1/del", middleware.LogPerRequest(middleware.CountRequestsMiddleware(http.HandlerFunc(v1Handler.HandleDeleteRequests))))

	return router
}
