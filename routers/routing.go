package routing

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func RRes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is RRes"))
}

func Routing() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api", func(api chi.Router) {
		api.Route("/users", UserRoutes)

		api.Get("/", RRes)

	})

	return r
}
