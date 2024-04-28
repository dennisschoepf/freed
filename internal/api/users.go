package api

import (
	"github.com/gofiber/fiber/v2"
)

func FetchAllUsersHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(&fiber.Map{"users": "none"})
}
