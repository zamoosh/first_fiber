package agency

import (
	"github.com/gofiber/fiber/v2"
)

func Retrieve(c *fiber.Ctx) error {
	return c.SendString("retrieving agency is completing...")
}
