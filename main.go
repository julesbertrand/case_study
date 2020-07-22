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

func connectDB() {
	var err error
	dataSourceName := "root:JH7WhS6c%%@tcp(localhost:3306)/customers_db?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	errorCheck(err)

	fmt.Println("Succesfully connected to MySQL database")

	db.AutoMigrate(&Customer{}, &Plan{})
}

func main() {
	router := NewRouter()

	connectDB()

	log.Fatal(http.ListenAndServe(":8080", router))
}
