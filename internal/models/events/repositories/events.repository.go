package eventRepositories

import (
	"errors"
	config "runners-api/configs"
	"runners-api/internal/entity"

	eventDtos "runners-api/internal/models/events/dtos"
)

func GetAll() ([]entity.Event, error) {
	var events []entity.Event

	if err := config.DB.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func Create(event *eventDtos.CreateEventDto) error {
	// eventToCreate := entity.Event{
	// 	Name:        event.Name,
	// 	Description: event.Description,
	// }

	// returnData := config.DB.Create(&eventToCreate)
	// if returnData.Error != nil {
	// 	return returnData.Error
	// }

	return errors.New("method not implemented")
}

func Update(id int, event *entity.Event) (*entity.Event, error) {
	return nil, errors.New("method not implemented")

	// return event, nil
}
