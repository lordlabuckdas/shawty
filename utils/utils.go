package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type URLError struct {
	Code int
	Msg string
}

func (err *URLError) Error() string {
	return fmt.Sprintf("%d - %s", err.Code, err.Msg)
}

func ShortenURL(longURL string) (string, error) {
	if len(longURL) == 0 {
		return "", &URLError{400, "Invalid long URL - Length is 0"}
	}
	hash := sha256.Sum256([]byte(longURL))
	shortURL := hex.EncodeToString(hash[:])[:8]
	return shortURL, nil
}

func HandleError(c *fiber.Ctx, err error) error {
	if err, ok := err.(*URLError); ok {
		return c.Status(err.Code).Render("index", fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(500).Render("index", fiber.Map{
		"error": "500 - Internal Error",
	})
}
