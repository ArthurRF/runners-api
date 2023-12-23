package eventUseCases

import (
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
	"runners-api/internal/shared"
)

func Update(id int, event *entity.Event) (*entity.Event, *shared.AppError) {
	eventUpdated, appErr := eventRepositories.Update(id, event)

	if appErr != nil {
		return nil, appErr
	}

	return eventUpdated, nil
}
