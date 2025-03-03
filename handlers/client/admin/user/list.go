package user

import (
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_Name"`
}

var (
	users = []User{
		{Username: "zamoosh", FirstName: "Mohammad", LastName: "Rahimi"},
		{Username: "reza", FirstName: "reza", LastName: "Rezayi"},
		{Username: "ef", FirstName: "hossein", LastName: "fakhri"},
		{Username: "amin", FirstName: "amin", LastName: "kashani"},
		{Username: "mlk", FirstName: "reza", LastName: "Malek Zadeh"},
	}
)


func List(c *fiber.Ctx) error {
	return c.JSON(users)
}
