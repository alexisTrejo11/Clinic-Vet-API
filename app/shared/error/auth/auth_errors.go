// Package autherror for application security errors
package autherror

import (
	"clinic-vet-api/app/shared/log"
	"net/http"

	apperror "clinic-vet-api/app/shared/error/application"

	"go.uber.org/zap"
)

func UnauthorizedError(message string) error {
	return &apperror.BaseApplicationError{
		Code:       "AUTHENTICATION_FAILED",
		Type:       "application",
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func ForbbidenError(requiredRole, userRole string) error {
	return apperror.BaseApplicationError{
		Code:       "AUTHORIZATION_FAILED",
		Type:       "application",
		Message:    "Insufficient permissions to perform this action",
		StatusCode: http.StatusForbidden,
		Details: map[string]string{
			"required_role": requiredRole,
			"user_role":     userRole,
		},
	}
}

func UnauthorizedCTXError() error {
	return UnauthorizedError("user not present in context")
}

func InvalidCredentialsError(err error, identifier string) error {
	log.Warn(
		"Error searching user by phone for authentication",
		[]zap.Field{
			log.WithOperation("authentication"),
			zap.String("auth_method", "phone"),
			zap.String("identifier", identifier),
			zap.String("error_type", "repository"),
			zap.Error(err),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "INVALID_CREDENTIALS",
		Type:       "application",
		Message:    "invalid email/phone or password",
		StatusCode: http.StatusUnauthorized,
	}
}
