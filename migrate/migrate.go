package main

import (
	config "github.com/ArthurRF/runners-api/configs"
	"github.com/ArthurRF/runners-api/internal/entity"
)

func init() {
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&entity.User{})
	config.DB.AutoMigrate(&entity.Event{})
}
