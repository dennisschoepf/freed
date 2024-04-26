package api

import (
	"errors"
	"freed/database"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

type Handlers struct {
	Repo *database.Repository
}

func NewHandlers(repo *database.Repository) *Handlers {
	return &Handlers{Repo: repo}
}

func Setup(app *fiber.App, repository *database.Repository) error {
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

	handlers := NewHandlers(repository)

	v1 := api.Group("/v1")
	v1.Get("/users", handlers.FetchAllUsersHandler)

	return nil
}

func (h *Handlers) FetchAllUsersHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(&fiber.Map{"users": "none"})
}
