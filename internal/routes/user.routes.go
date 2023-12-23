package routes

import (
	userControllers "runners-api/internal/models/users/controllers"

	"github.com/go-chi/chi"
)

func UserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/auth", userControllers.Auth)
	})
}
