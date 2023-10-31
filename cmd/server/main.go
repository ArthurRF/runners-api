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
	e := echo.New()

	// Routes
	e.GET("/", hello)

	port := ":3333"

	if viper.Get("WEB_SERVER_PORT") != nil {
		port = fmt.Sprintf(":%v", viper.Get("WEB_SERVER_PORT"))
	}
	e.Logger.Fatal(e.Start(port))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
