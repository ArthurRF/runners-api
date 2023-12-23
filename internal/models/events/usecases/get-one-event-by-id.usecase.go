package eventUseCases

import (
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
	"runners-api/internal/shared"
)

func GetOneByID(eventID uint) (*entity.Event, *shared.AppError) {
	event, err := eventRepositories.GetOneByID(eventID)
	if err != nil {
		return nil, err
	}

	return event, nil
}
