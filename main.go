package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getID(r *http.Request, idField string) uint {
	params := mux.Vars(r)
	id64, _ := strconv.ParseUint(params[idField], 10, 64)
	idToReturn := uint(id64)
	return idToReturn
}

func initDB(username string, password string, dbName string) {
	var err error
	dataSourceName := username + ":" + password + "@tcp(localhost:3306)/" + "?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	errorCheck(err)

	db.Exec("CREATE DATABASE " + dbName)
	db.Exec("USE " + dbName)
}

func connectDB(dbName string) {
	var err error
	dataSourceName := "root:JH7WhS6c%%@tcp(localhost:3306)/" + dbName + "?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		initDB("root", "JH7WhS6c%%", dbName)
		fmt.Printf("Succesfully created MySQL database: %v\n", dbName)
	} else {
		fmt.Printf("Succesfully connected to MySQL database: %v\n", dbName)
	}

	db.AutoMigrate(&Customer{}, &Plan{}, &Subscription{})
}

func main() {
	router := NewRouter()

	connectDB("customers_db")

	log.Fatal(http.ListenAndServe(":8080", router))
}
