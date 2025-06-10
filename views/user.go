package view

import (
	"encoding/json"
	"learn/models"
	"log"
	"net/http"
)

// func UserView(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodPost {

// 		w.Write([]byte("User View: GET request"))

// 	} else if r.Method == http.MethodPost {

// 		w.Write([]byte("User View: POST request"))

// 	}
// }

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	var db = models.GetDB()
	var users []models.User

	err := db.Find(&users).Error
	if err != nil {
		http.Error(w, "Error fetching users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var db = models.GetDB()
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	log.Println("Username:", user.Username, user.UserNameValidation())
	log.Println("Email:", user.Email, user.EmailValidation())

	if !(user.PasswordValidation()) {
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}

	err = db.Create(&user).Error
	if err != nil {
		http.Error(w, "Error creating user: "+err.Error(), http.StatusFailedDependency)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
