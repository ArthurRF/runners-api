package userUseCases

import (
	"net/http"
	"runners-api/internal/entity"
	userRepositories "runners-api/internal/models/users/repositories"
	"runners-api/internal/shared"
)

func Upsert(clerkId string, name string, avatarUrl string, email string) *shared.AppError {
	var user *entity.User
	var err *shared.AppError

	if clerkId != "" {
		user, err = userRepositories.FindByClerkID(clerkId)

		// error on searching the user
		if err != nil {
			return err
		}

		if user != nil {
			userUpdated, repoErr := userRepositories.Update(user.ID, clerkId, name, avatarUrl, email)

			if repoErr != nil {
				return repoErr
			}

			user = userUpdated
		}

		// user with this clerk id not found
		if user == nil {
			userCreated, repoErr := userRepositories.Create(clerkId, name, avatarUrl, email)

			if repoErr != nil {
				return repoErr
			}

			user = userCreated
		}
	}

	/// TODO: Which field to check if not social login
	if user == nil {
		return &shared.AppError{
			Message:    "error finding/creating an user",
			StatusCode: http.StatusInternalServerError,
		}
	}

	return nil
}
