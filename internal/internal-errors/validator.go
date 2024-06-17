package internalerrors

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func ValidatorStruct(obj interface{}) error {
	validate := validator.New()
	err := validate.Struct(obj)
	if err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		for _, validationError := range validationErrors {
			switch validationError.Tag() {
			case "required":
				return errors.New(validationError.StructField() + " is required")
			case "min":
				return errors.New(validationError.StructField() + " is too short")
			case "max":
				return errors.New(validationError.StructField() + " is too long")
			case "email":
				return errors.New(validationError.StructField() + " is invalid")
			case "dive":
				return errors.New(validationError.StructField() + " is invalid")
			default:
				break
			}
		}
	}

	return nil
}
