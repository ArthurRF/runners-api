package controllers

import (
	"runners-api/internal/models/events/usecases"

	"github.com/gofiber/fiber/v2"
)

func GetAllEvents(c *fiber.Ctx) error {
	events, err := usecases.GetAllEventsUseCase()

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(events)
}
