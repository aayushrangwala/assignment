package pkg

import (
	"net/http"

	"discovergy/api"
	"discovergy/internal"
	"discovergy/pkg/svc"

	"github.com/gorilla/mux"
)

// NewRouter is the function which creates the mux router according to the paths and their handlers
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range api.Routes {
		handler := internal.Logger(route.HandlerFunc, route.Name)

		router.
			Path(route.Pattern).
			Methods(route.Method).
			Name(route.Name).
			Handler(handler)
	}

	h := internal.Logger(http.HandlerFunc(svc.Reflect), "Reflector")
	router.PathPrefix("/").HandlerFunc(h).Methods(http.MethodGet)
	return router
}
