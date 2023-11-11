package initializers

import (
	"runners-api/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	routes.EventRoutes(app)
}
