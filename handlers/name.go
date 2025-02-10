package handlers

import "github.com/gofiber/fiber/v2"

func Name(ctx *fiber.Ctx) error {
	return ctx.SendString("name is: " + ctx.Params("name"))
}
