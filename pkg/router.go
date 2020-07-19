package pkg

import (
	"net/http"

	"discovergy/internal"
	"discovergy/pkg/svc"

	"github.com/gorilla/mux"
)

// NewRouter is the function which creates the mux router according to the paths and their handlers
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = internal.Logger(handler, route.Name)

		router.
			Path(route.Pattern).
			Methods(route.Method).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	router.PathPrefix("/").HandlerFunc(svc.Reflect).Methods(http.MethodGet)
	return router
}
