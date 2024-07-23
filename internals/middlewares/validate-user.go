package middlewares

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/database"
	jwt "github.com/masterghost2002/videotube/internals/token"
	"github.com/masterghost2002/videotube/types"
)

func ValidateUser(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(401).JSON(&types.Response{
			Error:       true,
			ErrorFields: []string{"token"},
			Message:     "unauthorized",
			Data:        nil,
		})
	}

	userData, err := jwt.ValidateJWT(token)
	if err != nil {
		return c.Status(401).JSON(&types.Response{
			Error:       true,
			ErrorFields: []string{"token"},
			Message:     "unauthorized",
			Data:        nil,
		})
	}
	userFromDB, err := database.Storage.GetUserByEmail(context.Background(), userData.Email)
	if err != nil {
		return err
	}
	c.Locals("logged_user", userFromDB)
	c.Next()
	return nil
}
