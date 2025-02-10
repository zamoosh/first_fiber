package handlers

import "github.com/gofiber/fiber/v2"

func Root(ctx *fiber.Ctx) error {
	return ctx.SendString("hello root!")
}
