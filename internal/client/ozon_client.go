package client

import (
	"github.com/diphantxm/ozon-api-client/ozon" // Adjust this import path according to your project structure
)

var apiClient *ozon.Client // Use the Client structure from the ozon package

// InitializeClient initializes the Ozon client with provided credentials
func InitializeClient(apiKey, clientID string) {
	opts := []ozon.ClientOption{
		ozon.WithAPIKey(apiKey),
		ozon.WithClientId(clientID),
	}
	apiClient = ozon.NewClient(opts...)
}

// GetClient returns the initialized Ozon client
func GetClient() *ozon.Client {
	return apiClient
}
