package routes

import (
	"encoding/json"
	"net/http"

	"github.com/rjvelazco/go-gorm-restapi/db"
	"github.com/rjvelazco/go-gorm-restapi/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	db.DB.Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Getting user..."))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating user: " + err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user..."))
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user..."))
}
