package middlewares

import (
	"encoding/json"

	"first_fiber/databases"
	"first_fiber/handlers"
	"first_fiber/library/utils/auth"
	"first_fiber/models/client"

	"github.com/charmbracelet/log"
	"github.com/gofiber/fiber/v2"
)

func Verify(c *fiber.Ctx) error {
	jwt := string(c.Request().Header.Peek("Authorization"))

	token, err := auth.GetToken(jwt)

	if token == nil || err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن معتبر نیست"})
	}

	var data map[string]any
	jsonData, _ := json.Marshal(token.Claims)
	_ = json.Unmarshal(jsonData, &data)

	if userId, ok := data["user_id"]; !ok || userId == 0 {
		log.Warn("token has been manipulated")
		return c.Status(fiber.StatusUnauthorized).JSON(handlers.Msg{Msg: "توکن معتبر نیست"})
	}

	var claim auth.JwtClaim
	_ = json.Unmarshal(jsonData, &claim)

	var user client.ClientUser
	db, _ := databases.GetPostgres()
	db.Where("id = ?", claim.GetUserId()).First(&user)
	c.Locals("user", user)
	return c.Next()
}
