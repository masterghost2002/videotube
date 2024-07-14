package validations

import (
	"github.com/go-playground/validator/v10"
)

type ErrorFields struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func init() {
	_validate := validator.New()
	_validate.RegisterValidation("password", PasswordValidation)
	_validate.RegisterValidation("fullName", FullNameValidation)
	_validate.RegisterValidation("email", EmailValidation)
	_validate.RegisterValidation("username", UsernameValidation)
	_validate.RegisterValidation("channelName", ValidateChannelName)
	Validate = _validate
}
