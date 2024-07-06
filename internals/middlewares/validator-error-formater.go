package middlewares

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/masterghost2002/videotube/internals/validations"
	"github.com/masterghost2002/videotube/types"
)

func ValidatorErrorFormator(c *fiber.Ctx) error {
	err := c.Next()
	if err == nil {
		return nil
	}
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var v_errors []validations.ErrorFields
		for _, curr_err := range validationErrors {
			v_errors = append(v_errors, validations.ErrorFields{
				Field:   curr_err.Field(),
				Message: curr_err.Tag(),
			})
		}
		return c.Status(403).JSON(&types.Response{
			Error:       true,
			Message:     "Invalid Fields",
			ErrorFields: v_errors,
			Data:        nil,
		})

	}
	return err
}
