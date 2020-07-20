package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	CustomerId   uint      `json:"customerId" gorm:"primary_key"`
	Email        string    `json:"email address"`
	FirstName    string    `json:"firstname"`
	LastName     string    `json:"lastname"`
	CreationDate time.Time `json:"creationdate"`
	CountryCode  string    `json:"email address" gorm:"countrykey"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:JH7WhS6c%%@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("Connection to database failed")
	}

	// create database and migrate to create table
	db.Exec("CREATE DATABASE customers_db")
	db.Exec("USE customers_db")
	db.AutoMigrate(&Customer{})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.PrintIn(w, "Welcome to the HomePage!")
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	json.NewDecoder(r.Body).Decode(&customer)
	db.Create(&customer)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var customers []Customer
	db.Find(&customers)
	json.NewEncoder(w).Encode(customers)
}

func main() {
	router := mux.NewRouter()
	// Create
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	//Read-all
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	initDB()

	// http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
