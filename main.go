package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Logger Configuration
	logger := log.New(os.Stdout, "datahow-api ", log.LstdFlags)

	// Handlers Configuration
	metricsHandler := handlers.NewMetrics(logger)
	logsHandler := handlers.NewLogs(logger)
	jsonResponseHandler := handlers.NewJsonResponse(logger)

	// Gorilla Mux Handler config
	sm := mux.NewRouter()
	sm.Use(jsonResponseHandler.MiddlewareAddJSONHeader)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/allmetrics", metricsHandler.GetMetrics)
	getRouter.HandleFunc("/metrics", metricsHandler.GetUniqueIPs)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/logs", logsHandler.AddMetricsLog)
	postRouter.Use(logsHandler.MiddlewareMetricsLogsValidation)

	// Web Server Config

	// Metrics Port - 9120
	metricsServer := http.Server{
		Addr:         ":9120",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Logs Port - 5000
	logsServer := http.Server{
		Addr:         ":5000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// Startup server
	go func() {
		logger.Println("\nStarting Metrics Server on port " + metricsServer.Addr)

		err := metricsServer.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server : %s\n", err)
			os.Exit(1)
		}
	}()

	logger.Println("\nStarting Logs Server on port " + logsServer.Addr)
	err := logsServer.ListenAndServe()
	if err != nil {
		logger.Printf("Error starting server : %s\n", err)
		os.Exit(1)
	}

	// Process to gracefully shutdown server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	logsServer.Shutdown(tc)
	metricsServer.Shutdown(tc)

}
