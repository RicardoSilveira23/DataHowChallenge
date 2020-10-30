package handlers

import (
	"log"
	"net/http"

	"../data"
)

// Init
type Metrics struct {
	logger *log.Logger
}

func NewMetrics(l *log.Logger) *Metrics {
	return &Metrics{l}
}

// Get All Metrics
func (m *Metrics) GetMetrics(rw http.ResponseWriter, r *http.Request) {
	m.logger.Println("Handle GET Metrics")

	metrics := data.GetMetrics()
	err := metrics.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode to json", http.StatusInternalServerError)
	}

	return
}
