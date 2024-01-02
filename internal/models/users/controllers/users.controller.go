package userControllers

import (
	"encoding/json"
	"net/http"
	"runners-api/internal/entity"
	userUseCases "runners-api/internal/models/users/usecases"
)

func Upsert(w http.ResponseWriter, r *http.Request) {
	user := new(entity.User)
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the property 'name' is required."))
		return
	}

	if user.ClerkID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("the property 'clerkId' is required."))
		return
	}

	appErr := userUseCases.Upsert(user.ClerkID, user.Name, user.AvatarUrl, user.Email)
	if appErr != nil {
		w.WriteHeader(int(appErr.StatusCode))
		w.Write([]byte(appErr.Message))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
