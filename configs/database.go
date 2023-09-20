package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := "host=berry.db.elephantsql.com user=vvywiiyq password=sxy9VQsOBHsdE1LJJ6zeoZcdlbj56djc dbname=vvywiiyq port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection Opened to Database")
}
