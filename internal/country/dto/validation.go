package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/samuellfa/copa-do-mundo-golang/internal/shared"
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
		return &shared.RequestError{
			StatusCode: 400,
			Err:        fmt.Errorf(string(json)),
		}
	}

	return nil
}
