package main

import (
	"flag"

	"first_fiber"
	"first_fiber/databases"
	"first_fiber/library/custom_log"
	"first_fiber/library/utils/auth"
	"first_fiber/models/client"
)

func main() {
	err := first_fiber.LoadConf()
	if err != nil {
		custom_log.L.Fatalf("Could not load project confings. %s", err.Error())
	}

	db := databases.GetPostgres()

	var username string
	var password string
	flag.StringVar(&username, "u", "", "username of the instance")
	flag.StringVar(&password, "p", "", "new password")
	flag.Parse()

	if len(password) == 0 {
		custom_log.L.Fatal("password must not be empty")
	}

	instance := new(client.User)
	db.Model(client.User{}).Where("username = ?", username).Find(instance)
	if instance.Id == 0 {
		custom_log.L.Fatal("instance not found")
	}

	instance.Password, _ = auth.Hash(password)

	if err := db.Model(instance).Select("password", "updated_at").Save(instance).Error; err != nil {
		custom_log.L.Fatal(err.Error())
	}

	custom_log.L.Success("password updated")
}
