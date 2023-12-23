package userUseCases

import (
	"net/http"
	"runners-api/internal/entity"
	userRepositories "runners-api/internal/models/users/repositories"
	"runners-api/internal/shared"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func Auth(name string, googleID int, avatarUrl string, email string) (*entity.UserTokens, *shared.AppError) {
	var user *entity.User
	var err *shared.AppError

	if googleID != 0 {
		user, err = userRepositories.FindByGoogleID(googleID)

		// error on searching the user
		if err != nil {
			return nil, err
		}

		// user with this google id not found
		if user == nil {
			userCreated, repoErr := userRepositories.Create(name, googleID, avatarUrl, email)

			if repoErr != nil {
				return nil, repoErr
			}

			user = userCreated
		}
	}

	/// TODO: Which field to check if not social login
	if user == nil {
		return nil, &shared.AppError{
			Message:    "treatment to check if not social login",
			StatusCode: http.StatusUnprocessableEntity,
		}
	}

	userTokens := &entity.UserTokens{}

	jwtExpiresInHours := viper.GetInt("JWT_EXPIRES_IN_HOURS")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.ID.String(),
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(jwtExpiresInHours)).Unix(),
	})

	jwtSecret := viper.GetString("JWT_SECRET")
	token, jwtErr := claims.SignedString([]byte(jwtSecret))
	if jwtErr != nil {
		return nil, &shared.AppError{
			Message:    "error generating the jwt token",
			StatusCode: http.StatusInternalServerError,
		}
	}

	userTokens.AccessToken = token

	return userTokens, nil
}
