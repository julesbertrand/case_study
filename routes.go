package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is the type for all routes containing name, CRUD method, pattern and handlerfunc
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is an array of Route
type Routes []Route

// NewRouter if a func to create a new router handling the routes created below
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"createCustomer",
		"POST",
		"/customers",
		createCustomer,
	},
	Route{
		"getCustomer",
		"GET",
		"/customers/{customerId}",
		getCustomer,
	},
	Route{
		"getCustomers",
		"GET",
		"/customers",
		getCustomers,
	},
	Route{
		"updateCustomer",
		"PUT",
		"/customers/{customerId}",
		updateCustomer,
	},
	Route{
		"deleteCustomer",
		"DELETE",
		"/customers/{customerId}",
		deleteCustomer,
	}}
