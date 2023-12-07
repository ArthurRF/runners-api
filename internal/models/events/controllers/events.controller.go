package eventControllers

import (
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
	// event := new(eventDtos.CreateEventDto)
	// if err := c.BodyParser(event); err != nil {
	// 	return err
	// }

	// validator := validator.New()
	// err := validator.Struct(event)
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	// }

	// err = eventUseCases.Create(event)

	// if err != nil {
	// 	return err
	// }

	return c.Status(fiber.StatusCreated).Send(nil)
}

func Update(c *fiber.Ctx) error {
	// id := c.Params("id")
	// parsedId, err := strconv.Atoi(id)

	// if err != nil {
	// 	return err
	// }

	// event := new(entity.Event)

	// if err := c.BodyParser(event); err != nil {
	// 	return err
	// }

	// eventUpdated, err := eventUseCases.Update(parsedId, event)

	// if err != nil {
	// 	return err
	// }

	return c.Status(fiber.StatusOK).JSON(nil)
}

func Delete(c *fiber.Ctx) error {
	return nil
}
