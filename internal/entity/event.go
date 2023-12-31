package entity

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string         `json:"name" gorm:"not null"`
	Description string         `json:"description"`
	Runners     int            `json:"runners" gorm:"default:0"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
}
