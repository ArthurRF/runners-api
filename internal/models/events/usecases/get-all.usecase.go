package eventUseCases

import (
	"runners-api/internal/entity"
	eventRepositories "runners-api/internal/models/events/repositories"
	"runners-api/internal/shared"
)

func GetAll() ([]entity.Event, *shared.AppError) {
	events, appErr := eventRepositories.GetAll()

	if appErr != nil {
		return nil, appErr
	}

	return events, nil
}
