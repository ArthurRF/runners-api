package entity

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInvalidEntity = errors.New("invalid entity")
)

type Event struct {
	gorm.Model
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (e *Event) Validate() error {
	if e.ID == 0 {
		return ErrInvalidEntity
	}

	return nil
}
