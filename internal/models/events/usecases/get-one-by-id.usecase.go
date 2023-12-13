package eventUseCases

import (
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
)

func GetOneByID(eventID uint) (*entity.Event, error) {
	event, err := eventRepositories.GetOneByID(eventID)
	if err != nil {
		return nil, err
	}

	return event, nil
}
