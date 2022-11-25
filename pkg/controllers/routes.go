package controllers

import "github.com/gorilla/mux"

func InitializeRoutes(router *mux.Router) {
	// Set the main route path
	var customersUrlPath = "/v1/customers"

	// Set the routes
	router.HandleFunc(customersUrlPath, GetCustomers).Methods("GET")
	router.HandleFunc(customersUrlPath+"/{id}", GetCustomerById).Methods("GET")
	router.HandleFunc(customersUrlPath, CreateCustomer).Methods("POST")
	router.HandleFunc(customersUrlPath+"/{id}", UpdateCustomer).Methods("PUT")
}
