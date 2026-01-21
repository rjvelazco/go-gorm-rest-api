package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=rjvelazco password=password dbname=gorm port=5432 sslmode=disable TimeZone=America/Caracas"

var DB *gorm.DB

func DBConnection() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		log.Println("Connected to the database")
	}
	return DB, nil
}
