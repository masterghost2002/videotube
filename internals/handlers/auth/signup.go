package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type UserData struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
