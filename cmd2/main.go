package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

const GoRoutineId = "ABCDEFGHIJKLMNOPQRSTUYWXYZ"

var idx = 0

func getHandlerId() string {
	c := GoRoutineId[idx%26]
	idx++
	return fmt.Sprintf("id-%v-%c", idx, c)
}

var (
	config = fiber.Config{
		CaseSensitive:     true,
		StrictRouting:     true,
		EnablePrintRoutes: true,
		AppName:           "first_fiber",
		ServerHeader:      "Server_1",
		Immutable:         false,
	}
)

func getName(ctx *fiber.Ctx) error {
	n := ctx.Params("name")

	fmt.Printf("%T %v %s", n, n, n)

	name := make([]byte, len(n))
	copy(name, n)

	id := getHandlerId()

	go func() {
		t := time.After(10 * time.Second)
		for {
			select {
			case <-t:
				log.Info("handler done", "id", id, "name", string(name))
				return
			default:
				log.Info("still running", "id", id, "name", string(name))
				time.Sleep(time.Second)
			}
		}
	}()

	log.Info("request received", "name", n)
	return nil
}

func main() {
	_ = godotenv.Load()
	log.Info("ENV Loaded")

	app := fiber.New(config)
	app.Get("/:name", getName).Name("name")

	err := app.Listen("0.0.0.0:3000")
	if err != nil {
		log.Fatalf("Could not listen on port 3000. err: %s \n", err)
	}
}
