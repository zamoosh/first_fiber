package agency

import (
	"first_fiber/databases"
	"first_fiber/models/agency"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Retrieve(c *fiber.Ctx) error {
	db := databases.GetPostgres()
	agencies := new([]agency.Agency)
	db.Where("deleted IS NULL").FindInBatches(agencies, 10, func(tx *gorm.DB, batch int) error { return nil })
	return c.JSON(agencies)
}
