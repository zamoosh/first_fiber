package main

import (
	"first_fiber"
	"first_fiber/databases/mongo"
	agencyAdmin "first_fiber/handlers/agency/admin"
	clientAuth "first_fiber/handlers/client/auth"
	"first_fiber/library/custom_log"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.mongodb.org/mongo-driver/v2/bson"
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

	c := mongo.CountDocuments(mongo.MCDColl(mongo.ActiveGpsLog), bson.D{})
	custom_log.L.Warnf("cound: %d", c)

	app := fiber.New(config)
	app.Use(logger.New(logger.Config{TimeZone: "Asia/Tehran"}))

	clientAuth.PreparePath(app)
	agencyAdmin.PreparePath(app)

	err = app.Listen("127.0.0.1:3000")
	if err != nil {
		log.Errorf("could not lister. %s", err)
	}
}
