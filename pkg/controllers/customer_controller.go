package controllers

import (
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"test-project-backend/pkg/database"
	"test-project-backend/pkg/entities"
	"test-project-backend/pkg/utils"

	"github.com/gorilla/mux"
)

const ServerFailMessage = "Internal fail"

func GetCustomers(w http.ResponseWriter, _ *http.Request) {
	// Create variable customer array
	var customers []entities.Customer
	var customerResponse []entities.CustomerResponse

	// Return all rows
	if err := database.Instance.Select("id, first_name, last_name, email").Find(&customers).Scan(&customerResponse).Error; err != nil {
		panic(ServerFailMessage)
	}

	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// Set status for response with code 200
	w.WriteHeader(http.StatusOK)

	// Return the all rows
	json.NewEncoder(w).Encode(customerResponse)
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
	var customerResponse entities.CustomerResponse

	// Search for the customer
	if err := database.Instance.Select("id, first_name, last_name, email").First(&customer,
		customerId).Scan(&customerResponse).Error; err != nil {
		panic(ServerFailMessage)
	}

	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// return the customer by id
	json.NewEncoder(w).Encode(customerResponse)
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
	if err := database.Instance.Create(&customer).Error; err != nil {
		panic(ServerFailMessage)
	}

	// return newly created customer
	json.NewEncoder(w).Encode(entities.CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	})
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
	if err := database.Instance.First(&customer, customerId).Error; err != nil {
		panic(ServerFailMessage)
	}

	// Assign requested parameters from body to variable
	json.NewDecoder(request.Body).Decode(&customer)

	// hash password
	password, err := utils.HashPassword(customer.Password)
	if err != nil {
		panic("Fail to create customer")
	}
	customer.Password = password

	// Update row in database
	if err := database.Instance.Save(&customer).Error; err != nil {
		panic(ServerFailMessage)
	}

	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// return newly updated customer
	json.NewEncoder(w).Encode(entities.CustomerResponse{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	})
}

func DeleteCustomer(w http.ResponseWriter, request *http.Request) {
	// Get customer id from request
	customerId := mux.Vars(request)["id"]

	// Check if customer exists
	if checkIfCustomerExists(customerId) == false {
		json.NewEncoder(w).Encode("Customer Not Found!")
		return
	}

	// Initiate customer based on entity Customer
	var customer entities.Customer

	// Update row in database
	if err := database.Instance.Delete(&customer, customerId).Error; err != nil {
		panic(ServerFailMessage)
	}

	// Set header to proper Type
	w.Header().Set("Content-Type", "application/json")

	// return newly updated customer
	json.NewEncoder(w).Encode("Customer deleted")
}

func checkIfCustomerExists(customerId string) bool {
	// Initiate customer based on entity Customer
	var customer entities.Customer

	// Search for the customer
	if err := database.Instance.First(&customer, customerId).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(ServerFailMessage)
		}
	}

	// Check if the customer is valid
	if customer.ID == 0 {
		return false
	}
	return true
}
