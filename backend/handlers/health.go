package handlers

import (
	"net/http"
	"time"

	"github.com/nimisean/intelligent-systems-lab/backend/response"
)

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Service   string `json:"service"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	data := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Service:   "intelligent-systems-backend",
	}

	response.JSON(w, http.StatusOK, data)
}
