package auth

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)



type VerifySerializer struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	Path string = "/api/client/auth/verify/"
)

func Verify(c *fiber.Ctx) error {
	var data VerifySerializer
	_ = json.Unmarshal(c.Request().Body(), &data)
	if data.Username == "zamoosh" && data.Password == "66569211" {
		return c.Status(fiber.StatusOK).SendString("Authorized Successfully")
	}
	return c.Status(fiber.StatusBadRequest).SendString("Username or Password is wrong")
}
