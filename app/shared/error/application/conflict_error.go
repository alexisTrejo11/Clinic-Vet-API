package apperror

import "net/http"

func ConflictError(resource, message string) error {
	return &BaseApplicationError{
		Code:       "RESOURCE_CONFLICT",
		Type:       "application",
		Message:    message,
		StatusCode: http.StatusConflict,
		Details: map[string]string{
			"resource": resource,
		},
	}
}
