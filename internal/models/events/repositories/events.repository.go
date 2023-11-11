package repositories

import (
	config "runners-api/configs"
	"runners-api/internal/entity"
)

func GetAll() ([]entity.Event, error) {
	var events []entity.Event

	if err := config.DB.Find(&events).Error; err != nil {
		return nil, err
	}

	return events, nil
}
