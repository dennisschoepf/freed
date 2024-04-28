package api

import (
	"database/sql"
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func Setup(app *fiber.App, db *sql.DB) error {
	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return errors.New("Could not read API_KEY from ENV file.")
	}

	api := app.Group("/api", keyauth.New(keyauth.Config{
		SuccessHandler: successHandler,
		ErrorHandler:   errHandler,
		KeyLookup:      "header:x-api-key",
		ContextKey:     "apiKey",
		Validator:      apiKeyValidator(apiKey),
	}))

	apiHandler := NewHandler(db)

	v1 := api.Group("/v1")
	v1.Post("/users", apiHandler.createUser)

	return nil
}
