package handlers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/models"
	repository "github.com/masterghost2002/videotube/internals/repository/database"
	jwt "github.com/masterghost2002/videotube/internals/token"
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
	token, err := jwt.GenerateJWT(jwt.UserPayload{
		FullName: userData.FullName,
		Email:    userData.Email,
	})
	if err != nil {
		c.SendStatus(201)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(124 * time.Hour),
		HTTPOnly: true,
	})
	return c.SendStatus(201)
}
