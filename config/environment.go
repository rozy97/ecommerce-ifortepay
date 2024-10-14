package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	AppPort        string
	SecretPassword string
	MysqlURI       string
}

func InitEnvironment() Environment {
	godotenv.Load()

	return Environment{
		AppPort:        os.Getenv("APP_PORT"),
		SecretPassword: os.Getenv("SECRET_PASSWORD"),
		MysqlURI:       os.Getenv("MYSQL_URI"),
	}
}
