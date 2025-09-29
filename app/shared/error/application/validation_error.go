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

func CommandDataValidationError(field, issue, commandName string) error {
	message := fmt.Sprintf("Invalid command data: field '%s' %s.", field, issue)
	log.Error(message, nil,
		append(
			[]zap.Field{log.WithOperation("command_validation")},
			zap.String("command", commandName),
			zap.String("result", "invalid_command_data"),
			zap.String("error_type", "validation"),
			zap.String("error_code", "INVALID_COMMAND_DATA"),
		)...,
	)

	return BaseApplicationError{
		Code:       "INVALID_COMMAND_DATA",
		Type:       "application",
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
		Details: map[string]string{
			"field":     field,
			"issue":     issue,
			"command":   commandName,
			"result":    "invalid command data",
			"errorType": "validation",
		},
	}
}

func QueryDataValidationError(field, issue, queryName string) error {
	message := fmt.Sprintf("Invalid query data: field '%s' %s.", field, issue)
	log.Error(message, nil,
		append(
			[]zap.Field{log.WithOperation("query_validation")},
			zap.String("query", queryName),
			zap.String("result", "invalid_query_data"),
			zap.String("error_type", "validation"),
			zap.String("error_code", "INVALID_QUERY_DATA"),
		)...,
	)

	return BaseApplicationError{
		Code:       "INVALID_QUERY_DATA",
		Type:       "application",
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
		Details: map[string]string{
			"field":     field,
			"issue":     issue,
			"query":     queryName,
			"result":    "invalid query data",
			"errorType": "validation",
		},
	}
}
