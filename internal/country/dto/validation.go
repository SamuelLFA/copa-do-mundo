package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field   string
	Message string
}

func ValidateInput[T any](body io.Reader, request *T) error {
	if err := json.NewDecoder(body).Decode(request); err != nil {
		return err
	}

	if err := validate.Struct(request); err != nil {
		var errors []ValidationError
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   validationError.Field(),
				Message: validationError.ActualTag(),
			})
		}
		json, _ := json.Marshal(errors)
		return fmt.Errorf(string(json))
	}

	return nil
}
