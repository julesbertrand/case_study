package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	// gorm.Model
	CustomerID  uint      `json:"customerId" gorm:"primary_key"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	CreatedAt   time.Time `json:"createdAt"`
	CountryCode string    `json:"countryCode"`
}

var db *gorm.DB

func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func connectDB() {
	var err error
	dataSourceName := "root:JH7WhS6c%%@tcp(localhost:3306)/customers_db?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	errorCheck(err)

	fmt.Println("Succesfully connected to MySQL database")

	db.AutoMigrate(&Customer{})
}

// func main() {
// 	router := mux.NewRouter()

// 	connectDB()

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	// Read
	router.HandleFunc("/customers/{customerId}", getCustomer).Methods("GET")
	// Read-all
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	// Update
	router.HandleFunc("/customers/{customerId}", updateCustomer).Methods("PUT")
	// Delete
	router.HandleFunc("/customers/{customerId}", deleteCustomer).Methods("DELETE")
	// Initialize db connection
	connectDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	db.Create(&customer)
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
	id64, _ := strconv.ParseUint(inputCustomerID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("customer_id = ?", idToDelete).Delete(&Customer{})
	w.WriteHeader(http.StatusNoContent)
	fmt.Printf("Succesfully deleted customer %v: %s", params["customerId"], params["email"])
}
