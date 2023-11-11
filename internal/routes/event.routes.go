package routes

import (
	eventControllers "runners-api/internal/models/events/controllers"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(app *fiber.App) {
	eventGroup := app.Group("/events")

	eventGroup.Get("/", eventControllers.GetAll)
	eventGroup.Post("/", eventControllers.Create)
	eventGroup.Put("/:id", eventControllers.Update)
	eventGroup.Delete("/:id", eventControllers.Delete)
}
