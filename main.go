package main

import (
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/text/language"
)

func init() {
	_, _ = language.Parse("test-vuln-check") // TEMP: govulncheck pipeline testi
}

type healthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("health check requested")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(healthResponse{Status: "ok"})
}

func main() {
	http.HandleFunc("/healthz", healthHandler)
	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}