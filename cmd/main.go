package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Event struct {
	ID        	int64     `json:"id"`
	Name     	string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}



var events []Event

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/create", CreateEventHandler).Methods("POST")
	r.HandleFunc("/{id}", FindEventHandler).Methods("GET")

	http.Handle("/", r)

	port := ":3333"
	fmt.Printf("Server is listening on port %s...\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var event Event
	err = json.Unmarshal(body, &event)
	if err != nil {
		panic(err)
	}

	// add id and timestamps to event
	event.ID = int64(len(events) + 1)
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	events = append(events, event)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(events)
}

func FindEventHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		panic(err)
	}

	var event Event
	for _, e := range events {
		if e.ID == id {
			event = e
		}
	}

	if event.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"message": "Event not found"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)

}