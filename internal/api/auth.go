package api

import (
	"crypto/sha256"
	"crypto/subtle"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var (
	apiKey        = "willbechanged"
	protectedURLs = []*regexp.Regexp{
		regexp.MustCompile("^/api$"),
	}
	errForbidden = &fiber.Error{
		Code:    403,
		Message: "API Key is missing or invalid",
	}
)

func validateAPIKey(_ *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func protectedRoutesFilter(ctx *fiber.Ctx) bool {
	originalURL := strings.ToLower(ctx.OriginalURL())

	for _, pattern := range protectedURLs {
		if pattern.MatchString(originalURL) {
			return false
		}
	}
	return true
}

func successHandler(ctx *fiber.Ctx) error {
	return ctx.Next()
}

func errHandler(ctx *fiber.Ctx, err error) error {
	ctx.Status(fiber.StatusForbidden)

	return ctx.JSON(errForbidden)
}
