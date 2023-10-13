package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=go_db user=viviana password=12345 dbname=ice_cream port=5432"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}

}
