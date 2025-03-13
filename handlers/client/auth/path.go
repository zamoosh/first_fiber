package auth

import (
	"first_fiber/handlers"
	"first_fiber/handlers/agency"
	clientAgency "first_fiber/handlers/client/agency"
	"first_fiber/handlers/client/auth/middlewares"

	"github.com/gofiber/fiber/v2"
)

const (
	Path = handlers.Path + "api/admin/" + agency.AppName + "/"
)

func PreparePath(app *fiber.App) {
	app.Post("/verify", Verify)

	app.Get("/is-verify", middlewares.Verify, IsVerify)
	app.Get("/get-agency", middlewares.Verify, clientAgency.GetAgency)
}
