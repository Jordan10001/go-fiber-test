package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configurations.
type Config struct {
	AppPort     string
	MongoDBURI  string
	MongoDBName string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	return &Config{
		AppPort:     os.Getenv("PORT"),
		MongoDBURI:  os.Getenv("MONGODB_URI"),
		MongoDBName: os.Getenv("MONGODB_NAME"),
	}
}