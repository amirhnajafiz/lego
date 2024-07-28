package metrics

func (m *Metrics) IncRequests() {
	m.requests++
}

func (m *Metrics) IncObjects() {
	m.objects++
}

func (m *Metrics) DecObjects() {
	m.objects--
}

func (m *Metrics) ObsMemory(size float64) {
	m.memory += size
}

func (m *Metrics) Export() map[string]interface{} {
	return map[string]interface{}{
		"total_requests_count": m.requests,
		"total_objects_stored": m.objects,
		"total_memory_used":    m.memory,
	}
}
