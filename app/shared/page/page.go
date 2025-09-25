// Package page provides utilities for paginating items queries in Go applications.
// It offers a type-safe, generic approach to handle common pagination requirements
// including page metadata calculation, sorting direction, and empty page handling.
package page

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"fmt"
	"math"
	"strings"
)

// SortDirection defines the direction for sorting results.
type SortDirection string

const (
	ASC  SortDirection = "ASC"  // Ascending order
	DESC SortDirection = "DESC" // Descending order
)

// Default values for pagination
const (
	DefaultPage     = 1
	DefaultPageSize = 10
	MaxPageSize     = 100
)

// UnmarshalParam unmarshals SortDirection from various string formats
func (sd *SortDirection) UnmarshalParam(param string) error {
	if param == "" {
		*sd = ASC // Default value
		return nil
	}

	switch strings.ToLower(strings.TrimSpace(param)) {
	case "asc", "ascending", "0":
		*sd = ASC
	case "desc", "descending", "1":
		*sd = DESC
	default:
		return fmt.Errorf("invalid sort direction: %s", param)
	}
	return nil
}

// IsValid checks if the SortDirection is valid
func (sd SortDirection) IsValid() bool {
	return sd == ASC || sd == DESC
}

// PaginationRequest (better name than PaginationRequest) contains parameters for paginating and sorting query results.
type PaginationRequest struct {
	Page          int           `json:"page" form:"page" validate:"omitempty,gte=1"`
	PageSize      int           `json:"page_size" form:"page_size" validate:"omitempty,gte=1,lte=100"`
	SortDirection SortDirection `json:"sort_direction" form:"sort_direction" validate:"omitempty,oneof=ASC DESC"`
	OrderBy       string        `json:"order_by" form:"order_by" validate:"omitempty,max=50"`
}

// NewPaginationRequest creates a new PaginationRequest with default values
func NewPaginationRequest() PaginationRequest {
	return PaginationRequest{
		Page:          DefaultPage,
		PageSize:      DefaultPageSize,
		SortDirection: ASC,
		OrderBy:       "",
	}
}

// Offset calculates the database offset based on page and page size
func (p PaginationRequest) Offset() int {
	if p.Page <= 1 {
		return 0
	}
	return (p.Page - 1) * p.PageSize
}

// Limit returns the page size (alias for better semantics)
func (p PaginationRequest) Limit() int {
	return p.PageSize
}

// IsValid checks if the pagination request is valid
func (p PaginationRequest) IsValid() bool {
	return p.Page >= 1 && p.PageSize >= 1 && p.PageSize <= MaxPageSize && p.SortDirection.IsValid()
}

// ToMap converts the pagination request to a map
func (p PaginationRequest) ToMap() map[string]any {
	return map[string]any{
		"page":           p.Page,
		"page_size":      p.PageSize,
		"sort_direction": p.SortDirection,
		"order_by":       p.OrderBy,
	}
}

// ToSpecPagination converts to specification.Pagination
func (p PaginationRequest) ToSpecPagination() specification.Pagination {
	return specification.Pagination{
		Limit:   p.Limit(),
		Offset:  p.Offset(),
		SortDir: string(p.SortDirection),
		OrderBy: p.OrderBy,
	}
}

// FromSpecPagination creates a PaginationRequest from specification.Pagination
func FromSpecPagination(pagi specification.Pagination) PaginationRequest {
	// Calculate page from offset and limit
	page := DefaultPage
	if pagi.Limit > 0 && pagi.Offset >= 0 {
		page = (pagi.Offset / pagi.Limit) + 1
	}

	sortDir := ASC
	if pagi.SortDir != "" {
		if err := sortDir.UnmarshalParam(pagi.SortDir); err != nil {
			sortDir = ASC // Fallback to default
		}
	}

	return PaginationRequest{
		Page:          page,
		PageSize:      pagi.Limit,
		SortDirection: sortDir,
		OrderBy:       pagi.OrderBy,
	}
}

// WithDefaults returns a new PaginationRequest with default values applied where needed
func (p PaginationRequest) WithDefaults() PaginationRequest {
	result := p

	if result.Page < 1 {
		result.Page = DefaultPage
	}

	if result.PageSize < 1 {
		result.PageSize = DefaultPageSize
	} else if result.PageSize > MaxPageSize {
		result.PageSize = MaxPageSize
	}

	if !result.SortDirection.IsValid() {
		result.SortDirection = ASC
	}

	return result
}

// PageMetadata contains comprehensive information about the pagination state.
type PageMetadata struct {
	TotalCount      int           `json:"total_count"`       // Total number of items across all pages
	TotalPages      int           `json:"total_pages"`       // Total number of pages
	CurrentPage     int           `json:"current_page"`      // Current page number
	PageSize        int           `json:"page_size"`         // Number of items per page
	SortDirection   SortDirection `json:"sort_direction"`    // Sorting direction applied
	HasNextPage     bool          `json:"has_next_page"`     // True if another page exists after current
	HasPreviousPage bool          `json:"has_previous_page"` // True if a page exists before current
}

// Page represents a paginated response containing items and metadata.
type Page[T any] struct {
	Items    []T          `json:"items"`    // The paginated items slice
	Metadata PageMetadata `json:"metadata"` // Pagination metadata
}

// NewPage creates a new Page instance with the provided items and metadata.
func NewPage[T any](items []T, totalCount int, request PaginationRequest) Page[T] {
	metadata := CalculateMetadata(totalCount, request)

	return Page[T]{
		Items:    items,
		Metadata: metadata,
	}
}

// NewNewEmptyPage creates an empty Page instance with properly initialized empty items.
func NewEmptyPage[T any](request PaginationRequest) Page[T] {
	emptyItems := make([]T, 0)
	metadata := CalculateMetadata(0, request)

	return Page[T]{
		Items:    emptyItems,
		Metadata: metadata,
	}
}

// CalculateMetadata calculates pagination metadata based on total items and pagination request.
func CalculateMetadata(totalCount int, request PaginationRequest) PageMetadata {
	req := request.WithDefaults()

	totalPages := 1
	if req.PageSize > 0 {
		totalPages = int(math.Ceil(float64(totalCount) / float64(req.PageSize)))
	}

	if totalPages < 1 {
		totalPages = 1
	}

	return PageMetadata{
		TotalCount:      totalCount,
		TotalPages:      totalPages,
		CurrentPage:     req.Page,
		PageSize:        req.PageSize,
		SortDirection:   req.SortDirection,
		HasNextPage:     req.Page < totalPages,
		HasPreviousPage: req.Page > 1,
	}
}

// FirstPage returns a PaginationRequest for the first page
func FirstPage(pageSize int) PaginationRequest {
	if pageSize <= 0 {
		pageSize = DefaultPageSize
	}
	return PaginationRequest{
		Page:     1,
		PageSize: pageSize,
	}
}

// AllItems returns a PaginationRequest to get all items (no pagination)
func AllItems() PaginationRequest {
	return PaginationRequest{
		Page:     1,
		PageSize: math.MaxInt32, // Use a large number to get all items
	}
}

// MapItems maps the items of a Page[T] to a Page[R] using the provided mapper function.
func MapItems[T any, R any](p Page[T], mapper func(T) R) Page[R] {
	newItems := make([]R, len(p.Items))
	for i, item := range p.Items {
		newItems[i] = mapper(item)
	}

	return Page[R]{
		Items:    newItems,
		Metadata: p.Metadata,
	}
}
