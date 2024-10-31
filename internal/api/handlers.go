package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"simple-ozohub-prjct/internal/client"

	"github.com/diphantxm/ozon-api-client/ozon"
	"github.com/gorilla/mux"
)

// CustomErrorResponse represents a standard format for error responses
type CustomErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// sendErrorResponse writes a JSON error response with a specified HTTP status code and message.
func sendErrorResponse(w http.ResponseWriter, statusCode int, errMessage, logMessage string) {
	log.Printf("Error: %s", logMessage)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(CustomErrorResponse{Error: http.StatusText(statusCode), Message: errMessage})
}

// FetchProduct retrieves product details from the marketplace API by product_id.
// Returns product details or an error if the retrieval fails.
func FetchProduct(productID int64) (*ozon.GetProductDetailsResponse, error) {
	ctx := context.Background()
	apiClient := client.GetClient()

	if apiClient == nil {
		return nil, fmt.Errorf("API client is not initialized")
	}

	resp, err := apiClient.Products().GetProductDetails(ctx, &ozon.GetProductDetailsParams{ProductId: productID})
	if err != nil {
		return nil, fmt.Errorf("error fetching product details from API: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from API: %d", resp.StatusCode)
	}
	return resp, nil
}

// GetProductHandler is an HTTP handler that fetches product details based on a product_id from the URL.
// Responds with JSON-formatted product details or an error message.
var FetchProductFunc = FetchProduct

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr := vars["product_id"]

	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "Invalid product ID format", fmt.Sprintf("invalid product ID: %s", productIDStr))
		return
	}

	product, err := FetchProductFunc(productID)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Error retrieving product details", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
