package eventRepositories

import (
	"errors"
	"fmt"
	"net/http"
	config "runners-api/configs"
	"runners-api/internal/entity"
	"runners-api/internal/shared"

	"gorm.io/gorm"
)

func GetAll() ([]entity.Event, *shared.AppError) {
	var events []entity.Event

	if err := config.DB.Find(&events).Error; err != nil {
		return nil, &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return events, nil
}

func GetOneByID(ID uint) (*entity.Event, *shared.AppError) {
	var eventFound entity.Event

	if err := config.DB.First(&eventFound, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &shared.AppError{
				Message:    fmt.Sprintf("event with ID %d not found.", ID),
				StatusCode: http.StatusNotFound,
			}
		}
		return nil, &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &eventFound, nil
}

func Create(name string, description string) *shared.AppError {
	eventToCreate := entity.Event{
		Name:        name,
		Description: description,
	}

	returnData := config.DB.Create(&eventToCreate)
	if returnData.Error != nil {
		return &shared.AppError{
			Message:    returnData.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func Delete(ID uint) *shared.AppError {
	returnData := config.DB.Delete(&entity.Event{}, ID)
	if returnData.Error != nil {
		return &shared.AppError{
			Message:    returnData.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}

func Update(id int, event *entity.Event) (*entity.Event, *shared.AppError) {
	return nil, &shared.AppError{
		Message:    "method not implemented.",
		StatusCode: http.StatusInternalServerError,
	}
}
