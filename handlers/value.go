package handlers

import "github.com/gofiber/fiber/v2"

func Value(ctx *fiber.Ctx) error {
	return ctx.SendString("value is: " + ctx.Params("value"))
}
