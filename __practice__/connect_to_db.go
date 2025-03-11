package main

import (
	"fmt"

	"first_fiber"
	"first_fiber/databases"
	"first_fiber/library/utils"
	"first_fiber/models/client"

	"github.com/charmbracelet/log"
)

func main() {
	err := first_fiber.LoadConf()
	if err != nil {
		log.Fatalf("Could not load project confings. %s", err.Error())
	}

	db, _ := databases.GetPostgres()

	// var users []client.ClientUser
	// if err := db.Debug().Find(&users).Error; err != nil {
	// 	log.Errorf("Could not get users. %s", err)
	// }
	// for i := 0; i < len(users); i++ {
	// 	fmt.Println(users[i])
	// }

	var zamoosh client.User
	db.Where("username = ?", "09392511300").First(&zamoosh)
	fmt.Println(zamoosh)

	var count int64
	db.Table(client.User{}.TableName()).Count(&count)
	fmt.Println(count)

	fmt.Println(utils.Hash("66569211"))

}
