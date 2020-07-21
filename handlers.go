package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)

	if db.NewRecord(customer) == false {
		fmt.Println("This customer id already exists.")
	}

	if err := db.Create(&customer).Error; errorCheck(err)

	// db.Create(&customer)
	fmt.Printf("Succesfully created customer %v: %s", customer.CustomerID, customer.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []Customer
	db.Find(&customers)
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputCustomerID := params["customerId"]

	var customer Customer
	db.First(&customer, inputCustomerID)
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var updatedCustomer Customer
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	db.Save(&updatedCustomer)
	fmt.Printf("Succesfully updated customer %v: %s", updatedCustomer.CustomerID, updatedCustomer.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCustomer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputCustomerID := params["customerId"]
	// Convert `orderId` string param to uint64
	id64, _ := strconv.ParseUint(inputCustomerID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("customer_id = ?", idToDelete).Delete(&Customer{})
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Succesfully deleted customer %v: %s", params["customerId"], params["email"])
}
