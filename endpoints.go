package main

import (
	"log"
	"net/http"
	"strings"
)

// Reflect endpoint reflects the requested endpoint URI
func Reflect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(strings.Trim(path, "/")))
	if err != nil {
		log.Fatal("Error writing or returning the response", err)
	}
}

// Health endpoint returns the healthy string if the service is reachable
func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "plain/text; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hi there, I am healthy!!"))
	if err != nil {
		log.Fatal("Error encoding or returning the response", err)
	}
}
