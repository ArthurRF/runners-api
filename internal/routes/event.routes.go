package routes

import (
	"runners-api/internal/models/events/controllers"

	"github.com/gofiber/fiber/v2"
)

func EventRoutes(app *fiber.App) {
	eventGroup := app.Group("/events")

	eventGroup.Get("/", controllers.GetAllEvents)
	// eventGroup.Post("/", controllers.CreateUser)
	// eventGroup.Put("/:id", controllers.UpdateUser)
	// eventGroup.Delete("/:id", controllers.DeleteUser)
}
