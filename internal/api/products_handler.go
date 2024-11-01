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

// GetListOfProductsHandler обработчик для получения списка всех продуктов и записи в БД
func GetListOfProductsHandler(w http.ResponseWriter, r *http.Request) {
	apiClient := client.GetClient()
	if apiClient == nil {
		http.Error(w, "API client is not initialized", http.StatusInternalServerError)
		return
	}

	// Инициализируем параметры запроса с пустым LastId
	params := &ozon.GetListOfProductsParams{
		Filter: ozon.GetListOfProductsFilter{
			OfferId:    []string{},
			ProductId:  []int64{},
			Visibility: "ALL",
		},
		LastId: "",
		Limit:  1000,
	}

	// Контекст с тайм-аутом для запроса
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Переменная для отслеживания `LastId`
	var lastID string
	var totalProducts int32

	for {
		resp, err := apiClient.Products().GetListOfProducts(ctx, params)
		if err != nil {
			log.Println("Error fetching product list:", err)
			http.Error(w, "Error fetching product list", http.StatusInternalServerError)
			return
		}

		// Печатаем информацию о полученных продуктах
		PrintProductsInfo(resp)

		// Сохраняем каждый продукт в базе данных
		for _, product := range resp.Result.Items {
			err := SaveProductList(product.ProductId, product.OfferId, resp.Result.LastId, "GetListOfProducts")
			if err != nil {
				log.Printf("Error saving product %d to database: %v\n", product.ProductId, err)
			}
		}

		// Обновляем LastId для следующего запроса
		lastID = resp.Result.LastId

		// Проверка условий завершения: если достигли total или LastId пустой
		if lastID == "" {
			break
		}

		// Обновляем параметр LastId для следующей страницы
		params.LastId = lastID

		// Обновляем контекст с тайм-аутом для каждого запроса
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":       "All products successfully saved to the database",
		"totalProducts": fmt.Sprintf("%d", totalProducts),
	})
}
