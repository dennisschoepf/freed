package api

import (
	"freed/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

var defaultUserError = fiber.NewError(fiber.ErrInternalServerError.Code, "Could not create user")

func (h *Handler) createUser(c *fiber.Ctx) error {
	user := new(model.User)

	userId, idErr := uuid.NewRandom()

	if idErr != nil {
		log.Warn(idErr)
		return defaultUserError
	}

	user.ID = userId.String()

	if parseErr := c.BodyParser(user); parseErr != nil {
		log.Warn(parseErr)
		return defaultUserError
	}

	validationErr := ValidateModel(user)

	if validationErr != nil {
		return validationErr
	}

	_, insertErr := h.db.Exec("INSERT INTO user (id, first_name, email) VALUES (?, ?, ?)", user.ID, user.FirstName, user.Email)

	if insertErr != nil {
		log.Warn(insertErr)
		return defaultUserError
	}

	c.SendStatus(201)
	return c.JSON(&fiber.Map{"userId": userId})
}
