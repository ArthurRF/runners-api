package eventControllers

import (
	"encoding/json"
	"net/http"
	eventUseCases "runners-api/internal/models/events/usecases"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	events, err := eventUseCases.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

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
	//     return err
	// }

	// event := new(entity.Event)

	// if err := c.BodyParser(event); err != nil {
	//     return err
	// }

	// eventUpdated, err := eventUseCases.Update(parsedId, event)

	// if err != nil {
	//     return err
	// }

	w.WriteHeader(http.StatusOK)
}
