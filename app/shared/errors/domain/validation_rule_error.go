package domain_error

type ValidationError struct {
	BaseDomainError
	Field string `json:"field"`
	Value string `json:"value"`
}

func NewValidationError(field, value, message string) *ValidationError {
	return &ValidationError{
		BaseDomainError: BaseDomainError{
			Code:    "VALIDATION_ERROR",
			Type:    "domain",
			Message: message,
			Data: map[string]interface{}{
				"field": field,
				"value": value,
			},
		},
		Field: field,
		Value: value,
	}
}
