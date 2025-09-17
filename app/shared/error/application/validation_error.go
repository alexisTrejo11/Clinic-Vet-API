package apperror

import (
	"clinic-vet-api/app/shared/log"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

func EntityNotFoundValidationError(entity, identifier, value string) error {
	message := fmt.Sprintf("The %s with %s '%s' was not found.", entity, identifier, value)
	log.Error(message, nil,
		append(
			log.WithEntity(entity, value),
			zap.String("identifier", identifier),
			zap.String("operation", "retrieving"),
			zap.String("result", "entity_not_found"),
			zap.String("error_type", "validation"),
			zap.String("error_code", "INVALID_ENTITY"),
		)...,
	)

	return BaseApplicationError{
		Code:       "INVALID_ENTITY",
		Type:       "application",
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
		Details: map[string]string{
			"entity":     entity,
			"identifier": identifier,
			"value":      value,
			"operation":  "retrieving",
			"result":     "entity not found",
		},
	}
}
