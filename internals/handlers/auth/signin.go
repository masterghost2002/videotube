package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/models"
	repository "github.com/masterghost2002/videotube/internals/repository/database"
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
	user := repository.FindUser(userData.Email)
	if user == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	isPasswordMatched := utils.ChechString(user.Password, userData.Password)
	if !isPasswordMatched {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	userResponse := models.UserResponse{
		FullName:  user.FullName,
		Email:     user.Email,
		ID:        user.ID,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}
	return c.Status(200).JSON(fiber.Map{
		"user": userResponse,
	})
}
