package config

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const dsn = "host=localhost port=5433 user=admin password=secret123 dbname=ticketing_db sslmode=disable"

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatalf("Error during connection: %s", err)
	}

	log.Printf("Successfully connected to PostgreSQL Database!\n")
	return db
}