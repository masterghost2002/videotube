package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/validations"
	"github.com/masterghost2002/videotube/utils"
)

func SignIn(c *fiber.Ctx) error {
	var userData validations.UserSignInFormData
	if err := c.BodyParser(&userData); err != nil {
		return err
	}
	if err := validations.Validate.Struct(userData); err != nil {
		errors := utils.FormatValidationErrors(err)
		return c.Status(424).JSON(fiber.Map{
			"eror": errors,
		})
	}
	return c.SendStatus(200)
}
