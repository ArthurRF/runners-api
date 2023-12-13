package eventUseCases

import (
	eventRepositories "runners-api/internal/models/events/repositories"
)

func Delete(eventID uint) error {
	_, err := eventRepositories.GetOneByID(eventID)
	if err != nil {
		return err
	}

	err = eventRepositories.Delete(eventID)
	if err != nil {
		return err
	}

	return nil
}
