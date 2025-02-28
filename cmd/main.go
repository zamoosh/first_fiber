package main

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Item struct {
	Id    int     `json:"Id"`
	Name  string  `json:"name"`
	Price float64 `json:"Price"`
}

type Msg struct {
	Msg string `json:"msg"`
}

var (
	config = fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "first_fiber",
		ServerHeader:  "first_fiber",
		Immutable:     true,
	}

	items = []Item{
		{Id: 1, Name: "Apple", Price: 10.43},
		{Id: 2, Name: "Banana", Price: 9.5},
		{Id: 3, Name: "Orange", Price: 12},
	}
)

func getItems(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(items)
}

func getItem(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(Msg{"ali ali ali"})
	}

	for _, item := range items {
		if item.Id == id {
			return ctx.JSON(item)
		}
	}

	return ctx.Status(400).JSON(Msg{"Item didn't found"})
}

func createItem(ctx *fiber.Ctx) error {
	var item Item
	err := ctx.BodyParser(&item)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(Msg{fmt.Sprintf("could not read. err: %s", err)})
	}

	log.Warn(item)
	item.Id = len(items) + 1

	items = append(items, item)
	return ctx.Status(fiber.StatusCreated).JSON(Msg{"Item created successfully!"})
}

func middleware(ctx *fiber.Ctx) error {
	fmt.Println("ali ali ali", ctx)
	return ctx.Next()
}

func main() {
	_ = godotenv.Load()
	log.Info("ENV Loaded")

	app := fiber.New(config)
	app.Get("/api/item", getItems)
	app.Get("/api/item/:id<int>", getItem)
	app.Post("/api/item", createItem)
	app.Use("/api", middleware)

	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Could not listen on port 3000. err: %s \n", err)
	}
}
