package appError

type ValidationError struct {
	BaseApplicationError
	Field string `json:"entity_type"`
	Value string `json:"entity_id"`
}

func NewValidationError(field, value, message string) *ValidationError {
	return &ValidationError{
		BaseApplicationError: BaseApplicationError{
			Code:    "VALIDATION_ERROR",
			Type:    "application",
			Message: message,
			Data: map[string]interface{}{
				"field": field,
				"value": value,
			},
			StatusCode: 400,
		},
		Field: field,
		Value: value,
	}
}
