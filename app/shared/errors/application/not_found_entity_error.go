package appError

import "fmt"

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
			Data: map[string]interface{}{
				"entity_type": entityType,
				"entity_id":   entityID,
			},
		},
		EntityType: entityType,
		EntityID:   entityID,
	}
}
