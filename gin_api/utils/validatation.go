package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatationRequired(fieldName string, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", fieldName)
	}
	return nil
}

func ValidatationStringLenght(fieldName, value string, min, max int) error {
	l := len(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be between %d and %d characters", fieldName, min, max)
	}
	return nil
}
func ValidatationIntRange(fieldName string, value int, min, max int) error {
	if value < min || value > max {
		return fmt.Errorf("%s must be between %d and %d", fieldName, min, max)
	}
	return nil
}
func HandleValidationError(err error) gin.H {
	fmt.Println("Check error:", err)
	errors := make(map[string]string)
	if error, ok := err.(validator.ValidationErrors); ok {
		for _, e := range error {
			fmt.Println(e)
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " must be greater than " + e.Param()
			case "required":
				errors[e.Field()] = "Trường này là bắt buộc"
			case "min":
				errors[e.Field()] = e.Field() + " must be at least " + e.Param() + " characters"
			case "max":
				errors[e.Field()] = e.Field() + " must be at most " + e.Param() + " characters"
			case "email":
				errors[e.Field()] = e.Field() + " must be a valid email address"
			default:
				errors[e.Field()] = e.Field() + " is invalid"
			}
		}
	}
	return gin.H{
		"error": "Yêu cầu không hợp lệ", "errors": errors,
	}
}
