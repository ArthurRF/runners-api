package main

import (
	config "runners-api/configs"
	"runners-api/internal/entity"
)

func init() {
	config.LoadConfig()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&entity.User{})
	config.DB.AutoMigrate(&entity.Event{})
}
