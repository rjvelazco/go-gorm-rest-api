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

	// Users routes
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.CreateUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	// Tasks routes
	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
