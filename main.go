package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rjvelazco/go-gorm-restapi/db"
	"github.com/rjvelazco/go-gorm-restapi/models"
	"github.com/rjvelazco/go-gorm-restapi/routes"
)

func main() {
	db.DBConnection()

	if err := db.DB.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
