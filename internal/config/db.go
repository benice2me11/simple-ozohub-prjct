package config

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	return DB.Ping()
}
