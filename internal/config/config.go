package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configurations.
type Config struct {
	AppPort         string
	PostgreSQLURI   string
	GoogleClientID  string
	GoogleSecret    string
	FrontendURL     string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	return &Config{
		AppPort:         os.Getenv("PORT"),
		PostgreSQLURI:   os.Getenv("POSTGRESQL_URI"),
		GoogleClientID:  os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleSecret:    os.Getenv("GOOGLE_CLIENT_SECRET"),
		FrontendURL:     os.Getenv("FRONTEND_REDIRECT_URL"),
	}
}