package apperror

import (
	"fmt"
	"net/http"
)

type EntityNotFoundError struct {
	BaseApplicationError
	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`
}

func NewEntityNotFoundError(entityType, entityID string) *EntityNotFoundError {
	return &EntityNotFoundError{
		BaseApplicationError: BaseApplicationError{
			Code:    "ENTITY_NOT_FOUND",
			Type:    "application",
			Message: fmt.Sprintf("%s with ID %s not found", entityType, entityID),
			Data: map[string]string{
				"entity_type": entityType,
				"entity_id":   entityID,
			},
			StatusCode: 404,
		},
		EntityType: entityType,
		EntityID:   entityID,
	}
}

func NewInternalServerError() *BaseApplicationError {
	return &BaseApplicationError{
		Code:       "INTERNAL_SERVER_ERROR",
		Type:       "server",
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}
}
