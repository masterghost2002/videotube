package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/models"
	repository "github.com/masterghost2002/videotube/internals/repository/database"
	"github.com/masterghost2002/videotube/utils"
)

type UserData struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(c *fiber.Ctx) error {
	var userData UserData
	if err := c.BodyParser(&userData); err != nil {
		return err
	}
	hashPassword := utils.HashString(userData.Password)
	user := models.User{FullName: userData.FullName, Email: userData.Email, Username: userData.Username, Password: hashPassword}

	if err := repository.CreateUser(user); err != nil {
		return c.Status(424).JSON(fiber.Map{
			"error": err,
		})
	}
	return c.SendStatus(201)
}
