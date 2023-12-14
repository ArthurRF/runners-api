package main

import (
	"fmt"
	"net/http"

	configs "runners-api/configs"

	"runners-api/internal/initializers"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"
)

func init() {
	configs.LoadConfig()
	configs.ConnectToDB()
}

func main() {
	r := chi.NewRouter()

	initializers.Routes(r)

	port := viper.Get("WEB_SERVER_PORT")

	fmt.Printf("Server starting on port :%s\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		panic(err)
	}
}
