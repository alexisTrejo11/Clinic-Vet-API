package apiResponse

import (
	"time"

	custom_error "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/errors"
)

type APIResponse struct {
	Success   bool        `json:"success"`
	Data      interface{} `json:"data,omitempty"`
	Error     *ErrorInfo  `json:"error,omitempty"`
	Meta      interface{} `json:"meta,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
	RequestID string      `json:"request_id,omitempty"`
}

type ErrorInfo struct {
	Code    string                 `json:"code"`
	Type    string                 `json:"type"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
	Stack   string                 `json:"stack,omitempty"` // Develop Only
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

type PaginationMeta struct {
	CurrentPage int  `json:"current_page"`
	PageSize    int  `json:"page_size"`
	Total       int  `json:"total"`
	TotalPages  int  `json:"total_pages"`
	HasNext     bool `json:"has_next"`
	HasPrev     bool `json:"has_prev"`
}

func (r *APIResponse) SuccessRequest(data interface{}) {
	r.Success = true
	r.Data = data
	r.Timestamp = time.Now()
}

func (r *APIResponse) SuccessWithMeta(data interface{}, meta interface{}) *APIResponse {
	r.Success = true
	r.Data = data
	r.Meta = meta
	r.Timestamp = time.Now()
	return r
}

func (r *APIResponse) SuccessWithPagination(data interface{}, pagination interface{}) *APIResponse {
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
	case custom_error.ApplicationError: // Or CASE INFRA
		errorInfo.Code = e.ErrorCode()
		errorInfo.Type = e.ErrorType()
		errorInfo.Message = e.Error()
		errorInfo.Details = e.Details()
	case custom_error.DomainError:
		errorInfo.Code = e.ErrorCode()
		errorInfo.Type = e.ErrorType()
		errorInfo.Message = e.Error()
		errorInfo.Details = e.Details()
	default:
		errorInfo.Message = err.Error()
	}

	return &APIResponse{
		Success:   false,
		Error:     errorInfo,
		Timestamp: time.Now(),
	}
}
