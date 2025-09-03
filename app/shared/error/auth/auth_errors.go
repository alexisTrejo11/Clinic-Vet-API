// Package autherror for application security errors
package autherror

import (
	"net/http"

	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type AuthenticationError struct {
	applicationErr apperror.BaseApplicationError
}

func UnauthorizedError(message string) *AuthenticationError {
	return &AuthenticationError{
		applicationErr: apperror.BaseApplicationError{
			Code:       "AUTHENTICATION_FAILED",
			Type:       "application",
			Message:    message,
			StatusCode: http.StatusUnauthorized,
		},
	}
}

type AuthorizationError struct {
	apperror.BaseApplicationError
	RequiredRole string `json:"required_role"`
	UserRole     string `json:"user_role"`
}

func ForbbidenError(requiredRole, userRole string) *AuthorizationError {
	return &AuthorizationError{
		BaseApplicationError: apperror.BaseApplicationError{
			Code:       "AUTHORIZATION_FAILED",
			Type:       "application",
			Message:    "Insufficient permissions to perform this action",
			StatusCode: http.StatusForbidden,
			Details: map[string]string{
				"required_role": requiredRole,
				"user_role":     userRole,
			},
		},
		RequiredRole: requiredRole,
		UserRole:     userRole,
	}
}
