package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	AppPort           string
	SecretPassword    string
	SecretKey         string
	PostgresURL       string
	MaxOpenConnection int
	MaxIdleConnection int
}

func InitEnvironment() *Environment {
	godotenv.Load()

	maxOpenConnection, _ := strconv.Atoi(os.Getenv("MAX_OPEN_CONNECTION"))
	maxIdleConnection, _ := strconv.Atoi(os.Getenv("MAX_IDLE_CONNECTION"))

	return &Environment{
		AppPort:           os.Getenv("APP_PORT"),
		SecretPassword:    os.Getenv("SECRET_PASSWORD"),
		SecretKey:         os.Getenv("SECRET_KEY"),
		PostgresURL:       os.Getenv("POSTGRES_URL"),
		MaxOpenConnection: maxOpenConnection,
		MaxIdleConnection: maxIdleConnection,
	}
}
