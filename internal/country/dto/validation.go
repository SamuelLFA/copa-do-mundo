package dto

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/samuellfa/copa-do-mundo-golang/internal/shared"
)

var validate = validator.New()

// ValidationError is the response sent if validation find any error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"error"`
}

// ValidateInput validate the input of the request
func ValidateInput(body io.Reader, request interface{}) error {
	if err := json.NewDecoder(body).Decode(request); err != nil {
		return err
	}

	if err := validate.Struct(request); err != nil {
		var errors []ValidationError
		validationErrors := err.(validator.ValidationErrors)
		for _, validationError := range validationErrors {
			errors = append(errors, ValidationError{
				Field:   validationError.Field(),
				Message: validationError.Tag(),
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
