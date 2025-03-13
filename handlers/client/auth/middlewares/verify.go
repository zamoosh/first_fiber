package middlewares

import (
	"errors"
	"time"

	"first_fiber/databases"
	"first_fiber/handlers"
	"first_fiber/library/utils/auth"
	"first_fiber/models/client"

	"github.com/gofiber/fiber/v2"
)

func checkToken(jwtStr string) (*auth.JwtClaim, error) {
	token, err := auth.GetToken(jwtStr)

	if token == nil || err != nil || !token.Valid {
		return nil, errors.New("توکن معتبر نیست")
	}

	exp, err := token.Claims.GetExpirationTime()
	if exp == nil || err != nil {
		return nil, errors.New("توکن معتبر نیست")
	}

	if time.Now().UTC().After(exp.Time) {
		return nil, errors.New("توکن شما باطل شده است")
	}

	return auth.ToJwtClaim(token.Claims)
}

func Verify(c *fiber.Ctx) error {
	jwtStr := string(c.Request().Header.Peek("Authorization"))
	jwtClaim, err := checkToken(jwtStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: err.Error()})
	}

	var user client.User
	db := databases.GetPostgres()
	db.Where("id = ?", jwtClaim.UserId).First(&user)
	c.Locals("user", user)
	return c.Next()
}

func VerifyAndIsAdmin(c *fiber.Ctx) error {
	jwtStr := string(c.Request().Header.Peek("Authorization"))
	jwtClaim, err := checkToken(jwtStr)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: err.Error()})
	}

	var user client.User
	db := databases.GetPostgres()
	db.Where("id = ?", jwtClaim.UserId).First(&user)
	c.Locals("user", user)
	return c.Next()
}
