package svc

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"discovergy/internal"

	"github.com/gorilla/mux"
)

// Reflect endpoint reflects the requested endpoint URI
func Reflect(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(strings.Trim(path, "/")))
	if err != nil {
		log.Fatal("Error writing or returning the response", err)
	}
}

// Health endpoint returns the healthy string if the svc is reachable
func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hi there, I am healthy!!"))
	if err != nil {
		log.Fatal("Error encoding or returning the response", err)
	}
}

// Encode endpoint will take the input string in the request and perform the encryption using ceaser cipher algorithm
// https://en.wikipedia.org/wiki/Caesar_cipher
func Encode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	input, present := vars["input"]
	if !present {
		log.Fatalf("Error input params. Required input string which is to be encoded: %+v", vars)
	}

	result := internal.CeaserCipher(input, internal.Shift, internal.CeaserCipherEncode)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Fatal("Error encoding or returning the response", err)
	}
}

// Decode endpoint will take the input string in the request and perform the decryption using ceaser cipher algorithm
func Decode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	input, present := vars["input"]
	if !present {
		log.Fatalf("Error input params. Required input string which is to be decoded: %+v", vars)
	}

	result := internal.CeaserCipher(input, internal.Shift, internal.CeaserCipherDecode)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Fatal("Error encoding or returning the response", err)
	}
}
