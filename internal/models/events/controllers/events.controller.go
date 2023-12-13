package eventControllers

import (
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	// events, err := eventUseCases.GetAll()

	// if err != nil {
	// 	return err
	// }

	w.WriteHeader(http.StatusOK)
}

func Create(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
}

func Update(w http.ResponseWriter, r *http.Request) {
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

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
