package middlewares

import (
	"time"

	"first_fiber/databases"
	"first_fiber/handlers"
	"first_fiber/library/utils/auth"
	"first_fiber/models/client"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Verify(c *fiber.Ctx) error {
	// r := recover()
	// if r != nil {
	// 	log.Errorf("internal error. %s", r)
	// }

	jwtStr := string(c.Request().Header.Peek("Authorization"))

	token, err := auth.GetToken(jwtStr)

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

	data, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن معتبر نیست"})
	}

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
