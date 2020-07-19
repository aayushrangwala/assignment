package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"

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

// Health endpoint returns the healthy string if the service is reachable
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

	result := ceaserCipherEncode(input, EncryptionShift)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Fatal("Error encoding or returning the response", err)
	}
}

func ceaserCipherEncode(in string, key int) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) {
			return r
		}
		// since the alphabets are all lower case in the requirement
		if unicode.IsUpper(r) {
			r += 32 // make lower
		}
		r += rune(key % AlphabetsLength)
		if r > 'z' {
			r = 'a' + (r - ('z' + 1))
		} else if r < 'a' {
			r = 'z' - (('a' - 1) - r)
		}
		return r
	}, in)
}
