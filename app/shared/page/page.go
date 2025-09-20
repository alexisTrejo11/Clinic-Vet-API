// Package page provides utilities for paginating items queries in Go applications.
// It offers a type-safe, generic approach to handle common pagination requirements
// including page metadata calculation, sorting direction, and empty page handling.
//
// The package includes:
// - Pagination input parameters with validation tags
// - Page metadata with comprehensive pagination information
// - Generic Page structure for type-safe paginated responses
// - Utility functions for creating pages and calculating metadata
package page

import (
	"clinic-vet-api/app/core/domain/specification"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// SortDirection defines the direction for sorting results.
type SortDirection string

const (
	ASC  SortDirection = "ASC"  // Ascending order
	DESC SortDirection = "DESC" // Descending order
)

func (sd *SortDirection) UnmarshalParam(param string) error {
	if param == "" {
		*sd = ASC // Valor por defecto
		return nil
	}

	// Manejar diferentes formatos
	switch strings.ToLower(param) {
	case "asc", "ascending", "0":
		*sd = ASC
	case "desc", "descending", "1":
		*sd = DESC
	default:
		if num, err := strconv.Atoi(param); err == nil {
			switch num {
			case 0:
				*sd = ASC
			case 1:
				*sd = DESC
			default:
				return fmt.Errorf("invalid sort direction: %s", param)
			}
		} else {
			return fmt.Errorf("invalid sort direction: %s", param)
		}
	}
	return nil
}

// PageInput (consider renaming to PaginationRequest) contains parameters
// for paginating and sorting query results.
// Uses JSON tags for API serialization and validate tags for input validation.
type PageInput struct {
	PageSize      int           `json:"page_limit" form:"page_limit" validate:"omitempty,gte=1,lte=100"`
	Page          int           `json:"page" form:"page" validate:"omitempty,gte=1"`
	SortDirection SortDirection `json:"sort_direction" form:"sort_direction" validate:"omitempty"`
	OrderBy       string        `json:"order_by" form:"order_by" validate:"omitempty,max=50"`
}

func (p PageInput) ToMap() map[string]any {
	return map[string]any{
		"page_limit":     p.PageSize,
		"page":           p.Page,
		"sort_direction": p.SortDirection,
		"order_by":       p.OrderBy,
	}
}

func (p PageInput) Validate() error {
	return nil
}
func FromSpecPagination(pagi specification.Pagination) PageInput {
	return PageInput{
		PageSize:      pagi.PageSize,
		Page:          pagi.Page,
		SortDirection: SortDirection(pagi.SortDir),
		OrderBy:       pagi.OrderBy,
	}
}

// SetDefaultsFieldsIfEmpty sets default values for PageInput fields if they are empty or invalid.
// Defaults:
// - SortDirection: ASC (if empty)
// - Page: 1 (if ≤ 0)
// - PageSize: 10 (if ≤ 0)
func (p *PageInput) SetDefaultsFieldsIfEmpty() {
	if p.SortDirection == "" {
		p.SortDirection = ASC
	}

	if p.Page <= 0 {
		p.Page = 1
	}

	if p.PageSize <= 0 {
		p.PageSize = 10
	}
}

func (p PageInput) Offset() int {
	return (p.Page - 1) * p.PageSize
}

// PageMetadata contains comprehensive information about the pagination state.
type PageMetadata struct {
	TotalCount      int           `json:"total_count"`       // Total number of items across all pages
	TotalPages      int           `json:"total_pages"`       // Total number of pages
	CurrentPage     int           `json:"current_page"`      // Current page number
	PageSize        int           `json:"page_limit"`        // Number of items per page
	SortDirection   SortDirection `json:"sort_direction"`    // Sorting direction applied
	HasNextPage     bool          `json:"has_next_page"`     // True if another page exists after current
	HasPreviousPage bool          `json:"has_previous_page"` // True if a page exists before current
}

// Page represents a paginated response containing items and metadata.
// Uses generics to provide type safety for the items field.
type Page[T any] struct {
	Items    []T          `json:"items"`    // The paginated items slice
	Metadata PageMetadata `json:"metadata"` // Pagination metadata
}

// NewPage creates a new Page instance with the provided items and metadata.
func NewPage[T any](items []T, metadata PageMetadata) Page[T] {
	page := &Page[T]{
		Items:    items,
		Metadata: metadata,
	}

	return *page
}

// EmptyPage creates an empty Page instance with properly initialized empty items.
// Handles slice types specially to ensure they're non-nil empty slices.
func EmptyPage[T any]() Page[T] {
	var emptyItems []T
	// Ensure emptyItems is a non-nil empty slice if T is a slice type
	if reflect.TypeOf(emptyItems).Kind() == reflect.Slice {
		emptyItems = make([]T, 0)
	}

	return Page[T]{
		Items: emptyItems,
		Metadata: PageMetadata{
			TotalCount:      0,
			TotalPages:      0,
			CurrentPage:     1,
			PageSize:        0,
			SortDirection:   ASC,
			HasNextPage:     false,
			HasPreviousPage: false,
		},
	}
}

// GetPageMetadata calculates pagination metadata based on total items and page input.
// Returns a PageMetadata pointer with all pagination information populated.
func GetPageMetadata(totalItems int, page PageInput) *PageMetadata {
	var totalPages int
	if page.PageSize > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(page.PageSize)))
	} else {
		totalPages = 1
	}

	return &PageMetadata{
		TotalCount:      totalItems,
		TotalPages:      totalPages,
		CurrentPage:     page.Page,
		PageSize:        page.PageSize,
		SortDirection:   page.SortDirection,
		HasNextPage:     page.Page < totalPages,
		HasPreviousPage: page.Page > 1,
	}
}
