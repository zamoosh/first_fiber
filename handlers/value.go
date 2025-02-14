package handlers

import "github.com/gofiber/fiber/v2"

//	@Summary		Get Value
//	@Description	Get value by parameter
//	@Tags			default
//	@Param			value	path		string	true	"Value"
//	@Success		200		{string}	string	"value is: <value>"
//	@Router			/value/{value} [get]
func Value(ctx *fiber.Ctx) error {
	return ctx.SendString("value is: " + ctx.Params("value"))
}
