package eventUseCases

import (
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
)

func Update(id int, event *entity.Event) (*entity.Event, error) {
	eventUpdated, err := eventRepositories.Update(id, event)

	if err != nil {
		return nil, err
	}

	return eventUpdated, nil
}
