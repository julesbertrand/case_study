package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

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
	},
}
