package domainerr

import (
	"fmt"
	"net/http"
)

type BaseDomainError struct {
	Code       string `json:"code"`
	Type       string `json:"type"`
	Message    string `json:"message"`
	StatusCode int
	Data       map[string]interface{} `json:"details,omitempty"`
}

func (e BaseDomainError) Error() string {
	return e.Message
}

func (e BaseDomainError) ErrorCode() string {
	return e.Code
}

func (e BaseDomainError) ErrorType() string {
	return e.Type
}

func (e BaseDomainError) Details() map[string]interface{} {
	return e.Data
}

type BusinessRuleError struct {
	BaseDomainError
	Rule string `json:"rule"`
}

func NewBusinessRuleError(rule, message string) *BusinessRuleError {
	return &BusinessRuleError{
		BaseDomainError: BaseDomainError{
			Code:    "BUSINESS_RULE_VIOLATION",
			Type:    "domain",
			Message: message,
			Data: map[string]interface{}{
				"rule": rule,
			},
		},
		Rule: rule,
	}
}

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

type EntityNotFoundError struct {
	BaseDomainError
	EntityType string `json:"entity_type"`
	EntityID   string `json:"entity_id"`
}

func NewEntityNotFoundError(entityType, entityID string) *EntityNotFoundError {
	return &EntityNotFoundError{
		EntityType: entityType,
		EntityID:   entityID,
		BaseDomainError: BaseDomainError{
			Code:    "ENTITY_NOT_FOUND",
			Type:    "application",
			Message: fmt.Sprintf("%s with ID %s not found", entityType, entityID),
			Data: map[string]interface{}{
				"entity_type": entityType,
				"entity_id":   entityID,
			},
			StatusCode: 404,
		},
	}
}

func NewInternalServerError() *BaseDomainError {
	return &BaseDomainError{
		Code:       "INTERNAL_SERVER_ERROR",
		Type:       "server",
		Message:    "Internal server error",
		StatusCode: 500,
	}
}

type ConflictError struct {
	BaseDomainError
	Resource string `json:"resource"`
}

func NewConflictError(resource, message string) *ConflictError {
	return &ConflictError{
		BaseDomainError: BaseDomainError{
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
