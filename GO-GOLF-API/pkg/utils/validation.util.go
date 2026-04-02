package utils

import (
	"GO-GOLF-API/pkg/response"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// ValidateStruct validates a struct and returns a formatted APIError if validation fails
func ValidateStruct(data interface{}, validate *validator.Validate) *response.AppError {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return response.NewAppError(http.StatusBadRequest, response.ErrValidationFailed, err)
	}

	var errorCodes []string
	for _, fieldErr := range validationErrs {
		fieldName := fieldErr.Field()

		code := getFieldCode(data, fieldName)
		errorCodes = append(errorCodes, code)
	}

	return response.NewAppErrorWithMessage(http.StatusBadRequest, response.ErrValidationFailed, strings.Join(errorCodes, "-"), gin.Error{})
}

func getFieldCode(data interface{}, fieldName string) string {
	t := reflect.TypeOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	field, ok := t.FieldByName(fieldName)
	if !ok {
		return ""
	}

	return field.Tag.Get("code")
}
