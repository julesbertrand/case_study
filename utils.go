package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

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

// // try to make a function to extract filters from url.Values
// func getFilters(f url.Values) map[string]interface{} {
// 	filters := make(map[string]interface{})
// 	for k, v := range f {
// 		filters[k] = v[0]
// 	}
// 	fmt.Println(filters)
// 	return filters
// }

func initDB(username string, password string, dbName string) {
	var err error
	dataSourceName := username + ":" + password + "@tcp(localhost:3306)/" + "?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)
	errorCheck(err)

	db.Exec("CREATE DATABASE " + dbName)
	db.Exec("USE " + dbName)
	fmt.Printf("Succesfully created MySQL database: %v\n", dbName)
}

func connectDB(dbName string) {
	var err error
	dataSourceName := "root:JH7WhS6c%%@tcp(localhost:3306)/" + dbName + "?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		initDB("root", "JH7WhS6c%%", dbName)
	} else {
		fmt.Printf("Succesfully connected to MySQL database: %v\n", dbName)
	}

	db.AutoMigrate(&Customer{}, &Plan{}, &Subscription{})
}
