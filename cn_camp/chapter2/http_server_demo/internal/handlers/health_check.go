package handlers

import (
	"http_server_demo/internal/logger"
	"net/http"
)

func HealthZHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	logger.RecordHttpClient(r, "healthz")
}
