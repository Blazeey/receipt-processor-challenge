package utils

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateRegex(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	pattern := fl.Param()

	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return false
	}
	return matched
}

func ValidateTimeOnly(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	pattern := fl.Param()
	_, err := time.Parse(pattern, value)
	return err == nil
}
