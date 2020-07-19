package main

import (
	"log"
	"net/http"
	"strconv"

	"discovergy/internal"
	"discovergy/pkg"
)

func main() {
	r := pkg.NewRouter()
	log.Printf("Listening at %d port...", internal.DefaultServicePort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(internal.DefaultServicePort), r))
}
