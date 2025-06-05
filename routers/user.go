package routing

import (
	"net/http"

	"github.com/go-chi/chi"
)

func UserRes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is user"))
}

func UserRoutes(users chi.Router) {

	users.Get("/", UserRes)

}
