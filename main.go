package main

import (
	_ "embed"
	"fmt"
	"freed/internal/database"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbFile := os.Getenv("DB_PATH")

	if dbFile == "" {
		log.Fatalf("No ENV value set for 'DB_PATH', could not initialize database. Please provide a valid path and filename")
	}

	_, err := database.Connect(dbFile)

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	fmt.Println("Setup finished")
	os.Exit(1)
}
