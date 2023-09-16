package entity

import (
	"errors"
	"time"
)

var (
	ErrInvalidEntity = errors.New("Invalid entity")
)

type Event struct {
	ID        	int64     `json:"id"`
	Name     	string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (e *Event) Validate () error {
	if e.ID == 0 {
		return ErrInvalidEntity
	}

	return nil
}
