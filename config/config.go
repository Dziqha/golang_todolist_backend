package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func GetConfig() Configuration {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	conf := Configuration{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}

	return conf
}
