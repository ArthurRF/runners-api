package config

import (
	"fmt"

	"database/sql"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var DB *sql.DB

func ConnectToDB() {
	var err error
	host := viper.Get("DB_HOST")
	user := viper.Get("DB_USER")
	password := viper.Get("DB_PASS")
	dbname := viper.Get("DB_NAME")
	port := viper.Get("DB_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("The database is connected")
}
