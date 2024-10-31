package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-ozohub-prjct/internal/api"    // Adjust the import path according to your project structure
	"simple-ozohub-prjct/internal/client" // Adjust the import path according to your project structure
	"simple-ozohub-prjct/internal/config" // Adjust the import path according to your project structure

	"github.com/gorilla/mux" // Import Gorilla Mux for routing
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// Create a client with your Client-Id and Api-Key
	client.InitializeClient(cfg.APIKey, cfg.ClientID)

	// Set up the router
	r := mux.NewRouter()
	r.HandleFunc("/products/{product_id}", api.GetProductHandler).Methods("GET")
	r.HandleFunc("/products/list/", api.GetListOfProductsHandler).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r)) // Listen on port 8080

	if err := config.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to database successfully.")
	// остальная инициализация сервиса
}
