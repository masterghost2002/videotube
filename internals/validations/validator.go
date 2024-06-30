package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ErrorFields struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
type UserSignUpFormData struct {
	FullName string `json:"fullName" validate:"required,fullName"`
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username"`
	Password string `json:"password" validate:"required,password"`
}

var Validate *validator.Validate

func FullNameValidation(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	if len(str) == 0 {
		return false
	}
	pattern := "^[a-zA-Z]+( [a-zA-Z]+)*$"
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
func PasswordValidation(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	if len(password) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case 'A' <= char && char <= 'Z':
			hasUpper = true
		case 'a' <= char && char <= 'z':
			hasLower = true
		case '0' <= char && char <= '9':
			hasNumber = true
		default:
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}
func UsernameValidation(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	if len(username) < 3 || len(username) > 20 {
		return false
	}

	// Regex pattern to match alphanumeric characters, underscores, and periods
	pattern := `^[a-zA-Z0-9_.]+$`
	matched, _ := regexp.MatchString(pattern, username)
	return matched
}
func EmailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Regex pattern for basic email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(pattern, email)
	return matched
}
func init() {
	_validate := validator.New()
	_validate.RegisterValidation("password", PasswordValidation)
	_validate.RegisterValidation("fullName", FullNameValidation)
	_validate.RegisterValidation("email", EmailValidation)
	_validate.RegisterValidation("username", UsernameValidation)
	Validate = _validate
}
