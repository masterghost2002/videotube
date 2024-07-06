package middlewares

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func SQLErrorHandler(c *fiber.Ctx) error {
	// first the handler functions will work
	err := c.Next() // handler functions return error

	// handling of error
	if err == nil {
		return nil
	}

	var pqError *pq.Error
	if errors.As(err, &pqError) {
		return c.Status(500).JSON(fiber.Map{
			"error":   true,
			"message": pqError.Message,
			"detail":  pqError.Detail,
			"code":    pqError.Code.Name(),
		})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"error":   true,
		"message": err.Error(),
	})
}
