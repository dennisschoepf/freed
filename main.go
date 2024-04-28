package main

import (
	_ "embed"
	"freed/internal/api"
	"freed/internal/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dbFile := os.Getenv("DB_FILE")

	if dbFile == "" {
		log.Fatalf("No ENV value set for 'DB_FILE', could not initialize database. Please provide a valid path and filename")
	}

	db, err := database.Connect(dbFile)

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	app := fiber.New()

	// Global Middlewares
	app.Use(logger.New())

	// Try to set up API routes
	if err := api.Setup(app, db); err != nil {
		log.Printf("Could not setup /api routes: %s", err)
	}

	app.Listen(":42069")
}
