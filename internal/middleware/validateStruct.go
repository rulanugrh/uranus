package middleware

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/rulanugrh/uranus/internal/domain/web"
)

func ValidateStruct(validate *validator.Validate, data interface{}) error {
	err := validate.Struct(data)
	
	if err != nil {
		errors := []web.ValidationList{}
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(data).FieldByName(err.Field())
			errors = append(errors, web.ValidationList{
				Field: field.Name,
				Error: err.Field() + "|" + err.ActualTag(),
			})

		}

		return web.ValidationError{
			Message: "validation errors",
			Errors: errors,
		}
		
	}
	return nil
}