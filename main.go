package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// type Customer struct {
// 	// gorm.Model
// 	CustomerID  uint      `json:"customerId" gorm:"primary_key"`
// 	FirstName   string    `json:"firstName"`
// 	LastName    string    `json:"lastName"`
// 	Email       string    `json:"email"`
// 	CreatedAt   time.Time `json:"createdAt"`
// 	CountryCode string    `json:"countryCode"`
// }

var db *gorm.DB

func errorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getId(r *http.Request, idField string) uint {
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
