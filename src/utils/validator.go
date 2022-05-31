package utils

import (
	"github.com/go-playground/validator/v10"
)

type customValidator struct {
	validator *validator.Validate
}

func SetValidator() *customValidator {
	validator := &customValidator{validator: validator.New()}
	return validator
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
