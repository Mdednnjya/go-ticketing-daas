package config

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const dsn = "host=localhost port=5433 user=admin password=secret123 dbname=ticketing_db sslmode=disable"

func ConnectDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Error during connection: %s", err)
	}

	// db connection pooling layer 
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(30 * time.Minute)

	log.Printf("Successfully connected to PostgreSQL Database!\n")
	return db
}