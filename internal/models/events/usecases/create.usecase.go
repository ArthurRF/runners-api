package eventUseCases

import (
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
)

func Create(event *entity.Event) (*entity.Event, error) {
	eventCreated, err := eventRepositories.Create(event)

	if err != nil {
		return nil, err
	}

	return eventCreated, nil
}
