// Package response contains all the centralized implementation of data strucutures and logic to output data, messages and status for this project
package response

import (
	"time"

	domainerr "clinic-vet-api/app/core/error"
	apperror "clinic-vet-api/app/shared/error/application"
)

// @Description Standardized API response structure.
type APIResponse struct {
	// Indicates whether the request was successful.
	Success bool `json:"success"`
	// The data payload of the response for successful requests.
	Data any `json:"data,omitempty"`
	// A message providing additional context about the response.
	Message string `json:"message,omitempty"`
	// Details of the error if the request was not successful.
	Error *ErrorInfo `json:"error,omitempty"`
	// Additional metadata for the response, such as pagination info.
	Meta any `json:"meta,omitempty"`
	// The timestamp when the response was generated.
	Timestamp time.Time `json:"timestamp"`
	// A unique identifier for the request.
	RequestID string `json:"request_id,omitempty"`
}

// @Description Detailed information about an error.
type ErrorInfo struct {
	// The error code.
	Code string `json:"code"`
	// The type of the error.
	Type string `json:"type"`
	// A descriptive error message.
	Message string `json:"message"`
	// A map of specific error details, often used for validation errors.
	Details map[string]string `json:"details,omitempty"`
	// The error stack trace (for development purposes).
	Stack string `json:"stack,omitempty"` // Develop Only
}

// @Description Metadata for a generic list response.
type Meta struct {
	// The current page number.
	Page int `json:"page,omitempty"`
	// The number of items per page.
	PageSize int `json:"page_size,omitempty"`
	// The total number of items.
	Total int `json:"total,omitempty"`
	// The total number of pages.
	TotalPages int `json:"total_pages,omitempty"`
}

// @Description Detailed pagination metadata.
type PaginationMeta struct {
	// The current page number.
	CurrentPage int `json:"current_page"`
	// The number of items per page.
	PageSize int `json:"page_size"`
	// The total number of items.
	Total int `json:"total"`
	// The total number of pages.
	TotalPages int `json:"total_pages"`
	// Indicates if there is a next page.
	HasNext bool `json:"has_next"`
	// Indicates if there is a previous page.
	HasPrev bool `json:"has_prev"`
}

func (r *APIResponse) SuccessRequest(data any, message string) {
	r.Success = true
	r.Data = data
	r.Message = message
	r.Timestamp = time.Now()
}

func (r *APIResponse) SuccessWithMeta(data any, meta any, message string) *APIResponse {
	r.Success = true
	r.Message = message
	r.Data = data
	r.Meta = meta
	r.Timestamp = time.Now()
	return r
}

func (r *APIResponse) SuccessWithPagination(data any, pagination any) *APIResponse {
	meta := &Meta{
		Page:       pagination.(PaginationMeta).CurrentPage,
		PageSize:   pagination.(PaginationMeta).PageSize,
		Total:      pagination.(PaginationMeta).Total,
		TotalPages: pagination.(PaginationMeta).TotalPages,
	}

	return &APIResponse{
		Success:   true,
		Data:      data,
		Meta:      meta,
		Timestamp: time.Now(),
	}
}

func (r *APIResponse) ErrorRequest(err error) *APIResponse {
	errorInfo := &ErrorInfo{
		Code:    "INTERNAL_SERVER_ERROR",
		Type:    "application",
		Message: "An unexpected error occurred",
	}

	switch e := err.(type) {
	case apperror.BaseApplicationError: // Or CASE INFRA
		errorInfo.Code = e.ErrorCode()
		errorInfo.Type = e.ErrorType()
		errorInfo.Message = e.Error()
		errorInfo.Details = e.DetailMap()
	case domainerr.BaseDomainError:
		errorInfo.Code = e.ErrorCode()
		errorInfo.Type = e.ErrorType()
		errorInfo.Message = e.Error()
		errorInfo.Details = e.DetailMap()
	default:
		errorInfo.Message = err.Error()
	}

	return &APIResponse{
		Success:   false,
		Error:     errorInfo,
		Timestamp: time.Now(),
	}
}
