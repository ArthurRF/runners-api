package controllers

import (
	"fmt"
	"runners-api/internal/models/events/usecases"

	"github.com/gofiber/fiber/v2"
)

func GetAllEvents(c *fiber.Ctx) error {
	events, err := usecases.GetAllEventsUseCase()

	fmt.Println("eita")

	if err != nil {
		return err
	}

	return c.Status(200).JSON(events)
}
