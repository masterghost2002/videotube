package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/database"
	jwt "github.com/masterghost2002/videotube/internals/token"
	"github.com/masterghost2002/videotube/internals/validations"
	"github.com/masterghost2002/videotube/types"
	"github.com/masterghost2002/videotube/utils"
)

func SignUp(c *fiber.Ctx) error {
	var userData validations.UserSignUpFormData
	if err := c.BodyParser(&userData); err != nil {
		return err
	}

	if err := validations.Validate.Struct(userData); err != nil {
		return err
	}
	hashPassword := utils.HashString(userData.Password)
	userParams := database.CreateUserParams{FullName: userData.FullName, Email: userData.Email, Username: userData.Username, Password: hashPassword}
	result, err := database.Storage.CreateUser(context.Background(), userParams)
	if err != nil {
		return err
	}
	token, err := jwt.GenerateJWT(jwt.UserPayload{
		FullName: result.FullName,
		Email:    result.Email,
	})
	if err != nil {
		c.SendStatus(500)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(124 * time.Hour),
		HTTPOnly: true,
	})
	return c.Status(201).JSON(fiber.Map{
		"user": &types.UserResponse{
			FullName:   result.FullName,
			Email:      result.Email,
			ChannelID:  result.ChannelID,
			Profileurl: result.Profileurl,
			Username:   result.Username,
			CreatedAt:  result.CreatedAt,
			UpdatedAt:  result.UpdatedAt,
		},
	})
}
