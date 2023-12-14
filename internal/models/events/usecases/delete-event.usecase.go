package eventUseCases

import (
	eventRepositories "runners-api/internal/models/events/repositories"
	"runners-api/internal/shared"
)

func Delete(eventID uint) *shared.AppError {
	_, appErr := eventRepositories.GetOneByID(eventID)
	if appErr != nil {
		return appErr
	}

	appErr = eventRepositories.Delete(eventID)
	if appErr != nil {
		return appErr
	}

	return nil
}
