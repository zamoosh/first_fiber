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
		// Prefork:       true,
		EnablePrintRoutes: true,
		CaseSensitive:     true,
		StrictRouting:     true,
		AppName:           "first_fiber",
		ServerHeader:      "first_fiber",
		Immutable:         true,
	}

	items = []Item{
		{Id: 1, Name: "Apple", Price: 10.43},
		{Id: 2, Name: "Banana", Price: 9.5},
		{Id: 3, Name: "Orange", Price: 12},
	}
)

func getItems(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(items)
}

func getItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(Msg{"ali ali ali"})
	}

	for _, item := range items {
		if item.Id == id {
			return c.JSON(item)
		}
	}

	return c.Status(400).JSON(Msg{"Item didn't found"})
}

func createItem(c *fiber.Ctx) error {
	var item Item
	err := c.BodyParser(&item)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Msg{fmt.Sprintf("could not read. err: %s", err)})
	}

	log.Warn(item)
	item.Id = len(items) + 1

	items = append(items, item)
	return c.Status(fiber.StatusCreated).JSON(Msg{"Item created successfully!"})
}

func middleware(c *fiber.Ctx) error {
	return c.Next()
}

func getError(c *fiber.Ctx) error {
	// return fiber.NewError(fiber.StatusBadRequest, "ali ali ali")
	// c.Response().Header.Add("Content-Type", "application/json; charset=utf-8")
	return c.SendStatus(fiber.StatusFailedDependency)
}

func main() {
	_ = godotenv.Load()
	log.Info("ENV Loaded")

	app := fiber.New(config)
	app.Use("/api", middleware)
	app.Get("/api/item", getItems)
	app.Get("/api/item/:id<int>", getItem)
	app.Post("/api/item", createItem)
	app.Get("/api/error", getError)

	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Could not listen on port 3000. err: %s \n", err)
	}
}
