package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getMetrics(r *mux.Router) {
	r.HandleFunc("/api/metrics", func(w http.ResponseWriter, r *http.Request) {
		// Send a response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Metrics endpoint"})
	}).Methods("GET")
}