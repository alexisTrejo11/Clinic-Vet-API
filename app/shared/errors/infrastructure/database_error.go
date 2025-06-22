package infra_error

import "net/http"

type DatabaseError struct {
	BaseInfrastructureError
	Operation string `json:"operation"`
}

func NewDatabaseError(operation, message string) *DatabaseError {
	return &DatabaseError{
		BaseInfrastructureError: BaseInfrastructureError{
			Code:       "DATABASE_ERROR",
			Type:       "infrastructure",
			Message:    message,
			StatusCode: http.StatusInternalServerError,
			Data: map[string]interface{}{
				"operation": operation,
			},
		},
		Operation: operation,
	}
}
