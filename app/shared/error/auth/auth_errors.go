// Package autherror for application security errors
package autherror

import (
	"net/http"

	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
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
