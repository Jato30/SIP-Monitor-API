package api

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Metrics struct {
	mu     sync.Mutex
	Counts map[string]int `json:"counts"`
}

func NewMetrics() *Metrics {
	return &Metrics{
		Counts: make(map[string]int),
	}
}

// Increment counter for a SIP method, thread-safe
func (m *Metrics) Increment(method string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Counts[method]++
}

// HTTP handler for /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// HTTP handler for /metrics
func (m *Metrics) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	m.mu.Lock()
	defer m.mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(m.Counts)
}
