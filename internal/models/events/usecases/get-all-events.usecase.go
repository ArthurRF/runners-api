package usecases

import (
	"runners-api/internal/entity"
	"runners-api/internal/models/events/repositories"
)

func GetAllEventsUseCase() ([]entity.Event, error) {
	events, err := repositories.GetAll()

	if err != nil {
		return nil, err
	}

	return events, nil
}
