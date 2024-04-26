package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func Setup(app *fiber.App) {
	api := app.Group("/api", keyauth.New(keyauth.Config{
		SuccessHandler: successHandler,
		ErrorHandler:   errHandler,
		KeyLookup:      "header:x-api-key",
		ContextKey:     "apiKey",
		Validator:      validateAPIKey,
	}))

	v1 := api.Group("/v1")

	v1.Get("/users", FetchAllUsersHandler)
}
