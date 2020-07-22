package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)

	// errorCheck(err)
	db.Create(&customer)
	fmt.Printf("Succesfully created customer %v: %s \n", customer.CustomerID, customer.Email)
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
	inputCustomerID := getId(r, "customerId")

	var customer Customer
	db.First(&customer, inputCustomerID)
	if customer.CustomerID == 0 {
		fmt.Println("Trying to get a non-existing customer")
		return
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var updatedCustomer Customer
	json.NewDecoder(r.Body).Decode(&updatedCustomer)
	updatedCustomer.CustomerID = getId(r, "customerId")
	fmt.Printf("%+v\n", updatedCustomer)
	db.Save(&updatedCustomer)

	fmt.Printf("Succesfully updated customer %v: %s \n", updatedCustomer.CustomerID, updatedCustomer.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCustomer)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	idToDelete := getId(r, "customerId")
	var customer Customer
	db.First(&customer, idToDelete)
	if customer.CustomerID == 0 {
		fmt.Println("Trying to delete a non-existing customer")
		return
	} else {
		db.Where("customer_id = ?", idToDelete).Delete(&Customer{})
		w.WriteHeader(http.StatusNoContent)
		fmt.Printf("Succesfully deleted customer %v: %s \n", customer.CustomerID, customer.Email)
	}
}
