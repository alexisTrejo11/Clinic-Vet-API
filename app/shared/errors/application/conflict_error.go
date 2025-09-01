package apperror

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
			StatusCode: 409,
			Data: map[string]any{
				"resource": resource,
			},
		},
		Resource: resource,
	}
}
