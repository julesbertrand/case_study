package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Customer handlers
func createCustomer(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	var item Customer
	json.NewDecoder(r.Body).Decode(&item)

	item.CreatedAt = time.Now()
	db.Create(&item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

	fmt.Printf("Succesfully created item %v: %s \n", item.CustomerID, item.Email)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	var items []Customer

	// filters := getFilters(r.URL.Query())
	// if len(filters) != 0 {
	// 	db.Where(filters).Find(&items)
	// 	if len(items) == 0 {
	// 		fmt.Fprintln(w, "No items in database for these filters")
	// 		return
	// 	}
	// } else {
	// 	db.Find(&items)
	// }
	db.Find(&items)

	json.NewEncoder(w).Encode(items)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	inputItemID := getID(r, "customerId")
	var item Customer
	db.First(&item, inputItemID)

	if item.CustomerID == 0 {
		fmt.Fprintln(w, "Trying to get a non-existing item")
		return
	}

	json.NewEncoder(w).Encode(item)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	var item Customer
	json.NewDecoder(r.Body).Decode(&item)
	item.CustomerID = getID(r, "customerId")

	db.Save(&item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

	fmt.Printf("Succesfully updated item %v: %s \n", item.CustomerID, item.Email)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	idToDelete := getID(r, "customerId")
	var item Customer
	db.First(&item, idToDelete)

	if item.Email == "" && item.CustomerID == 0 {
		fmt.Fprintln(w, "Trying to delete a non-existing item")
		return
	}

	db.Where("customer_id = ?", idToDelete).Delete(&Customer{})
	w.WriteHeader(http.StatusNoContent)

	fmt.Printf("Succesfully deleted item %v: %s \n", item.CustomerID, item.Email)
}

// Plans handlers
func createPlan(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	var item Plan
	json.NewDecoder(r.Body).Decode(&item)

	item.CreatedAt = time.Now()
	db.Create(&item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

	fmt.Printf("Succesfully created item %v: %s \n", item.PlanID, item.PlanName)
}

func getPlans(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	var items []Plan

	// filters := getFilters(r.URL.Query())
	// if len(filters) != 0 {
	// 	db.Where(filters).Find(&items)
	// 	if len(items) == 0 {
	// 		fmt.Fprintln(w, "No items in database for these filters")
	// 		return
	// 	}
	// } else {
	// 	db.Find(&items)
	// }
	db.Find(&items)

	json.NewEncoder(w).Encode(items)
}

func getPlan(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")

	inputItemID := getID(r, "planId")
	var item Plan

	db.First(&item, inputItemID)
	if item.PlanID == 0 {
		fmt.Println("Trying to get a non-existing item")
		return
	}

	json.NewEncoder(w).Encode(item)
}

func updatePlan(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	var item Plan
	json.NewDecoder(r.Body).Decode(&item)
	item.PlanID = getID(r, "planId")

	db.Save(&item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

	fmt.Printf("Succesfully updated item %v: %s \n", item.PlanID, item.PlanName)
}

func deletePlan(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	idToDelete := getID(r, "planId")
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

// Subscriptions handler
func addSubscription(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	var item Subscription
	json.NewDecoder(r.Body).Decode(&item)

	deactivateSubscriptions(item)

	item.CreatedAt = time.Now()
	item.ActiveStatus = true
	db.Create(&item)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

	fmt.Printf("Succesfully created subscription for plan %v for customer %v \n", item.PlanID, item.CustomerID)
}

func deactivateSubscriptions(item Subscription) {
	db.Model(&item).Where("customer_id = ? AND active_status = ?", item.CustomerID, true).Updates(map[string]interface{}{"activeStatus": false, "endedAt": time.Now()})
}

func getSubscriptions(w http.ResponseWriter, r *http.Request) {
	connectDB("customers_db")
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	var items []Subscription
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}
