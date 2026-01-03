package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/health", handleHealth)
	http.HandleFunc("/api/greet", handleGreet)

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := Response{
		Message: "Welcome to the GitHub Actions Tutorial API",
		Version: "1.0.0",
		Status:  "running",
	}
	json.NewEncoder(w).Encode(resp)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := HealthResponse{
		Status: "healthy",
	}
	json.NewEncoder(w).Encode(resp)
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	
	w.Header().Set("Content-Type", "application/json")
	resp := Response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Version: "1.0.0",
		Status:  "success",
	}
	json.NewEncoder(w).Encode(resp)
}
