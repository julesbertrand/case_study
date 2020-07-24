package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
