package eventRepositories

import (
	"errors"
	"fmt"
	config "runners-api/configs"
	"runners-api/internal/entity"

	"gorm.io/gorm"
)

func GetAll() ([]entity.Event, error) {
	var events []entity.Event

	if err := config.DB.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}

func GetOneByID(ID uint) (*entity.Event, error) {
	var eventFound entity.Event

	if err := config.DB.First(&eventFound, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("event with ID %d not found", ID)
		}
		return nil, err
	}

	return &eventFound, nil
}

func Create(name string, description string) error {
	eventToCreate := entity.Event{
		Name:        name,
		Description: description,
	}

	returnData := config.DB.Create(&eventToCreate)
	if returnData.Error != nil {
		return returnData.Error
	}

	return nil
}

func Delete(ID uint) error {
	returnData := config.DB.Delete(&entity.Event{}, ID)
	if returnData.Error != nil {
		return returnData.Error
	}

	return nil
}

func Update(id int, event *entity.Event) (*entity.Event, error) {
	return nil, errors.New("method not implemented")

	// return event, nil
}
