package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Get Root docs
//	@Description	Root endpoint with ali ali ali
//	@Tags			default
//	@Success		200	{string}	string	"hello root!"
//	@Router			/ [get]
func Root(ctx *fiber.Ctx) error {
	return ctx.SendString("hello root!")
}
