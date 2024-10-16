package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	AppPort        string
	SecretPassword string
	SecretKey      string
	PostgresURL    string
}

func InitEnvironment() *Environment {
	godotenv.Load()

	return &Environment{
		AppPort:        os.Getenv("APP_PORT"),
		SecretPassword: os.Getenv("SECRET_PASSWORD"),
		SecretKey:      os.Getenv("SECRET_KEY"),
		PostgresURL:    os.Getenv("POSTGRES_URL"),
	}
}
