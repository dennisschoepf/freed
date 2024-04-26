package main

import (
	"freed/internal/database"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_, err := database.New("./freed.db")

	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	app := fiber.New()

	// Routes
	app.Get("/status", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"application": "freed",
			"version":     "0.0.1",
			"status":      "up",
		})
		return c.SendStatus(http.StatusOK)
	})

	app.Listen(":42069")
}
