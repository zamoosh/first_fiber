package handlers

import "github.com/gofiber/fiber/v2"

type Msg struct {
	Msg string `json:"msg"`
}

func BadRequest(c *fiber.Ctx, msg string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Msg{Msg: msg})
}

func Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Msg{Msg: "Unauthorized"})
}
