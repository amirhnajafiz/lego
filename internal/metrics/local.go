package metrics

// localMetrics is a struct that supports Metrics interface. In this system we are
// storing our metrics in local variables. There are other available tools like Prometheus
// or VictoriaMetrics that can manage metrics.
type localMetrics struct {
	requests int
	objects  int
	memory   float64
}

func (m *localMetrics) IncRequests() {
	m.requests++
}

func (m *localMetrics) IncObjects() {
	m.objects++
}

func (m *localMetrics) DecObjects() {
	m.objects--
}

func (m *localMetrics) ObsMemory(size float64) {
	m.memory += size
}

func (m *localMetrics) Export() map[string]interface{} {
	return map[string]interface{}{
		"total_requests_count": m.requests,
		"total_objects_stored": m.objects,
		"total_memory_used":    m.memory,
	}
}
