// Package autherror for application security errors
package autherror

import (
	"net/http"

	ApplicationError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
)

type AuthenticationError struct {
	applicationErr ApplicationError.BaseApplicationError
}

func UnauthorizedError(message string) *AuthenticationError {
	return &AuthenticationError{
		applicationErr: ApplicationError.BaseApplicationError{
			Code:       "AUTHENTICATION_FAILED",
			Type:       "application",
			Message:    message,
			StatusCode: http.StatusUnauthorized,
		},
	}
}

type AuthorizationError struct {
	ApplicationError.BaseApplicationError
	RequiredRole string `json:"required_role"`
	UserRole     string `json:"user_role"`
}

func ForbbidenError(requiredRole, userRole string) *AuthorizationError {
	return &AuthorizationError{
		BaseApplicationError: ApplicationError.BaseApplicationError{
			Code:       "AUTHORIZATION_FAILED",
			Type:       "application",
			Message:    "Insufficient permissions to perform this action",
			StatusCode: http.StatusForbidden,
			Data: map[string]interface{}{
				"required_role": requiredRole,
				"user_role":     userRole,
			},
		},
		RequiredRole: requiredRole,
		UserRole:     userRole,
	}
}
