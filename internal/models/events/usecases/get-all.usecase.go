package eventUseCases

import (
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
)

func GetAll() ([]entity.Event, error) {
	events, err := eventRepositories.GetAll()

	if err != nil {
		return nil, err
	}

	return events, nil
}
