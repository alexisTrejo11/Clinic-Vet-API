// Package http provides HTTP error handling with comprehensive logging
package http

import (
	"fmt"
	"net/http"

	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/log"

	"go.uber.org/zap"
)

func RequestURLParamError(err error, field string, value string) error {
	log.Warn(
		"Invalid URL parameter provided",
		[]zap.Field{
			zap.String("error_code", "INVALID_URL_PARAM"),
			zap.String("error_type", "routing"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.String("field", field),
			zap.String("value", value),
			zap.Error(err),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "INVALID_URL_PARAM",
		Type:    "routing",
		Message: fmt.Sprintf("invalid URL parameter '%s' with value '%s': %s", field, value, err.Error()),
		Details: map[string]string{
			"field":         field,
			"value":         value,
			"originalError": err.Error(),
		},
		StatusCode: http.StatusBadRequest,
	}
}

func InternalServerError(err error) error {
	log.Error(
		"Internal server error occurred",
		err,
		[]zap.Field{
			zap.String("error_code", "INTERNAL_SERVER_ERROR"),
			zap.String("error_type", "server"),
			zap.String("component", "http"),
			zap.String("severity", "critical"),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "INTERNAL_SERVER_ERROR",
		Type:    "server",
		Message: "an internal server error occurred",
		Details: map[string]string{
			"originalError": err.Error(),
		},
		StatusCode: http.StatusInternalServerError,
	}
}

func RequestURLQueryError(err error, queryURL string) error {
	log.Warn(
		"Invalid query parameters in request",
		[]zap.Field{
			zap.String("error_code", "INVALID_URL_PARAMS"),
			zap.String("error_type", "routing"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.String("query_url", queryURL),
			zap.Error(err),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "INVALID_URL_PARAMS",
		Type:    "routing",
		Message: "invalid query parameters",
		Details: map[string]string{
			"query":         queryURL,
			"originalError": err.Error(),
			"context":       "parsing query parameters",
		},
		StatusCode: http.StatusBadRequest,
	}
}

func RequestBodyDataError(err error) error {
	log.Warn(
		"Invalid request body format",
		[]zap.Field{
			zap.String("error_code", "REQUEST_BODY_ERROR"),
			zap.String("error_type", "body_format"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.Error(err),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "REQUEST_BODY_ERROR",
		Type:    "body_format",
		Message: fmt.Sprintf("invalid request body format: %s", err.Error()),
		Details: map[string]string{
			"originalError": err.Error(),
		},
		StatusCode: http.StatusBadRequest,
	}
}

func InvalidDataError(err error) error {
	log.Warn(
		"Data validation failed",
		[]zap.Field{
			zap.String("error_code", "INVALID_DATA_VALUES"),
			zap.String("error_type", "data_validation"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.Error(err),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "INVALID_DATA_VALUES",
		Type:    "data_validation",
		Message: err.Error(),
		Details: map[string]string{
			"originalError": err.Error(),
		},
		StatusCode: http.StatusUnprocessableEntity,
	}
}

// ValidationError handles field-specific validation errors
func ValidationError(field, value, reason string) error {
	log.Warn(
		"Field validation failed",
		[]zap.Field{
			zap.String("error_code", "FIELD_VALIDATION_ERROR"),
			zap.String("error_type", "validation"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.String("field", field),
			zap.String("value", value),
			zap.String("reason", reason),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "FIELD_VALIDATION_ERROR",
		Type:    "validation",
		Message: fmt.Sprintf("validation failed for field '%s': %s", field, reason),
		Details: map[string]string{
			"field":  field,
			"value":  value,
			"reason": reason,
		},
		StatusCode: http.StatusBadRequest,
	}
}

// MissingRequiredFieldError handles missing required fields
func MissingRequiredFieldError(field string) error {
	log.Warn(
		"Required field missing in request",
		[]zap.Field{
			zap.String("error_code", "MISSING_REQUIRED_FIELD"),
			zap.String("error_type", "validation"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.String("field", field),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "MISSING_REQUIRED_FIELD",
		Type:    "validation",
		Message: fmt.Sprintf("required field '%s' is missing", field),
		Details: map[string]string{
			"field": field,
		},
		StatusCode: http.StatusBadRequest,
	}
}

// RequestTooLargeError handles request size limit exceeded
func RequestTooLargeError(maxSize, actualSize int64) error {
	log.Warn(
		"Request size exceeds limit",
		[]zap.Field{
			zap.String("error_code", "REQUEST_TOO_LARGE"),
			zap.String("error_type", "request_size"),
			zap.String("component", "http"),
			zap.String("severity", "medium"),
			zap.Int64("max_size", maxSize),
			zap.Int64("actual_size", actualSize),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "REQUEST_TOO_LARGE",
		Type:    "request_size",
		Message: fmt.Sprintf("request size %d bytes exceeds maximum allowed size %d bytes", actualSize, maxSize),
		Details: map[string]string{
			"max_size":    fmt.Sprintf("%d", maxSize),
			"actual_size": fmt.Sprintf("%d", actualSize),
		},
		StatusCode: http.StatusRequestEntityTooLarge,
	}
}

// UnsupportedMediaTypeError handles unsupported content types
func UnsupportedMediaTypeError(contentType string, supportedTypes []string) error {
	log.Warn(
		"Unsupported media type in request",
		[]zap.Field{
			zap.String("error_code", "UNSUPPORTED_MEDIA_TYPE"),
			zap.String("error_type", "media_type"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.String("content_type", contentType),
			zap.Strings("supported_types", supportedTypes),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "UNSUPPORTED_MEDIA_TYPE",
		Type:    "media_type",
		Message: fmt.Sprintf("unsupported media type '%s'", contentType),
		Details: map[string]string{
			"content_type":    contentType,
			"supported_types": fmt.Sprintf("%v", supportedTypes),
		},
		StatusCode: http.StatusUnsupportedMediaType,
	}
}

// MethodNotAllowedError handles HTTP method not allowed
func MethodNotAllowedError(method, endpoint string, allowedMethods []string) error {
	log.Warn(
		"HTTP method not allowed for endpoint",
		[]zap.Field{
			zap.String("error_code", "METHOD_NOT_ALLOWED"),
			zap.String("error_type", "http_method"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.String("method", method),
			zap.String("endpoint", endpoint),
			zap.Strings("allowed_methods", allowedMethods),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "METHOD_NOT_ALLOWED",
		Type:    "http_method",
		Message: fmt.Sprintf("method '%s' not allowed for endpoint '%s'", method, endpoint),
		Details: map[string]string{
			"method":          method,
			"endpoint":        endpoint,
			"allowed_methods": fmt.Sprintf("%v", allowedMethods),
		},
		StatusCode: http.StatusMethodNotAllowed,
	}
}

// NotFoundError handles resource not found
func NotFoundError(resource string) error {
	log.Info(
		"Resource not found",
		[]zap.Field{
			zap.String("error_code", "RESOURCE_NOT_FOUND"),
			zap.String("error_type", "routing"),
			zap.String("component", "http"),
			zap.String("severity", "low"),
			zap.String("resource", resource),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "RESOURCE_NOT_FOUND",
		Type:    "routing",
		Message: fmt.Sprintf("resource '%s' not found", resource),
		Details: map[string]string{
			"resource": resource,
		},
		StatusCode: http.StatusNotFound,
	}
}

// RateLimitExceededError handles rate limiting
func RateLimitExceededError(clientIP string, limit int, windowSeconds int) error {
	log.Warn(
		"Rate limit exceeded",
		[]zap.Field{
			zap.String("error_code", "RATE_LIMIT_EXCEEDED"),
			zap.String("error_type", "rate_limiting"),
			zap.String("component", "http"),
			zap.String("severity", "medium"),
			zap.String("client_ip", clientIP),
			zap.Int("limit", limit),
			zap.Int("window_seconds", windowSeconds),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "RATE_LIMIT_EXCEEDED",
		Type:    "rate_limiting",
		Message: fmt.Sprintf("rate limit exceeded: %d requests per %d seconds", limit, windowSeconds),
		Details: map[string]string{
			"client_ip":      clientIP,
			"limit":          fmt.Sprintf("%d", limit),
			"window_seconds": fmt.Sprintf("%d", windowSeconds),
		},
		StatusCode: http.StatusTooManyRequests,
	}
}

// RequestTimeoutError handles request timeout
func RequestTimeoutError(timeoutSeconds int) error {
	log.Warn(
		"Request timeout exceeded",
		[]zap.Field{
			zap.String("error_code", "REQUEST_TIMEOUT"),
			zap.String("error_type", "timeout"),
			zap.String("component", "http"),
			zap.String("severity", "medium"),
			zap.Int("timeout_seconds", timeoutSeconds),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "REQUEST_TIMEOUT",
		Type:    "timeout",
		Message: fmt.Sprintf("request timeout after %d seconds", timeoutSeconds),
		Details: map[string]string{
			"timeout_seconds": fmt.Sprintf("%d", timeoutSeconds),
		},
		StatusCode: http.StatusRequestTimeout,
	}
}

// ConflictError handles resource conflicts
func ConflictError(resource, reason string) error {
	log.Warn(
		"Resource conflict detected",
		[]zap.Field{
			zap.String("error_code", "RESOURCE_CONFLICT"),
			zap.String("error_type", "conflict"),
			zap.String("component", "http"),
			zap.String("severity", "medium"),
			zap.String("resource", resource),
			zap.String("reason", reason),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "RESOURCE_CONFLICT",
		Type:    "conflict",
		Message: fmt.Sprintf("conflict with resource '%s': %s", resource, reason),
		Details: map[string]string{
			"resource": resource,
			"reason":   reason,
		},
		StatusCode: http.StatusConflict,
	}
}

// ServiceUnavailableError handles service unavailability
func ServiceUnavailableError(service string, retryAfterSeconds int) error {
	log.Error(
		"Service unavailable",
		nil,
		[]zap.Field{
			zap.String("error_code", "SERVICE_UNAVAILABLE"),
			zap.String("error_type", "service"),
			zap.String("component", "http"),
			zap.String("severity", "high"),
			zap.String("service", service),
			zap.Int("retry_after_seconds", retryAfterSeconds),
		}...,
	)

	return &apperror.BaseApplicationError{
		Code:    "SERVICE_UNAVAILABLE",
		Type:    "service",
		Message: fmt.Sprintf("service '%s' is temporarily unavailable", service),
		Details: map[string]string{
			"service":             service,
			"retry_after_seconds": fmt.Sprintf("%d", retryAfterSeconds),
		},
		StatusCode: http.StatusServiceUnavailable,
	}
}
