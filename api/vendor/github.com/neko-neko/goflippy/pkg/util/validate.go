package util

import validator "gopkg.in/go-playground/validator.v9"

var validate = validator.New()

// ValidateStruct validate struct members
func ValidateStruct(data interface{}) validator.ValidationErrors {
	err := validate.Struct(data)
	if err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}
