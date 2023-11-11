package config

import (
	"path/filepath"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASS"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	JwtExpiresIn  string `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuthKey  *jwtauth.JWTAuth
}

func LoadConfig() (*conf, error) {
	envFile := filepath.Join(".", ".env")

	viper.SetConfigType("env")
	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuthKey = jwtauth.New("HS256", []byte(cfg.JwtSecret), nil)

	return cfg, err
}
