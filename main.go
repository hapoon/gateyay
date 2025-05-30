package main

import (
	"gateyay/handler"
	"log"
	"net/http"
)

const (
	name    = "gateyay" //nolint:unused
	version = "0.0.1"   //nolint:unused
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/health", handler.HealthHandler)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
