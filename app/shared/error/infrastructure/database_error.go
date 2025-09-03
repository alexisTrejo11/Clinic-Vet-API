package infraerr

import (
	"fmt"
	"net/http"
)

type DatabaseError struct {
	BaseInfrastructureError
	Operation string `json:"operation"`
}

func NewDatabaseError(operation, message string) error {
	return &DatabaseError{
		BaseInfrastructureError: BaseInfrastructureError{
			Code:       "DATABASE_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Details: map[string]string{
				"operation": operation,
			},
		},
		Operation: operation,
	}
}

func EntityNotFoundError(entityType, entityID string) error {
	return &BaseInfrastructureError{
		Code:    "ENTITY_NOT_FOUND",
		Type:    "application",
		Message: fmt.Sprintf("%s with ID %s not found", entityType, entityID),
		Details: map[string]string{
			"entity_type": entityType,
			"entity_id":   entityID,
		},
		StatusCode: http.StatusNotFound,
	}
}

func InternalServerError() error {
	return &BaseInfrastructureError{
		Code:       "INTERNAL_SERVER_ERROR",
		Type:       "server",
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}
