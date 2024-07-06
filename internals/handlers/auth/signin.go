package handlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/database"
	"github.com/masterghost2002/videotube/internals/validations"
	"github.com/masterghost2002/videotube/types"
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
	user, err := database.Storage.GetUserByEmail(context.Background(), userData.Email)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return c.Status(404).JSON(fiber.Map{
			"error": "No user found",
		})
	}
	fmt.Println(user)
	isPasswordMatched := utils.ChechString(user.Password, userData.Password)
	if !isPasswordMatched {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"user": &types.UserResponse{
			ID:         user.ID,
			Email:      user.Email,
			Username:   user.Username,
			ChannelID:  user.ChannelID,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
			Profileurl: user.Profileurl,
		},
	})
}
