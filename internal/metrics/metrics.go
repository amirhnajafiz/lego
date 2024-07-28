package metrics

// Metrics is a struct that holds system metrics.
type Metrics struct {
	requests int
	objects  int
	memory   float64
}

func NewMetrics() *Metrics {
	return &Metrics{
		requests: 0,
		objects:  0,
		memory:   0,
	}
}
