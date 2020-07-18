package main

import "github.com/gorilla/mux"

// NewRouter is the function which creates the mux router according to the paths and their handlers
func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.Path("/healthz").HandlerFunc(Health)
	router.PathPrefix("/").HandlerFunc(Reflect)

	return router
}
