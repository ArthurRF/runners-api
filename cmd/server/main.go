package main

import (
	"fmt"
	"net/http"

	config "github.com/ArthurRF/runners-api/configs"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	config.ConnectToDB()
	config.LoadConfig(".env")
}

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", viper.Get("WEB_SERVER_PORT"))))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
