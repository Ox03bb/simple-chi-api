package view

import (
	"encoding/json"
	"learn/models"
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
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

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	var db = models.GetDB()
	var user models.User

	id := chi.URLParam(r, "id")

	err := db.First(&user, id).Error
	if err != nil {
		http.Error(w, "User with id "+id+" Not Found", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var db = models.GetDB()
	var user models.User

	var updateData struct {
		Username  *string `json:"username"`
		Email     *string `json:"email"`
		IsAdmin   *bool   `json:"is_admin"`
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
		Phone     *string `json:"phone"`
	}

	id := chi.URLParam(r, "id")

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	err = db.First(&user, uid).Error
	if err != nil {
		http.Error(w, "User with id "+id+" Not Found", 404)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if updateData.Username != nil {
		user.Username = *updateData.Username
	}
	if updateData.Email != nil {
		user.Email = *updateData.Email
	}
	if updateData.IsAdmin != nil {
		user.IsAdmin = updateData.IsAdmin
	}
	if updateData.FirstName != nil {
		user.FirstName = updateData.FirstName
	}
	if updateData.LastName != nil {
		user.LastName = updateData.LastName
	}
	if updateData.Phone != nil {
		user.Phone = updateData.Phone
	}

	if err := db.Save(&user).Error; err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var db = models.GetDB()
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

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
