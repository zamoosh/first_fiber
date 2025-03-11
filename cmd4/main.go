package main

import (
	"first_fiber"
	agencyAdmin "first_fiber/handlers/agency/admin"
	clientAgency "first_fiber/handlers/client/agency"
	"first_fiber/handlers/client/auth"
	"first_fiber/handlers/client/auth/middlewares"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	config = fiber.Config{
		// Prefork:       true,
		EnablePrintRoutes: true,
		CaseSensitive:     true,
		AppName:           "first_fiber",
		ServerHeader:      "first_fiber",
		Immutable:         true,
	}
)

func main() {
	err := first_fiber.LoadConf()
	if err != nil {
		log.Fatalf("Could not load project confings. %s", err.Error())
	}

	app := fiber.New(config)
	app.Use(logger.New(logger.Config{TimeZone: "Asia/Tehran"}))
	app.Post("/verify", auth.Verify)

	app.Get("/is-verify", middlewares.Verify, auth.IsVerify)
	app.Get("/get-agency", middlewares.Verify, clientAgency.GetAgency)

	agencyAdmin.PreparePath(app)

	err = app.Listen("127.0.0.1:3000")
	if err != nil {
		log.Errorf("could not lister. %s", err)
	}
}
