// Package domainerr contains all the custom buissness logic error applied to entiy package
package domainerr

import (
	"fmt"
	"net/http"
)

type BaseDomainError struct {
	Code       string            `json:"code"`
	Type       string            `json:"type"`
	Message    string            `json:"message"`
	StatusCode int               `json:"-"`
	Details    map[string]string `json:"details,omitempty"`
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

func (e BaseDomainError) HTTPStatus() int {
	if e.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return e.StatusCode
}

func (e BaseDomainError) DetailMap() map[string]string {
	return e.Details
}

func NewBusinessRuleError(rule, entity, field string) error {
	return &BaseDomainError{
		Code:    "BUSINESS_RULE_VIOLATION",
		Type:    "domain",
		Message: rule,
		Details: map[string]string{
			"rule":   rule,
			"entity": entity,
			"field":  field,
		},
	}
}

type ValidationError struct {
	BaseDomainError
}

func NewValidationError(field, value, message string) *ValidationError {
	return &ValidationError{
		BaseDomainError: BaseDomainError{
			Code:    "VALIDATION_ERROR",
			Type:    "domain",
			Message: message,
			Details: map[string]string{
				"field": field,
				"value": value,
			},
		},
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

func EntityNotFoundError(entityType, entityID string) error {
	return &BaseDomainError{
		Code:    "ENTITY_NOT_FOUND",
		Type:    "application",
		Message: fmt.Sprintf("%s with ID %s not found", entityType, entityID),
		Details: map[string]string{
			"entity_type": entityType,
			"entity_id":   entityID,
		},
		StatusCode: 404,
	}
}

type ConflictError struct {
	BaseDomainError
}

func NewConflictError(resource, message string) *ConflictError {
	return &ConflictError{
		BaseDomainError: BaseDomainError{
			Code:       "RESOURCE_CONFLICT",
			Type:       "application",
			Message:    message,
			StatusCode: http.StatusConflict,
			Details: map[string]string{
				"resource": resource,
			},
		},
	}
}
