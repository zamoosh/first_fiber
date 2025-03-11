package auth

import (
	"encoding/json"
	"fmt"

	"first_fiber/databases"
	"first_fiber/handlers"
	"first_fiber/library/utils/auth"
	"first_fiber/models/client"

	"github.com/charmbracelet/log"
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
	err := json.Unmarshal(c.Request().Body(), &data)
	if err != nil {
		return handlers.BadRequest(c, "دیتای ارسالی معتبر نیست")
	}

	var user client.User
	db := databases.GetPostgres()
	result := db.Where("username = ? OR cellphone = ?", data.Username, data.Username).First(&user)
	if result.Error != nil {
		log.Errorf("db error hapened. %s", result.Error)
	}

	if !auth.Compare(user.Password, data.Password) {
		return handlers.BadRequest(c, "نام کاربری یا رمز عبور اشتباه است")
	}

	t, _ := auth.GenerateToken(user.Id, auth.AccessToken)

	c.Response().Header.Set("Authorization", t)
	return c.Status(fiber.StatusOK).JSON(handlers.Msg{Msg: "شما با موفقیت وارد شدید"})
}

func IsVerify(c *fiber.Ctx) error {
	user := c.Locals("user").(client.User)
	fmt.Printf("User: %s\n", user)
	return c.SendString("client is verified")
}
