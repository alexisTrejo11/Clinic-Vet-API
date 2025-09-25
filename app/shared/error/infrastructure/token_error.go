package infraerr

import "net/http"

func TokenExpiredError(tokenType string) error {
	return &BaseInfrastructureError{
		Message:    "token expired",
		Code:       "TOKEN_EXPIRED",
		StatusCode: http.StatusUnauthorized,
		Details: map[string]string{
			"token_type": tokenType,
			"errorType":  "expired",
			"module":     "token_factory",
		},
	}
}
func InvalidTokenError(tokenType string) error {
	return &BaseInfrastructureError{
		Message:    "invalid token",
		Code:       "INVALID_TOKEN",
		StatusCode: http.StatusUnauthorized,
		Details: map[string]string{
			"token_type": tokenType,
			"errorType":  "invalid",
			"module":     "token_factory",
		},
	}
}

func TokenGenerationError(tokenType string, err error) error {
	return &BaseInfrastructureError{
		Message:    "error generating token: " + err.Error(),
		Code:       "TOKEN_GENERATION_ERROR",
		StatusCode: http.StatusInternalServerError,
		Details: map[string]string{
			"token_type": tokenType,
			"errorType":  "generation_error",
			"module":     "token_factory",
		},
	}
}

func TokeNotFoundError(tokenType string) error {
	return &BaseInfrastructureError{
		Message:    "token not found",
		Code:       "TOKEN_NOT_FOUND",
		StatusCode: http.StatusNotFound,
		Details: map[string]string{
			"token_type": tokenType,
			"errorType":  "not_found",
			"module":     "token_manager",
		},
	}
}
