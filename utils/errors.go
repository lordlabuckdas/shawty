package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type URLError struct {
	Code int
	Msg  string
}

func (err *URLError) Error() string {
	return fmt.Sprintf("%d - %s", err.Code, err.Msg)
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
