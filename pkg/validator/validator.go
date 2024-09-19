package validator

import (
	"log"

	lib "github.com/go-playground/validator/v10"
)

var (
	v = lib.New()
)

// Validate function for any type that can be validated
func Validate[T any](request T) error {
	err := v.Struct(request)

	if err != nil {
		for _, err := range err.(lib.ValidationErrors) {
			log.Printf("Validation failed for field '%s', condition: %s", err.Field(), err.ActualTag())
		}

		return err
	}

	return nil
}

// Map function for convert validation errors in map[string]string (JSON-parseable)
func Map(err error) map[string]string {
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(lib.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}
