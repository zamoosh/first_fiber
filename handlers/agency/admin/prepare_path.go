package admin

import (
	"fmt"

	api "first_fiber/handlers/agency/admin/agency"
	"first_fiber/handlers/client/auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

func PreparePath(app *fiber.App) {
	app.Use(middlewares.Verify)
	app.Get(fmt.Sprintf("%s:id<int>", Path), api.Retrieve)
}
