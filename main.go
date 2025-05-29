package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	name    = "gateyay" //nolint:all
	version = "0.0.1"   //nolint:all
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, World!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "OK")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
