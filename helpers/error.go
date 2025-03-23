package helpers

import (
	"github.com/go-playground/validator/v10"
)

func FormatValidatorError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func Errorlogin(msg string) string {
	var errors []string
	errors = append(errors, msg)

	return errors[0]
}
