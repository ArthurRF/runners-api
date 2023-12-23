package initializers

import (
	"runners-api/internal/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Routes(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routes.EventRoutes(r)
	routes.UserRoutes(r)
}
