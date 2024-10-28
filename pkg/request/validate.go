package request

import (
	"github.com/go-playground/validator"
)

func IsValid[T any](payload T) error {
	validator := validator.New()
	err := validator.Struct(payload)
	return err
}
