package api

import (
	"fmt"
	"freed/internal/model"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ModelConstraint interface {
	*model.User | *model.Feed
}

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateModel[T ModelConstraint](s T) *fiber.Error {
	var errorMessages []string
	validate := validator.New()
	err := validate.Struct(s)

	fmt.Printf("%#v", err)

	if err == nil {
		return nil
	}

	for _, err := range err.(validator.ValidationErrors) {
		err := fmt.Sprintf("Field %s is invalid, reason: %s", err.StructNamespace(), err.Tag())

		errorMessages = append(errorMessages, err)
	}

	return fiber.NewError(fiber.StatusBadRequest, strings.Join(errorMessages, ". "))
}
