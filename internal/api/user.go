package api

import (
	"freed/internal/model"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) createUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return err
	}

	result, err := h.db.Exec("INSERT INTO user (first_name, email) VALUES (?, ?)", user.FirstName, user.Email)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	c.SendStatus(201)
	return c.JSON(&fiber.Map{"userId": id})
}
