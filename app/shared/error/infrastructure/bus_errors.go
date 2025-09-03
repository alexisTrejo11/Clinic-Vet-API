package infraerr

import (
	"fmt"
	"net/http"
)

func NotRegistredCommandErr(commandName, entityBus string) *BaseInfrastructureError {
	return &BaseInfrastructureError{
		Message:    fmt.Sprintf("%v not registres on bus", commandName),
		Code:       "NOT_REGISTRED_COMMAND",
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"entityBus":     entityBus,
			"comandName":    commandName,
			"cqrsOperation": "command",
		},
	}
}

func InvalidCommandErr(message, commandName, entityBus string) *BaseInfrastructureError {
	return &BaseInfrastructureError{
		Message: message,
		Code:    "INVALID_COMMAND", StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"entityBus":     entityBus,
			"comandName":    commandName,
			"cqrsOperation": "command",
		},
	}
}

func NotRegistredQueryErr(queryName, entityBus string) *BaseInfrastructureError {
	return &BaseInfrastructureError{
		Message:    fmt.Sprintf("query with name %v not registres on bus", queryName),
		Code:       "NOT_REGISTRED_COMMAND",
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"cqrsOperation": "query",
			"entityBus":     entityBus,
			"queryName":     queryName,
		},
	}
}

func InvalidQuerryErr(message, commandName, entityBus string) *BaseInfrastructureError {
	return &BaseInfrastructureError{
		Message: message,
		Code:    "INVALID_COMMAND", StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"cqrsOperation": "query",
			"entityBus":     entityBus,
			"comandName":    commandName,
		},
	}
}
