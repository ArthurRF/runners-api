package eventControllers

import (
	"runners-api/internal/entity"
	eventUseCases "runners-api/internal/models/events/usecases"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	events, err := eventUseCases.GetAll()

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func Create(c *fiber.Ctx) error {
	event := new(entity.Event)
	if err := c.BodyParser(event); err != nil {
		return err
	}

	eventCreated, err := eventUseCases.Create(event)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(eventCreated)
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}
