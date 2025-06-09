package view

import (
	"net/http"
)

func UserView(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		w.Write([]byte("User View: GET request"))

	} else if r.Method == http.MethodPost {

		w.Write([]byte("User View: POST request"))

	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

}
