package eventUseCases

import (
	eventDtos "runners-api/internal/models/events/dtos"
	eventRepositories "runners-api/internal/models/events/repositories"
)

func Create(event *eventDtos.CreateEventDto) error {
	err := eventRepositories.Create(event)

	if err != nil {
		return err
	}

	return nil
}
