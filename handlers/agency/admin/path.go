package admin

import (
	"fmt"

	"first_fiber/handlers"
	"first_fiber/handlers/agency"
	api "first_fiber/handlers/agency/admin/agency"
	"first_fiber/handlers/client/auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

const (
	Path = handlers.Path + "api/admin/" + agency.AppName + "/"
)

func PreparePath(app *fiber.App) {
	app.Use(middlewares.VerifyAndIsAdmin)
	app.Get(fmt.Sprintf("%s:id<int>", Path), api.Retrieve)
}

