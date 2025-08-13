package domainErr

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

func InvalidEnumValue(field, value, message string) *ValidationError {
	return NewValidationError(field, value, message)
}

func InvalidFormat(field, value, message string) *ValidationError {
	return NewValidationError(field, value, message)
}

func RequiredField(field, message string) *ValidationError {
	return NewValidationError(field, "", message)
}
