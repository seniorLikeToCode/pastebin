package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq" // assuming you are using PostgreSQL
	"github.com/seniorLikeToCode/pastebin/cmd/api"
)

func main() {
	// Database connection setup (example for PostgreSQL)
	connStr := "user=pastebinuser dbname=pastebin password=1234 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	// Ensure the database is available
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
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
