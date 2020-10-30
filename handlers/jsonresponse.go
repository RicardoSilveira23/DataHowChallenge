package handlers

import (
	"log"
	"net/http"
)

type JsonResponse struct {
	logger *log.Logger
}

func NewJsonResponse(l *log.Logger) *JsonResponse {
	return &JsonResponse{l}
}

// Adds JSON response Headers
func (jp *JsonResponse) MiddlewareAddJSONHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(rw, r)

	})
}
