package controller

import (
	"net/http"
)

// HealthCheck is a Handler function for satisfying Kubernetes Health checks.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
