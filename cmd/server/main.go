package main

import (
	"fmt"

	"github.com/ArthurRF/runners-api/internal/entity"
)

func main() {
	user, err := entity.NewUser("Alexandre", "emailteste", "abcdefg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("main", user)

	isValid := user.ValidatePassword("abcdefg")
	fmt.Println("main", isValid)
}
