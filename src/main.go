package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Define a route
	r.HandleFunc("/api/metrics", getMetrics).Methods("GET")

	// Start the server
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", r)
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	// Send a response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Metrics endpoint"})
}