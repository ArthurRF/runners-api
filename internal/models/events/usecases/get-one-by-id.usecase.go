package eventUseCases

import (
	"errors"
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
)

func GetOneByID(id string) (*entity.Event, error) {
	event, err := eventRepositories.GetOneById(id)
	if err != nil {
		return nil, errors.New("event not found")
	}

	return &event, nil
}
