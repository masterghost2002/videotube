package handlers

import (
	"context"

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
		return err
	}
	user, err := database.Storage.GetUserByEmail(context.Background(), userData.Email)
	if err != nil {
		return err
	}
	isPasswordMatched := utils.ChechString(user.Password, userData.Password)
	if !isPasswordMatched {
		return c.Status(fiber.StatusUnauthorized).JSON(
			&types.Response{
				Error:       true,
				Message:     "Incorrect Password",
				ErrorFields: []string{"Password"},
				Data:        nil,
			},
		)
	}

	return c.Status(200).JSON(
		&types.Response{
			Error:       false,
			Message:     "Sign In  success",
			ErrorFields: nil,
			Data: &types.UserResponse{
				ID:         user.ID,
				Email:      user.Email,
				Username:   user.Username,
				ChannelID:  user.ChannelID,
				CreatedAt:  user.CreatedAt,
				UpdatedAt:  user.UpdatedAt,
				Profileurl: user.Profileurl,
			},
		},
	)
}
