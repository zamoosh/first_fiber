package middlewares

import (
	"fmt"
	"reflect"
	"time"

	"first_fiber/databases"
	"first_fiber/handlers"
	"first_fiber/library/utils/auth"
	"first_fiber/models/client"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func Verify(c *fiber.Ctx) error {
	// r := recover()
	// if r != nil {
	// 	log.Errorf("internal error. %s", r)
	// }

	jwt := string(c.Request().Header.Peek("Authorization"))

	token, err := auth.GetToken(jwt)

	if token == nil || err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن معتبر نیست"})
	}

	exp, err := token.Claims.GetExpirationTime()
	if exp == nil || err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن معتبر نیست"})
	}

	if time.Now().UTC().After(exp.Time) {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن شما باطل شده است"})
	}

	claimValue := reflect.ValueOf(token.Claims)
	if claimValue.Kind() != reflect.Map {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن معتبر نیست"})
	}

	data := make(map[string]any)
	iter := claimValue.MapRange()
	for iter.Next() {
		data[iter.Key().String()] = iter.Value().Interface()
	}
	fmt.Println(data)

	if userId, ok := data["user_id"]; !ok || userId == 0 {
		log.Warn("token has been manipulated")
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن معتبر نیست"})
	}


	var user client.ClientUser
	db, _ := databases.GetPostgres()
	db.Where("id = ?", data["user_id"]).First(&user)
	c.Locals("user", user)
	return c.Next()
}
