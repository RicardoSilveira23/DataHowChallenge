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

// Get Unique ID's
func (m *Metrics) GetUniqueIPs(rw http.ResponseWriter, r *http.Request) {
	m.logger.Println("Handle GET Unique IPs")

	uniqueMetrics := data.GetUniqueIPAdrresses()
	err := uniqueMetrics.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to encode to json", http.StatusInternalServerError)
	}

	return
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
