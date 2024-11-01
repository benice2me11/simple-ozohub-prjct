package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dsn := os.Getenv("DATABASE_URL")

	for attempts := 1; attempts <= 5; attempts++ {
		DB, err = sql.Open("postgres", dsn) // Убедитесь, что "postgres" указан в sql.Open
		if err == nil && DB.Ping() == nil {
			log.Println("Successfully connected to the database")
			return nil
		}
		log.Printf("Attempt %d: Unable to connect to the database. Retrying in 2 seconds...\n", attempts)
		time.Sleep(2 * time.Second)
	}

	return err
}
