package controllers

import (
	"encoding/json"
	"net/http"
	"test-project-backend/pkg/database"
	"test-project-backend/pkg/entities"
	"test-project-backend/pkg/utils"

	"github.com/gorilla/mux"
)

func GetCustomers(w http.ResponseWriter, _ *http.Request) {
	// Create variable customer array
	var customers []entities.Customer

	// Return all rows
	database.Instance.Find(&customers)

	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// Set status for response with code 200
	w.WriteHeader(http.StatusOK)

	// Return the all rows
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerById(w http.ResponseWriter, request *http.Request) {
	// Get customerId from request
	customerId := mux.Vars(request)["id"]

	// Check if customer exists
	if checkIfCustomerExists(customerId) == false {
		json.NewEncoder(w).Encode("Customer Not Found!")
		return
	}

	// Initiate customer based on entity Customer
	var customer entities.Customer

	// Search for the customer
	database.Instance.First(&customer, customerId)

	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// return the customer by id
	json.NewEncoder(w).Encode(customer)
}

func CreateCustomer(w http.ResponseWriter, request *http.Request) {
	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// Initiate customer based on entity Customer
	var customer entities.Customer

	// Assign requested parameters from body to variable
	json.NewDecoder(request.Body).Decode(&customer)

	// hash password
	password, err := utils.HashPassword(customer.Password)
	if err != nil {
		panic("Fail to create customer")
	}
	customer.Password = password

	// Insert row
	database.Instance.Create(&customer)

	// return newly created customer
	json.NewEncoder(w).Encode(customer)
}

func UpdateCustomer(w http.ResponseWriter, request *http.Request) {
	// Get customer id from request
	customerId := mux.Vars(request)["id"]

	// Check if customer exists
	if checkIfCustomerExists(customerId) == false {
		json.NewEncoder(w).Encode("Customer Not Found!")
		return
	}

	// Initiate customer based on entity Customer
	var customer entities.Customer

	// Search for the customer
	database.Instance.First(&customer, customerId)

	// Assign requested parameters from body to variable
	json.NewDecoder(request.Body).Decode(&customer)

	// hash password
	password, err := utils.HashPassword(customer.Password)
	if err != nil {
		panic("Fail to create customer")
	}
	customer.Password = password

	// Update row in database
	database.Instance.Save(&customer)

	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// return newly updated customer
	json.NewEncoder(w).Encode(customer)
}

func checkIfCustomerExists(customerId string) bool {
	// Initiate customer based on entity Customer
	var customer entities.Customer

	// Search for the customer
	database.Instance.First(&customer, customerId)

	// Check if the customer is valid
	if customer.ID == 0 {
		return false
	}
	return true
}
