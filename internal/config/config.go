package config

import (
	"log"
	"os"

	"github.com/joho/godotenv" // Import godotenv for loading env variables
)

// Config holds the application configuration
type Config struct {
	APIKey   string
	ClientID string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	apiKey := os.Getenv("API_KEY")
	clientID := os.Getenv("CLIENT_ID")

	if apiKey == "" || clientID == "" {
		return nil, log.Output(2, "API_KEY and CLIENT_ID must be set")
	}

	return &Config{
		APIKey:   apiKey,
		ClientID: clientID,
	}, nil
}
