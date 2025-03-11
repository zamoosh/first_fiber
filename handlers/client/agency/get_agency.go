package agency

import (
	"first_fiber/databases"
	"first_fiber/models/client"

	"github.com/gofiber/fiber/v2"
)

func GetAgency(c *fiber.Ctx) error {
	user := c.Locals("user").(client.User)
	db := databases.GetPostgres()
	db.Model(client.User{}).Preload("AgencySet", "deleted IS NULL").Where("id = ?", user.Id).Find(&user)
	return c.JSON(user.AgencySet)
}
