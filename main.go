package main

import (
	"discovergy/internal"
	"discovergy/pkg"
	"log"
	"net/http"
)

func main() {
	r := pkg.NewRouter()
	log.Printf("Listening at %d port...", internal.DefaultServicePort)
	log.Fatal(http.ListenAndServe(":3333", r))
}
