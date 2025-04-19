package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	SecretKey        string
	Host             string
	Port             string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresHost     string
	PostgresPort     string
}

func Get() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}
	return Config{
		SecretKey:        os.Getenv("SECRET_KEY"),
		Host:             os.Getenv("HOST"),
		Port:             os.Getenv("PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDatabase: os.Getenv("POSTGRES_DB"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
	}
}
