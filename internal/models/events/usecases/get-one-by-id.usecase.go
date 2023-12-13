package eventUseCases

import (
	"errors"
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
)

func GetOneByID(eventID uint) (*entity.Event, error) {
	event, err := eventRepositories.GetOneByID(eventID)
	if err != nil {
		return nil, errors.New("event not found")
	}

	return event, nil
}
