package main

import (
	"fmt"

	config "github.com/ArthurRF/runners-api/configs"
	"github.com/ArthurRF/runners-api/internal/entity"
)

func init() {
	config.ConnectToDB()
}

func main() {
	user, err := entity.NewUser("Alexandre", "emailteste", "abcdefg")
	if err != nil {
		fmt.Println(err)
	}

	db := config.DB
	pk := db.Create(&user)

	fmt.Println(pk)
}
