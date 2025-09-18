// Package dberr provides database error handling with comprehensive logging
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
			zap.String("component", "database"),
			zap.String("severity", "high"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "DATABASE_ERROR",
		Type:       "infrastructure",
		Message:    fmt.Sprintf("Database operation '%s' failed on table '%s': %s", operation, tableName, originalErr.Error()),
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"operation":     operation,
			"table":         tableName,
			"databaseType":  dbType,
			"originalError": originalErr.Error(),
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
			zap.String("component", "database"),
			zap.String("severity", "medium"),
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
			zap.String("component", "server"),
			zap.String("severity", "critical"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "INTERNAL_SERVER_ERROR",
		Type:       "server",
		Message:    "Internal server error occurred",
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"originalError": originalErr.Error(),
		},
	}
}

func ConstraintViolationError(operation, tableName, constraintName string, originalErr error) error {
	errorMessage := fmt.Sprintf("Constraint violation in %s operation on table %s", operation, tableName)

	log.Error(
		errorMessage,
		originalErr,
		[]zap.Field{
			zap.String("operation", operation),
			zap.String("table", tableName),
			zap.String("constraint", constraintName),
			zap.String("error_code", "CONSTRAINT_VIOLATION"),
			zap.String("error_type", "database"),
			zap.String("component", "database"),
			zap.String("severity", "medium"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "CONSTRAINT_VIOLATION",
		Type:       "database",
		Message:    errorMessage,
		StatusCode: http.StatusConflict,
		Details: map[string]string{
			"operation":     operation,
			"table":         tableName,
			"constraint":    constraintName,
			"originalError": originalErr.Error(),
		},
	}
}

func DatabaseTimeoutError(operation, tableName string, originalErr error) error {
	errorMessage := fmt.Sprintf("Database timeout during %s operation on table %s", operation, tableName)

	log.Error(
		errorMessage,
		originalErr,
		[]zap.Field{
			zap.String("operation", operation),
			zap.String("table", tableName),
			zap.String("error_code", "DATABASE_TIMEOUT"),
			zap.String("error_type", "infrastructure"),
			zap.String("component", "database"),
			zap.String("severity", "high"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "DATABASE_TIMEOUT",
		Type:       "infrastructure",
		Message:    errorMessage,
		StatusCode: http.StatusGatewayTimeout,
		Details: map[string]string{
			"operation":     operation,
			"table":         tableName,
			"originalError": originalErr.Error(),
		},
	}
}

// DatabaseConnectionError handles database connection failures
func DatabaseConnectionError(dbType string, originalErr error) error {
	errorMessage := fmt.Sprintf("Failed to connect to %s database", dbType)

	log.Error(
		errorMessage,
		originalErr,
		[]zap.Field{
			zap.String("database_type", dbType),
			zap.String("error_code", "DATABASE_CONNECTION_ERROR"),
			zap.String("error_type", "infrastructure"),
			zap.String("component", "database"),
			zap.String("severity", "critical"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "DATABASE_CONNECTION_ERROR",
		Type:       "infrastructure",
		Message:    errorMessage,
		StatusCode: http.StatusServiceUnavailable,
		Details: map[string]string{
			"databaseType":  dbType,
			"originalError": originalErr.Error(),
		},
	}
}

// TransactionError handles database transaction failures
func TransactionError(operation string, originalErr error) error {
	errorMessage := fmt.Sprintf("Transaction failed during %s operation", operation)

	log.Error(
		errorMessage,
		originalErr,
		[]zap.Field{
			zap.String("operation", operation),
			zap.String("error_code", "TRANSACTION_ERROR"),
			zap.String("error_type", "database"),
			zap.String("component", "database"),
			zap.String("severity", "high"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "TRANSACTION_ERROR",
		Type:       "database",
		Message:    errorMessage,
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"operation":     operation,
			"originalError": originalErr.Error(),
		},
	}
}

// DuplicateKeyError handles duplicate key constraint violations
func DuplicateKeyError(tableName, keyField, keyValue string, originalErr error) error {
	errorMessage := fmt.Sprintf("Duplicate key violation: %s '%s' already exists in table %s", keyField, keyValue, tableName)

	log.Warn(
		errorMessage,
		[]zap.Field{
			zap.String("table", tableName),
			zap.String("key_field", keyField),
			zap.String("key_value", keyValue),
			zap.String("error_code", "DUPLICATE_KEY_ERROR"),
			zap.String("error_type", "database"),
			zap.String("component", "database"),
			zap.String("severity", "low"),
			zap.Error(originalErr),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "DUPLICATE_KEY_ERROR",
		Type:       "database",
		Message:    errorMessage,
		StatusCode: http.StatusConflict,
		Details: map[string]string{
			"table":         tableName,
			"keyField":      keyField,
			"keyValue":      keyValue,
			"originalError": originalErr.Error(),
		},
	}
}

// QueryExecutionError handles SQL query execution failures
func QueryExecutionError(query, operation string, originalErr error) error {
	errorMessage := fmt.Sprintf("Query execution failed during %s operation", operation)

	log.Error(
		errorMessage,
		originalErr,
		[]zap.Field{
			zap.String("operation", operation),
			zap.String("query_preview", truncateQuery(query, 100)), // Log only first 100 chars
			zap.String("error_code", "QUERY_EXECUTION_ERROR"),
			zap.String("error_type", "database"),
			zap.String("component", "database"),
			zap.String("severity", "high"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "QUERY_EXECUTION_ERROR",
		Type:       "database",
		Message:    errorMessage,
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"operation":     operation,
			"originalError": originalErr.Error(),
		},
	}
}

// DataIntegrityError handles data integrity violations
func DataIntegrityError(tableName, field, value string, originalErr error) error {
	errorMessage := fmt.Sprintf("Data integrity violation: invalid %s value '%s' for table %s", field, value, tableName)

	log.Error(
		errorMessage,
		originalErr,
		[]zap.Field{
			zap.String("table", tableName),
			zap.String("field", field),
			zap.String("value", value),
			zap.String("error_code", "DATA_INTEGRITY_ERROR"),
			zap.String("error_type", "database"),
			zap.String("component", "database"),
			zap.String("severity", "medium"),
		}...,
	)

	return &infraerr.BaseInfrastructureError{
		Code:       "DATA_INTEGRITY_ERROR",
		Type:       "database",
		Message:    errorMessage,
		StatusCode: http.StatusUnprocessableEntity,
		Details: map[string]string{
			"table":         tableName,
			"field":         field,
			"value":         value,
			"originalError": originalErr.Error(),
		},
	}
}

// Helper function to truncate query for logging
func truncateQuery(query string, maxLength int) string {
	if len(query) <= maxLength {
		return query
	}
	return query[:maxLength] + "..."
}
