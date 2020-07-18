package main

import (
	"log"
	"net/http"
)

func main() {
	r := NewRouter()
	log.Printf("Listening at %d port...", DefaultServicePort)
	log.Fatal(http.ListenAndServe(":3333", r))
}
