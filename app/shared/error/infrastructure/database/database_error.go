package dberr

import (
	"fmt"
	"net/http"

	infraerr "clinic-vet-api/app/shared/error/infrastructure"
	"clinic-vet-api/app/shared/log"

	"go.uber.org/zap"
)

func DatabaseOperationError(operation, tableName, dbType string, originalErr error) error {
	log.Error(
		"Database operation failed",
		originalErr,
		[]zap.Field{
			zap.String("operation", operation),
			zap.String("table", tableName),
			zap.String("database_type", dbType),
			zap.String("error_code", "DATABASE_ERROR"),
			zap.String("error_type", "infrastructure"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "DATABASE_ERROR",
		Type:       "infrastructure",
		Message:    originalErr.Error(),
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"operation":    operation,
			"table":        tableName,
			"databaseType": dbType,
		},
	}
}

func EntityNotFoundError(paramName, paramValue, operation, tableName, dbType string) error {
	errorMessage := fmt.Sprintf("%s with %s '%s' not found", tableName, paramName, paramValue)

	log.Warn(
		errorMessage,
		[]zap.Field{
			zap.String("param_name", paramName),
			zap.String("param_value", paramValue),
			zap.String("operation", operation),
			zap.String("table", tableName),
			zap.String("database_type", dbType),
			zap.String("error_code", "ENTITY_NOT_FOUND"),
			zap.String("error_type", "application"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:    "ENTITY_NOT_FOUND",
		Type:    "application",
		Message: errorMessage,
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

func InternalServerError(originalErr error) error {
	log.Error(
		"Internal server error occurred",
		originalErr,
		[]zap.Field{
			zap.String("error_code", "INTERNAL_SERVER_ERROR"),
			zap.String("error_type", "server"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "INTERNAL_SERVER_ERROR",
		Type:       "server",
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}

func ConstraintViolationError(operation, tableName, constraintName string, originalErr error) error {
	errorMessage := fmt.Sprintf("Constraint violation in %s operation", operation)

	log.Error(
		errorMessage,
		originalErr,
		[]zap.Field{
			zap.String("operation", operation),
			zap.String("table", tableName),
			zap.String("constraint", constraintName),
			zap.String("error_code", "CONSTRAINT_VIOLATION"),
			zap.String("error_type", "database"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "CONSTRAINT_VIOLATION",
		Type:       "database",
		Message:    errorMessage,
		StatusCode: http.StatusConflict,
		Details: map[string]string{
			"operation":  operation,
			"table":      tableName,
			"constraint": constraintName,
		},
	}
}

func DatabaseTimeoutError(operation, tableName string, originalErr error) error {
	errorMessage := fmt.Sprintf("Database timeout during %s operation", operation)

	log.Warn(
		errorMessage,
		zap.String("operation", operation),
		zap.String("table", tableName),
		zap.String("error_code", "DATABASE_TIMEOUT"),
		zap.String("error_type", "infrastructure"),
		zap.Error(originalErr),
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "DATABASE_TIMEOUT",
		Type:       "infrastructure",
		Message:    errorMessage,
		StatusCode: http.StatusGatewayTimeout, // 504 Gateway Timeout
		Details: map[string]string{
			"operation": operation,
			"table":     tableName,
		},
	}
}
