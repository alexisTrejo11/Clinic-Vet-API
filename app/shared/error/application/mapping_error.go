package apperror

import (
	"fmt"
	"net/http"
	"strings"
)

type FieldDataError struct {
	BaseApplicationError
}

func FieldValidationError(field, value, message string) *FieldDataError {
	return &FieldDataError{
		BaseApplicationError: BaseApplicationError{
			Code:    "FIELD_DATA_ERROR",
			Type:    "mapping",
			Message: message,
			Details: map[string]string{
				"field": field,
				"value": value,
			},
			StatusCode: 422,
		},
	}
}

func MappingError(errorMessage []string, from, to, entityName string) *FieldDataError {
	return &FieldDataError{
		BaseApplicationError: BaseApplicationError{
			Code:    "FIELD_DATA_ERROR",
			Type:    "mapping",
			Message: "error courred while mapping " + from,
			Details: map[string]string{
				"errorsMessages": strings.Join(errorMessage, ","),
				"from":           from,
				"to":             to,
				"entity":         entityName,
			},
			StatusCode: http.StatusUnprocessableEntity,
		},
	}
}

func InvalidFieldFormatError(field, format string) error {
	return BaseApplicationError{
		Code:    "INVALID_DATE_FORMAT",
		Type:    "DATE FORMAT",
		Message: fmt.Sprintf("Invalid date format, expected format: %s", format),
		Details: map[string]string{
			"field":  field,
			"format": format,
		},
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func InvalidParseFieldError(field string, value, meesage string) error {
	return BaseApplicationError{
		Code:    "INVALID_PARSE_DATA",
		Type:    "DATA PARSING",
		Message: fmt.Sprintf("Invalid data for field '%s': %s", field, value),
		Details: map[string]string{
			"field":   field,
			"value":   value,
			"message": meesage,
		},

		StatusCode: http.StatusUnprocessableEntity,
	}
}

func ValidationError(message string) error {
	return BaseApplicationError{
		Code:       "VALIDATION_ERROR",
		Type:       "validation",
		Message:    message,
		Details:    map[string]string{},
		StatusCode: http.StatusUnprocessableEntity,
	}
}
