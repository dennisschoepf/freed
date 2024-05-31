package api

import (
	"fmt"
	"freed/internal/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/mattn/go-sqlite3"
)

func (h *Handler) createFeed(c *fiber.Ctx) error {
	userID := c.Params("userID")
	feed := new(model.Feed)

	if parseErr := c.BodyParser(feed); parseErr != nil {
		log.Warn(parseErr)
		return defaultUserError
	}

	validationErr := ValidateModel(feed)

	if validationErr != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, validationErr.Message)
	}

	userExists := h.userExists(userID)
	fmt.Printf("exists? %v", userExists)

	if h.userExists(userID) == false {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "No existing user found, check user id path param.")
	}

	// TODO: Get name from feed
	// TODO: Get type from feed
	// TODO: Either schedule or request items from feed immediately

	sqlResult, insertErr := h.db.Exec("INSERT INTO feed (userId, name, url) VALUES (?, ?, ?)", userID, feed.Url, feed.Url)

	if sqliteErr, ok := insertErr.(sqlite3.Error); ok {
		log.Warn(insertErr)

		if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fiber.NewError(fiber.StatusBadRequest, "Feed with url already exists")
		}

		return fiber.NewError(fiber.ErrInternalServerError.Code, "Could not create feed")
	}

	feedID, idErr := sqlResult.LastInsertId()

	if idErr != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code, "Created feed, but could not retrieve its ID, check database.")
	}

	c.SendStatus(201)
	return c.JSON(&fiber.Map{"id": feedID})
}
