package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"simple-ozohub-prjct/internal/client"

	"github.com/diphantxm/ozon-api-client/ozon"
)

// PrintProductsInfo выводит отформатированный список продуктов
func PrintProductsInfo(response *ozon.GetListOfProductsResponse) {
	fmt.Println("Product List:")
	for _, item := range response.Result.Items {
		fmt.Printf("Product ID: %d, Offer ID: %s\n", item.ProductId, item.OfferId)
	}
	fmt.Printf("Total Products: %d\n", response.Result.Total)
	fmt.Printf("Last ID: %s\n", response.Result.LastId)
}

// GetListOfProductsHandler обновленный обработчик с визуализацией
func GetListOfProductsHandler(w http.ResponseWriter, r *http.Request) {
	apiClient := client.GetClient()
	if apiClient == nil {
		http.Error(w, "API client is not initialized", http.StatusInternalServerError)
		return
	}

	params := &ozon.GetListOfProductsParams{
		Filter: ozon.GetListOfProductsFilter{
			OfferId:    []string{},
			ProductId:  []int64{},
			Visibility: "ALL",
		},
		LastId: "",
		Limit:  10,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := apiClient.Products().GetListOfProducts(ctx, params)
	if err != nil {
		log.Printf("Error fetching product list: %v\n", err)
		http.Error(w, "Error fetching product list", http.StatusInternalServerError)
		return
	}

	PrintProductsInfo(resp)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}