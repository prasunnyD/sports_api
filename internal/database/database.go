package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/marcboeker/go-duckdb"
)

// InitDB initializes the connection to MotherDuck
func InitDB() (*sql.DB, error) {
	// Get MotherDuck token from environment
	token := os.Getenv("MOTHERDUCK_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("MOTHERDUCK_TOKEN environment variable is required")
	}

	// Create connection string for MotherDuck
	connStr := fmt.Sprintf("md:?motherduck_token=%s", token)
	
	db, err := sql.Open("duckdb", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to MotherDuck")
	return db, nil
} 