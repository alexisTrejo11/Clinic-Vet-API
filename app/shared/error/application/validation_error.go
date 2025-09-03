package apperror

import (
	"fmt"
	"net/http"
)

func EntityValidationError(entity, identifier, value string) error {
	return BaseApplicationError{
		Code:       "INVALID_ENTITY",
		Type:       "application",
		Message:    fmt.Sprintf("Invalid id provided for %s ", entity),
		StatusCode: http.StatusUnprocessableEntity,
		Details: map[string]string{
			"enitty":     entity,
			"identifier": identifier,
			"operation":  "retrieving",
			"result":     "entity not found",
		},
	}
}
