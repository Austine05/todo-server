package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Austine05/todo-server/models"
)

// Register a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = uuid.New().String()
	err := models.RegisterUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(newUser)
}

// Update an existing user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := models.GetUserByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	user.FirstName = r.FormValue("firstName")
	user.LastName = r.FormValue("lastName")
	user.Email = r.FormValue("email")
	user.Mobile = r.FormValue("mobile")
	err = models.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Update the password of a user
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := models.GetUserByID(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	user.Password = r.FormValue("password")
	err = models.UpdatePassword(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(user)
}
