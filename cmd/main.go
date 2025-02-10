package main

import (
	"log"
	"time"

	"first_fiber/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "first_fiber",
		ServerHeader:  "first_fiber",
	}

	app := fiber.New(config)

	// app.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.SendString("ali ali ali")
	// })

	app.Static(
		"/static",
		"./static",
		fiber.Static{
			Browse:        true,
			ByteRange:     true,
			CacheDuration: 10 * time.Second,
			MaxAge:        60,
		},
	)
	app.Get("/", handlers.Root)
	app.Get("/value/:value", handlers.Value)
	app.Get("/name/:name?", handlers.Name)
	app.Get(
		"/err", func(ctx *fiber.Ctx) error {
			return fiber.NewError(400, "some dummy error!")
		},
	)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalln("Could not listen on port 3000")
	}
}
