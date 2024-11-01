package api

import (
	"log"
	"simple-ozohub-prjct/internal/config"
)

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
