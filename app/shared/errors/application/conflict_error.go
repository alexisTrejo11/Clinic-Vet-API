package appError

import "net/http"

type ConflictError struct {
	BaseApplicationError
	Resource string `json:"resource"`
}

func NewConflictError(resource, message string) *ConflictError {
	return &ConflictError{
		BaseApplicationError: BaseApplicationError{
			Code:       "RESOURCE_CONFLICT",
			Type:       "application",
			Message:    message,
			StatusCode: http.StatusConflict,
			Data: map[string]interface{}{
				"resource": resource,
			},
		},
		Resource: resource,
	}
}
