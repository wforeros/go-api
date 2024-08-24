package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var dsn string = "host=localhost user=postgres password=postgres dbname=tasks-go port=5434"
	var error error
	DB, error = gorm.Open(postgres.Open(dsn))
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Succesfully connected to Postg")
	}
}
