package agency

import (
	"first_fiber/databases"
	"first_fiber/handlers"
	"first_fiber/handlers/agency/admin/agency/serializers"
	"first_fiber/models/agency"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func initialValidate() {
	if validate == nil {
		validate = validator.New(validator.WithRequiredStructEnabled())
	}
}

func Update(c *fiber.Ctx) error {
	initialValidate()
	db := databases.GetPostgres()

	id, err := c.ParamsInt("id")
	if err != nil {
		return handlers.BadRequest(c, "آی‌دی باید عددی باشد")
	}

	instance := new(agency.Agency)
	db.Model(agency.Agency{}).Where("id = ?", id).Preload("User").First(instance)
	if instance.Id == 0 {
		return handlers.BadRequest(c, "نماینده یافت نشد")
	}

	var body serializers.WriteAgencySerializerAdmin
	if err = c.BodyParser(&body); err != nil {
		return handlers.BadRequest(c, "آی‌دی باید عددی باشد")
	}

	err = validate.Struct(body)
	if err != nil {
		return handlers.BadRequest(c, "validation error")
	}

	instance.Name = body.Name
	if err = db.Model(instance).Select("name").Save(instance).Error; err != nil {
		return handlers.BadRequest(c, err.Error())
	}

	return c.JSON(handlers.Msg{Msg: "نمایندگی با موفقیت بروز شد"})
}
