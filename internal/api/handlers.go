package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"marketplace-api-client/internal/client" // Adjust this import path according to your project structure

	"github.com/diphantxm/ozon-api-client/ozon"
	"github.com/gorilla/mux"
)

// FetchProduct retrieves product details from the marketplace API by product_id.
func FetchProduct(productID int64) (*ozon.GetProductDetailsResponse, error) {
	ctx := context.Background()
	apiClient := client.GetClient()

	// Send request with parameters to get product details
	resp, err := apiClient.Products().GetProductDetails(ctx, &ozon.GetProductDetailsParams{
		ProductId: productID,
	})

	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error when getting product details: %s", err)
	}

	return resp, nil
}

// GetProductHandler handles the GET request to fetch a product.
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["product_id"]

	// Convert productID to int64
	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		http.Error(w, "invalid product_id", http.StatusBadRequest)
		return
	}

	product, err := FetchProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Output the product details to the console
	fmt.Printf("Product ID: %d, Name: %s, Price: %s\n", product.Result.Id, product.Result.Name, product.Result.Price)

	// Return the product as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product.Result) // Return only the result part of the response
}
