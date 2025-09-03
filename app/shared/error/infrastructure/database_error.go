package infraerr

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
			Data: map[string]any{
				"operation": operation,
			},
		},
		Operation: operation,
	}
}

func MapFieldErrorToDatabaseError(err error, operation string) *DatabaseError {
	return &DatabaseError{
		BaseInfrastructureError: BaseInfrastructureError{
			Code:       "DATABASE_ERROR",
			Type:       "infrastructure",
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
			Data: map[string]any{
				"operation": operation,
			},
		},
		Operation: operation,
	}
}
