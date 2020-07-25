package api

import (
	"net/http"

	"discovergy/pkg/svc"
)

// Route struct defines the http route mapping
type Route struct {
	// Name of the Route
	Name string

	// Method is the http method to be used for that route
	Method string

	// Pattern is the route path pattern for the endpoint
	Pattern string

	// HandlerFunc is the http handler function for the endpoint route
	HandlerFunc http.HandlerFunc
}

// routes is the list of all the routes with all the paths and its related handler functions and pattern
var Routes = []Route{
	// This Route with the name Health represents the path, method and handler for a route
	{
		Name:        "Health",
		Method:      http.MethodGet,
		Pattern:     "/healthz",
		HandlerFunc: svc.Health,
	},

	// This Route with the name Encoder represents the path, method and handler for an encoding route
	{
		Name:        "Encoder",
		Method:      http.MethodGet,
		Pattern:     "/encode/{input}",
		HandlerFunc: svc.Encode,
	},

	// This Route with the name Decoder represents the path, method and handler for an decoding endpoint
	{
		Name:        "Decoder",
		Method:      http.MethodGet,
		Pattern:     "/decode/{input}",
		HandlerFunc: svc.Decode,
	},
}
