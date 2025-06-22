package authError

import (
	"net/http"

	appError "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors/application"
)

type AuthenticationError struct {
	applicationErr appError.BaseApplicationError
}

func NewAuthenticationError(message string) *AuthenticationError {

	return &AuthenticationError{
		applicationErr: appError.BaseApplicationError{
			Code:       "AUTHENTICATION_FAILED",
			Type:       "application",
			Message:    message,
			StatusCode: http.StatusUnauthorized,
		},
	}
}

type AuthorizationError struct {
	appError.BaseApplicationError
	RequiredRole string `json:"required_role"`
	UserRole     string `json:"user_role"`
}

func NewAuthorizationError(requiredRole, userRole string) *AuthorizationError {
	return &AuthorizationError{
		BaseApplicationError: appError.BaseApplicationError{
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
