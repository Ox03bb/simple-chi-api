package routing

import (
	view "learn/views"

	"github.com/go-chi/chi"
)

// the first method of routing
// the second method using HandleFunc()
func UserRoutes(users chi.Router) {

	users.Get("/", view.UserView)
	users.Post("/", view.CreateUser)

}
