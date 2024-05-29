package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // assuming you are using PostgreSQL
	"github.com/seniorLikeToCode/pastebin/cmd/api"
)

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Ensure the database is available
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	// create a table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS data (
		id SERIAL PRIMARY KEY,
		uid TEXT NOT NULL,
		content TEXT NOT NULL
	)`)

	if err != nil {
		log.Fatalf("Error creating the table: %v", err)
	}

	// Get the server address from environment variable or default
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":5000"
	}

	server := api.NewAPIServer(addr, db)
	if err := server.Run(); err != nil {
		log.Fatalf("Error running the server: %v", err)
	}
}
