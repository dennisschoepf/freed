package api

import (
	"database/sql"
	"errors"
	"fmt"
	"freed/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
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
		return fiber.NewError(fiber.ErrInternalServerError.Code, validationErr.Message)
	}

	_, insertErr := h.db.Exec("INSERT INTO user (id, first_name, email) VALUES (?, ?, ?)", user.ID, user.FirstName, user.Email)

	if sqliteErr, ok := insertErr.(sqlite3.Error); ok {
		log.Warn(insertErr)

		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fiber.NewError(fiber.StatusBadRequest, "User with email already exists")
		}

		return defaultUserError
	}

	c.SendStatus(201)
	return c.JSON(&fiber.Map{"userId": userId})
}

func (h *Handler) getUserByID(id string) (*model.User, error) {
	var user *model.User

	if err := h.db.QueryRow("SELECT * FROM user WHERE id = ?", id).Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			notFoundErrMessage := fmt.Sprintf("No user found with ID: %s", id)
			return nil, errors.New(notFoundErrMessage)
		}

		unexpectedErrMessage := fmt.Sprintf("Unexpected error occured, could not get user with ID: %s", id)
		return nil, errors.New(unexpectedErrMessage)
	}

	return user, nil
}

func (h *Handler) userExists(id string) bool {
	var userID string

	if err := h.db.QueryRow("SELECT id FROM user WHERE id = ?", id).Scan(&userID); err != nil {
		return false
	}

	return true
}
