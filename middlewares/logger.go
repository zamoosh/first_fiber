package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func lo(c *fiber.Ctx) error {
	logger.New()
	return nil
}
