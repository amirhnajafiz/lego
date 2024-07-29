package metrics

// Metrics is a interface that shows required system's metrics.
type Metrics interface {
	IncRequests()
	IncObjects()
	DecObjects()
	ObsMemory(size float64)
	Export() map[string]interface{}
}

// NewMetrics returns a struct that supports Metrics interface, in our case
// it returns a localMetrics instance.
func NewMetrics() Metrics {
	return &localMetrics{
		requests: 0,
		objects:  0,
		memory:   0,
	}
}
