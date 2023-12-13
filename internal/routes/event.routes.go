package routes

import (
	eventControllers "runners-api/internal/models/events/controllers"

	"github.com/go-chi/chi"
)

func EventRoutes(r chi.Router) {
	r.Route("/events", func(r chi.Router) {
		r.Get("/", eventControllers.GetAll)
		r.Get("/{eventID}", eventControllers.GetOneByID)
		r.Post("/", eventControllers.Create)
		r.Delete("/{eventID}", eventControllers.Delete)
	})
}
