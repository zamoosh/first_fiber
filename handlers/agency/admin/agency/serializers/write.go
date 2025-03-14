package serializers

import (
	"first_fiber/databases"
	"first_fiber/handlers"
	"first_fiber/models/agency"
	"first_fiber/models/client"

	"github.com/gofiber/fiber/v2"
)

type WriteAgencySerializerAdminInterface interface {
	ValidateUserId(c *fiber.Ctx, instance agency.Agency, value interface{}) error
	ValidateName(c *fiber.Ctx, instance agency.Agency, value interface{}) error
	// Validate(c *fiber.Ctx) error

	Run()
}
type WriteAgencySerializerAdmin struct {
	// WriteAgencySerializerAdminInterface
	User uint64 `json:"user" validate:"required"`
	Name string `json:"name" validate:"required"`
}

func (WriteAgencySerializerAdmin) ValidateUserId(c *fiber.Ctx, instance agency.Agency, value interface{}) error {
	db := databases.GetPostgres()

	var exists bool
	_ = db.Model(client.User{}).Where("id = ?", uint64(value.(float64))).Find(&exists).Error
	if !exists {
		return handlers.BadRequest(c, "کاربر یافت نشد")
	}
	return nil
}

func (WriteAgencySerializerAdmin) ValidateName(c *fiber.Ctx, instance agency.Agency, value interface{}) error {
	return nil
}

func Run() {

}
