package config

import (
	"errors"
	"os"
)

// Config holds the application configuration
type Config struct {
	APIKey   string
	ClientID string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("API_KEY")
	clientID := os.Getenv("CLIENT_ID")

	if apiKey == "" || clientID == "" {
		return nil, errors.New("API_KEY and CLIENT_ID must be set")
	}

	return &Config{
		APIKey:   apiKey,
		ClientID: clientID,
	}, nil
}
