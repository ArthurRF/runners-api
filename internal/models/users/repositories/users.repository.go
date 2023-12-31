package userRepositories

import (
	"errors"
	"net/http"
	config "runners-api/configs"
	"runners-api/internal/entity"
	"runners-api/internal/shared"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FindByClerkID(clerkID string) (*entity.User, *shared.AppError) {
	var userFound entity.User

	if err := config.DB.First(&userFound, "clerk_id = ?", clerkID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, &shared.AppError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &userFound, nil
}

func Create(clerkID string, name string, avatarUrl string, email string) (*entity.User, *shared.AppError) {
	user := entity.User{
		Name:      name,
		Email:     email,
		AvatarUrl: avatarUrl,
		ClerkID:   clerkID,
	}

	returnData := config.DB.Create(&user)
	if returnData.Error != nil {
		return nil, &shared.AppError{
			Message:    returnData.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &user, nil
}

func Update(ID uuid.UUID, clerkID string, name string, avatarUrl string, email string) (*entity.User, *shared.AppError) {
	user := entity.User{
		ID:        ID,
		Name:      name,
		Email:     email,
		AvatarUrl: avatarUrl,
		ClerkID:   clerkID,
	}

	returnData := config.DB.Save(&user)
	if returnData.Error != nil {
		return nil, &shared.AppError{
			Message:    returnData.Error.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &user, nil
}
