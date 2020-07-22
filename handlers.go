package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Customer handlers
func createCustomer(w http.ResponseWriter, r *http.Request) {
	var item Customer
	json.NewDecoder(r.Body).Decode(&item)
	item.CreatedAt = time.Now()
	db.Create(&item)
	fmt.Printf("Succesfully created item %v: %s \n", item.CustomerID, item.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []Customer
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	inputItemID := getId(r, "customerId")

	var item Customer
	db.First(&item, inputItemID)
	if item.CustomerID == 0 {
		fmt.Println("Trying to get a non-existing item")
		return
	}
	json.NewEncoder(w).Encode(item)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var item Customer
	json.NewDecoder(r.Body).Decode(&item)
	item.CustomerID = getId(r, "customerId")
	fmt.Printf("%+v\n", item)
	db.Save(&item)

	fmt.Printf("Succesfully updated item %v: %s \n", item.CustomerID, item.Email)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	idToDelete := getId(r, "customerId")
	var item Customer
	db.First(&item, idToDelete)
	if item.Email == "" && item.CustomerID == 0 {
		fmt.Println("Trying to delete a non-existing item")
		return
	}
	db.Where("customer_id = ?", idToDelete).Delete(&Customer{})
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Succesfully deleted item %v: %s \n", item.CustomerID, item.Email)
}

// Plans handlers
func createPlan(w http.ResponseWriter, r *http.Request) {
	var item Plan
	json.NewDecoder(r.Body).Decode(&item)
	item.CreationDate = time.Now()
	db.Create(&item)
	fmt.Printf("Succesfully created item %v: %s \n", item.PlanID, item.PlanName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func getPlans(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []Plan
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func getPlan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	inputItemID := getId(r, "planId")

	var item Plan
	db.First(&item, inputItemID)
	if item.PlanID == 0 {
		fmt.Println("Trying to get a non-existing item")
		return
	}
	json.NewEncoder(w).Encode(item)
}

func updatePlan(w http.ResponseWriter, r *http.Request) {
	var item Plan
	json.NewDecoder(r.Body).Decode(&item)
	item.PlanID = getId(r, "planId")
	fmt.Printf("%+v\n", item)
	db.Save(&item)

	fmt.Printf("Succesfully updated item %v: %s \n", item.PlanID, item.PlanName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func deletePlan(w http.ResponseWriter, r *http.Request) {
	idToDelete := getId(r, "planId")
	var item Plan
	db.First(&item, idToDelete)
	if item.PlanName == "" && item.PlanID == 0 {
		fmt.Println("Trying to delete a non-existing item")
		return
	}
	db.Where("plan_id = ?", idToDelete).Delete(&Plan{})
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Succesfully deleted item %v: %s \n", item.PlanID, item.PlanName)
}
