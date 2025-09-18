// Package autherror provides authentication and authorization error handling with comprehensive logging
package autherror

import (
	"fmt"
	"net/http"

	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/log"

	"go.uber.org/zap"
)

func UnauthorizedError(message string) error {
	log.Warn(
		"Authentication failed",
		[]zap.Field{
			zap.String("error_code", "AUTHENTICATION_FAILED"),
			zap.String("error_type", "authentication"),
			zap.String("component", "auth"),
			zap.String("severity", "medium"),
			zap.String("message", message),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "AUTHENTICATION_FAILED",
		Type:       "authentication",
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func ForbiddenError(requiredRole, userRole string) error {
	log.Warn(
		"Authorization failed: insufficient permissions",
		[]zap.Field{
			zap.String("error_code", "AUTHORIZATION_FAILED"),
			zap.String("error_type", "authorization"),
			zap.String("component", "auth"),
			zap.String("severity", "medium"),
			zap.String("required_role", requiredRole),
			zap.String("user_role", userRole),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "AUTHORIZATION_FAILED",
		Type:       "authorization",
		Message:    "Insufficient permissions to perform this action",
		StatusCode: http.StatusForbidden,
		Details: map[string]string{
			"required_role": requiredRole,
			"user_role":     userRole,
		},
	}
}

func UnauthorizedCTXError() error {
	log.Warn(
		"User context missing: authentication required",
		[]zap.Field{
			zap.String("error_code", "AUTHENTICATION_FAILED"),
			zap.String("error_type", "context"),
			zap.String("component", "auth"),
			zap.String("severity", "medium"),
		}...,
	)

	return UnauthorizedError("user not present in context")
}

func InvalidCredentialsError(err error, identifier string) error {
	log.Warn(
		"Invalid credentials provided during authentication",
		[]zap.Field{
			zap.String("operation", "authentication"),
			zap.String("auth_method", "phone"),
			zap.String("identifier", identifier),
			zap.String("error_code", "INVALID_CREDENTIALS"),
			zap.String("error_type", "authentication"),
			zap.String("component", "auth"),
			zap.String("severity", "low"),
			zap.Error(err),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "INVALID_CREDENTIALS",
		Type:       "authentication",
		Message:    "invalid email/phone or password",
		StatusCode: http.StatusUnauthorized,
		Details: map[string]string{
			"identifier": identifier,
		},
	}
}

// TokenExpiredError handles expired authentication tokens
func TokenExpiredError(tokenType string) error {
	log.Warn(
		"Authentication token expired",
		[]zap.Field{
			zap.String("error_code", "TOKEN_EXPIRED"),
			zap.String("error_type", "authentication"),
			zap.String("component", "auth"),
			zap.String("severity", "low"),
			zap.String("token_type", tokenType),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "TOKEN_EXPIRED",
		Type:       "authentication",
		Message:    fmt.Sprintf("%s token has expired", tokenType),
		StatusCode: http.StatusUnauthorized,
		Details: map[string]string{
			"token_type": tokenType,
		},
	}
}

// InvalidTokenError handles malformed or invalid tokens
func InvalidTokenError(tokenType string, reason string) error {
	log.Warn(
		"Invalid authentication token",
		[]zap.Field{
			zap.String("error_code", "INVALID_TOKEN"),
			zap.String("error_type", "authentication"),
			zap.String("component", "auth"),
			zap.String("severity", "medium"),
			zap.String("token_type", tokenType),
			zap.String("reason", reason),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "INVALID_TOKEN",
		Type:       "authentication",
		Message:    "invalid authentication token",
		StatusCode: http.StatusUnauthorized,
		Details: map[string]string{
			"token_type": tokenType,
			"reason":     reason,
		},
	}
}

// AccountLockedError handles locked user accounts
func AccountLockedError(userID, reason string) error {
	log.Warn(
		"Account locked authentication attempt",
		[]zap.Field{
			zap.String("error_code", "ACCOUNT_LOCKED"),
			zap.String("error_type", "authentication"),
			zap.String("component", "auth"),
			zap.String("severity", "medium"),
			zap.String("user_id", userID),
			zap.String("lock_reason", reason),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "ACCOUNT_LOCKED",
		Type:       "authentication",
		Message:    "account is temporarily locked",
		StatusCode: http.StatusForbidden,
		Details: map[string]string{
			"user_id":     userID,
			"lock_reason": reason,
		},
	}
}

// TooManyAttemptsError handles rate limiting for authentication attempts
func TooManyAttemptsError(identifier string, remainingTime string) error {
	log.Warn(
		"Too many authentication attempts",
		[]zap.Field{
			zap.String("error_code", "TOO_MANY_ATTEMPTS"),
			zap.String("error_type", "rate_limiting"),
			zap.String("component", "auth"),
			zap.String("severity", "medium"),
			zap.String("identifier", identifier),
			zap.String("remaining_time", remainingTime),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "TOO_MANY_ATTEMPTS",
		Type:       "rate_limiting",
		Message:    fmt.Sprintf("too many authentication attempts, try again in %s", remainingTime),
		StatusCode: http.StatusTooManyRequests,
		Details: map[string]string{
			"identifier":     identifier,
			"remaining_time": remainingTime,
		},
	}
}

// SessionExpiredError handles expired user sessions
func SessionExpiredError(sessionID string) error {
	log.Info(
		"User session expired",
		[]zap.Field{
			zap.String("error_code", "SESSION_EXPIRED"),
			zap.String("error_type", "session"),
			zap.String("component", "auth"),
			zap.String("severity", "low"),
			zap.String("session_id", sessionID),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "SESSION_EXPIRED",
		Type:       "session",
		Message:    "user session has expired",
		StatusCode: http.StatusUnauthorized,
		Details: map[string]string{
			"session_id": sessionID,
		},
	}
}

// PermissionDeniedError handles specific permission violations
func PermissionDeniedError(userID, resource, action string) error {
	log.Warn(
		"Permission denied for resource action",
		[]zap.Field{
			zap.String("error_code", "PERMISSION_DENIED"),
			zap.String("error_type", "authorization"),
			zap.String("component", "auth"),
			zap.String("severity", "medium"),
			zap.String("user_id", userID),
			zap.String("resource", resource),
			zap.String("action", action),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:       "PERMISSION_DENIED",
		Type:       "authorization",
		Message:    fmt.Sprintf("access denied to %s resource for %s action", resource, action),
		StatusCode: http.StatusForbidden,
		Details: map[string]string{
			"user_id":  userID,
			"resource": resource,
			"action":   action,
		},
	}
}
