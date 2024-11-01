package api

import (
	"log"
	"simple-ozohub-prjct/internal/config"
	"time"
)

// SaveProductDetails сохраняет или обновляет детали продукта в таблице products.
func SaveProductDetails(productID int64, offerID, name, price, oldPrice, currencyCode, primaryImage string, sku int64, updatedAt time.Time) error {
	query := `
        INSERT INTO products (product_id, offer_id, name, price, old_price, currency_code, primary_image, sku, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (product_id) DO UPDATE 
        SET offer_id = EXCLUDED.offer_id,
            name = EXCLUDED.name,
            price = EXCLUDED.price,
            old_price = EXCLUDED.old_price,
            currency_code = EXCLUDED.currency_code,
            primary_image = EXCLUDED.primary_image,
            sku = EXCLUDED.sku,
            updated_at = EXCLUDED.updated_at;
    `
	_, err := config.DB.Exec(query, productID, offerID, name, price, oldPrice, currencyCode, primaryImage, sku, updatedAt)
	if err != nil {
		log.Printf("Error saving product to database: %v\n", err)
		return err
	}
	log.Println("Product saved successfully!")
	return nil
}

// SaveProductList
func SaveProductList(productID int64, offerID, lastID, serviceSource string) error {
	query := `
        INSERT INTO products_list (product_id, offer_id, last_id, service_source)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (product_id) DO UPDATE 
        SET last_id = EXCLUDED.last_id, 
            offer_id = EXCLUDED.offer_id,
            service_source = EXCLUDED.service_source;
    `
	_, err := config.DB.Exec(query, productID, offerID, lastID, serviceSource)
	if err != nil {
		log.Printf("Error saving product to database: %v\n", err)
		return err
	}
	log.Println("Product saved successfully!")
	return nil
}
