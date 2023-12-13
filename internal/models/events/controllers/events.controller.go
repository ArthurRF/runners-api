package eventControllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"runners-api/internal/entity"
	eventUseCases "runners-api/internal/models/events/usecases"
	"strconv"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	events, err := eventUseCases.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetOneByID(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "eventID")

	convertedEventID, err := strconv.Atoi(eventID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	event, err := eventUseCases.GetOneByID(uint(convertedEventID))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	event := new(entity.Event)
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = eventUseCases.Create(event.Name, event.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	// id := c.Params("id")
	// parsedId, err := strconv.Atoi(id)

	// if err != nil {
	//     return err
	// }

	// event := new(entity.Event)

	// if err := c.BodyParser(event); err != nil {
	//     return err
	// }

	// eventUpdated, err := eventUseCases.Update(parsedId, event)

	// if err != nil {
	//     return err
	// }

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "eventID")

	convertedID, err := strconv.Atoi(eventID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = eventUseCases.Delete(uint(convertedID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
