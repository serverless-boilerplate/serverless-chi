package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/serverless-boilerplate/serverless-chi/app/schema"
)

func ServiceHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schema.ServiceHealthResponse{
		ServerHealth: "Healthy",
	})
}

func ServiceTime(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(schema.ServiceTimeResponse{
		ServerTime: time.Now().Format(time.RFC3339Nano),
	})
}
