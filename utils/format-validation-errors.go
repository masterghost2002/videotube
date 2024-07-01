package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/masterghost2002/videotube/internals/validations"
)

func FormatValidationErrors(err error) []validations.ErrorFields {
	var errors []validations.ErrorFields
	for _, validationError := range err.(validator.ValidationErrors) {
		errors = append(errors, validations.ErrorFields{
			Field:   validationError.Field(),
			Message: validationError.Tag(),
		})
	}
	return errors
}
