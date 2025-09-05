package dberr

import (
	"fmt"
	"net/http"

	infraerr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure"
)

func DatabaseOperationError(operation, tableName, dbType, errorMessage string) error {
	return &infraerr.BaseInfrastructureError{
		Code:       "DATABASE_ERROR",
		Type:       "infrastructure",
		Message:    errorMessage,
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"operation":    operation,
			"table":        tableName,
			"databaseType": dbType,
		},
	}
}

func EntityNotFoundError(paramName, paramValue, operation, tableName, dbType string) error {
	return &infraerr.BaseInfrastructureError{
		Code:    "ENTITY_NOT_FOUND",
		Type:    "application",
		Message: fmt.Sprintf("%s with %s not found", tableName, paramName),
		Details: map[string]string{
			"paramName":    paramName,
			"paramValue":   paramValue,
			"operation":    operation,
			"table":        tableName,
			"databaseType": dbType,
		},
		StatusCode: http.StatusNotFound,
	}
}

func InternalServerError() error {
	return &infraerr.BaseInfrastructureError{
		Code:       "INTERNAL_SERVER_ERROR",
		Type:       "server",
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}
