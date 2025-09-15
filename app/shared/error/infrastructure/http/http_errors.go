package http

import (
	"net/http"

	apperror "clinic-vet-api/app/shared/error/application"
)

func RequestURLParamError(err error, field string, value string) error {
	return apperror.BaseApplicationError{
		Code:       "INVALID_URL_PARAM",
		Type:       "ROUTING",
		Message:    err.Error(),
		Details:    map[string]string{},
		StatusCode: http.StatusBadRequest,
	}
}

func InternalServerError(err error) error {
	return apperror.BaseApplicationError{
		Code:       "INTERNAL_SERVER_ERROR",
		Type:       "SERVER",
		Message:    "an internal server error occurred",
		Details:    map[string]string{"error": err.Error()},
		StatusCode: http.StatusInternalServerError,
	}
}

func RequestURLQueryError(err error, queryURL string) error {
	return apperror.BaseApplicationError{
		Code:    "INVALID_URL_PARAMS",
		Type:    "ROUTING",
		Message: "invalid query parameters",
		Details: map[string]string{
			"query":   queryURL,
			"error":   err.Error(),
			"context": "parsing query parameters",
		},
		StatusCode: http.StatusBadRequest,
	}
}

func RequestBodyDataError(err error) error {
	return apperror.BaseApplicationError{
		Code:       "REQUEST_BODY_ERROR",
		Type:       "BODY_FORMAT",
		Message:    err.Error(),
		Details:    map[string]string{},
		StatusCode: http.StatusBadRequest,
	}
}

func InvalidDataError(err error) error {
	return apperror.BaseApplicationError{
		Code:       "INVALID_DATA_VALUES",
		Type:       "Data Validation",
		Message:    err.Error(),
		Details:    map[string]string{},
		StatusCode: http.StatusUnprocessableEntity,
	}
}
