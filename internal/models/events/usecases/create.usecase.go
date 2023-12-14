package eventUseCases

import (
	eventRepositories "runners-api/internal/models/events/repositories"
	"runners-api/internal/shared"
)

func Create(name string, description string) *shared.AppError {
	appErr := eventRepositories.Create(name, description)

	if appErr != nil {
		return appErr
	}

	return nil
}
