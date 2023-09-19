package main

import (
	"fmt"

	config "github.com/ArthurRF/runners-api/configs"
)

func main() {
	cfg, _ := config.LoadConfig(".")
	fmt.Println("main", cfg.DBDriver)
}
