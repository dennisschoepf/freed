package main

import (
	"freed/internal/api"
	"freed/internal/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	dbFile = "./freed.db"
)

func main() {
	err := database.Connect(dbFile)

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(&fiber.Map{
			"application": "freed",
			"version":     "0.0.1",
			"status":      "running",
		})
	})

	api.Setup(app)

	app.Listen(":42069")
}
