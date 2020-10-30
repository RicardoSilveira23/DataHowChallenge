package handlers

import (
	"context"
	"log"
	"net/http"

	"../data"
)

// Init
type Logs struct {
	logger *log.Logger
}

func NewLogs(l *log.Logger) *Logs {
	return &Logs{l}
}

// Add New Log
func (ml *Logs) AddMetricsLog(rw http.ResponseWriter, r *http.Request) {
	ml.logger.Println("Handle POST LOG")

	log := r.Context().Value(KeyMetricsLog{}).(data.MetricsLog)

	ml.logger.Printf("LOG: %#v", log)
	data.AddMetricsLog(&log)
}

type KeyMetricsLog struct{}

// Simple Validation on JSON passed
func (p *Logs) MiddlewareMetricsLogsValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.MetricsLog{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.logger.Println("[ERROR] deserializing log", err)
			http.Error(rw, "Error reading log", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyMetricsLog{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)

	})
}
