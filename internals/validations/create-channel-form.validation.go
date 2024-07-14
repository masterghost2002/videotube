package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type CreateChannelFormData struct {
	Name string  `json:"name" validate:"required,channelName"`
	Logo *string `json:"logo"`
}

func ValidateChannelName(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	if len(str) == 0 {
		return false
	}

	pattern := "^[A-Za-z]+$"
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
