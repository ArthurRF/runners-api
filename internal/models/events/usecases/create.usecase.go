package eventUseCases

import (
	eventRepositories "runners-api/internal/models/events/repositories"
)

func Create(name string, description string) error {
	err := eventRepositories.Create(name, description)

	if err != nil {
		return err
	}

	return nil
}
