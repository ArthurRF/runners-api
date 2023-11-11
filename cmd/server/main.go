package main

import (
	"fmt"

	configs "runners-api/configs"

	"runners-api/internal/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	configs.LoadConfig()
	configs.ConnectToDB()
}

func main() {
	app := fiber.New()

	initializers.Routes(app)

	port := viper.Get("WEB_SERVER_PORT")

	app.Listen(fmt.Sprintf(":%s", port))
}
