package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter is the function which creates the mux router according to the paths and their handlers
func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.Path("/healthz").HandlerFunc(Health).Methods(http.MethodGet)
	router.Path("/encode/{input}").HandlerFunc(Encode).Methods(http.MethodGet)
	router.PathPrefix("/").HandlerFunc(Reflect).Methods(http.MethodGet)

	return router
}
