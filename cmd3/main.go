package main

import (
	"fmt"

	"first_fiber/handlers/client/admin/user"
	"first_fiber/handlers/client/auth"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	config = fiber.Config{
		CaseSensitive:     true,
		EnablePrintRoutes: true,
		AppName:           "first_fiber",
		ServerHeader:      "Server_1",
		Immutable:         false,
	}
)

func main() {
	_ = godotenv.Load()
	log.Info("ENV Loaded")

	app := fiber.New(config)

	userApi := app.Group(user.Path)
	userApi.Get("", user.List)

	authApi := app.Group(auth.Path)
	authApi.Post("", auth.Verify)

	fmt.Println(app.GetRoutes()[0].Path + "/")

	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Could not listen on port 3000. err: %s \n", err)
	}
}
