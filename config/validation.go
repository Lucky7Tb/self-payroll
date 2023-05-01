package config

import (
	"fmt"
	"net/http"
	"self-payroll/common/structs"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	Validator *validator.Validate
}

func messageMap(validatorTag string, fieldName string, validatorParam string) string {
	switch validatorTag {
	case "required":
		return fmt.Sprintf("%s field is required", fieldName)
	case "numeric":
		return fmt.Sprintf("%s field must be numeric only", fieldName)
	case "min":
		return fmt.Sprintf("%s field minimum is %s", fieldName, validatorParam)
	default:
		return ""
	}
}

func (v *Validator) Validate(data interface{}) error {
	response := &structs.Response{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}
	if fieldErrors := v.Validator.Struct(data); fieldErrors != nil {
		errors := make([]string, len(fieldErrors.(validator.ValidationErrors)))
		for index, err := range fieldErrors.(validator.ValidationErrors) {
			errors[index] = messageMap(err.Tag(), err.Field(), err.Param())
		}
		response.Errors = &errors
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	return nil
}
