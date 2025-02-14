package main

import (
	"log"

	"first_fiber/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config := fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "first_fiber",
		ServerHeader:  "first_fiber",
	}
	app := fiber.New(config)

	app.Get("/", handlers.Root)
	app.Get("/value/:value", handlers.Value)
	app.Get("/name/:name?", handlers.Name)
	app.Get(
		"/err/*", func(ctx *fiber.Ctx) error {
			return fiber.NewError(400, "some dummy error!")
		},
	)

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		log.Fatalln("Could not listen on port 8000", err)
	}
}
