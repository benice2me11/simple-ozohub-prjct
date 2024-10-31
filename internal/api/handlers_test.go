package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	core "github.com/diphantxm/ozon-api-client"
	"github.com/diphantxm/ozon-api-client/ozon"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

// mockFetchProduct simulates the FetchProduct function for testing.
func mockFetchProduct(productID int64) (*ozon.GetProductDetailsResponse, error) {
	return &ozon.GetProductDetailsResponse{
		CommonResponse: core.CommonResponse{StatusCode: http.StatusOK},
		Result:         ozon.ProductDetails{ /* заполните нужными данными */ },
	}, nil
}

// TestGetProductHandler tests the GetProductHandler function with mockFetchProduct.
func TestGetProductHandler(t *testing.T) {
	// Подменим FetchProductFunc на mockFetchProduct
	originalFetchProduct := FetchProductFunc
	FetchProductFunc = mockFetchProduct
	defer func() { FetchProductFunc = originalFetchProduct }() // Восстановление после теста

	req := httptest.NewRequest("GET", "/products/123", nil)
	req = mux.SetURLVars(req, map[string]string{"product_id": "123"})
	w := httptest.NewRecorder()

	GetProductHandler(w, req)

	res := w.Result()
	defer res.Body.Close()

	var product ozon.GetProductDetailsResponse
	json.NewDecoder(res.Body).Decode(&product)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, http.StatusOK, product.CommonResponse.StatusCode) // Проверка StatusCode
}
