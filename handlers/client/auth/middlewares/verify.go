package middlewares

import (
	"first_fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func Verify(c *fiber.Ctx) error {
	jwt := string(c.Request().Header.Peek("Authorization"))
	if jwt != "Bearer 1234567890" {
		return handlers.Unauthorized(c)
	}
	return c.Next()
}
