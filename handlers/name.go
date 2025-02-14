package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type response struct {
	Name string `json:"name"`
}

//	@Summary		Get Name
//	@Description	Get name by optional parameter
//	@Tags			default
//	@Param			name	path		string	false	"Name"
//	@Success		200		{string}	string	"name is: <name>"
func Name(ctx *fiber.Ctx) error {
	j := response{
		Name: "",
	}

	if ctx.Params("name") != "" {
		j.Name = ctx.Params("name")
	}

	fmt.Println(j)
	return ctx.JSON(j)
}
