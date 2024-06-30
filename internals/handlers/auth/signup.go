package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/models"
	repository "github.com/masterghost2002/videotube/internals/repository/database"
	"github.com/masterghost2002/videotube/internals/validations"
	"github.com/masterghost2002/videotube/utils"
)

func SignUp(c *fiber.Ctx) error {
	var userData validations.UserSignUpFormData
	if err := c.BodyParser(&userData); err != nil {
		return err
	}
	hashPassword := utils.HashString(userData.Password)
	user := models.User{FullName: userData.FullName, Email: userData.Email, Username: userData.Username, Password: hashPassword}

	if err := validations.Validate.Struct(userData); err != nil {
		var errors []validations.ErrorFields
		for _, validationError := range err.(validator.ValidationErrors) {
			errors = append(errors, validations.ErrorFields{
				Field:   validationError.Field(),
				Message: validationError.Tag(),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}
	if err := repository.CreateUser(user); err != nil {
		return c.Status(424).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.SendStatus(201)
}
