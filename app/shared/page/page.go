package page

import "math"

type SortDirection string

const ASC SortDirection = "ASC"
const DESC SortDirection = "DESC"

type PageData struct {
	PageSize      int           `json:"page_limit" validate:"omitempty,min=1,max=100"`
	PageNumber    int           `json:"page_number" validate:"omitempty,min=1"`
	SortDirection SortDirection `json:"sort_direction" validate:"omitempty,min=0"`
}

type PageMetadata struct {
	TotalCount      int           `json:"total_count"`
	TotalPages      int           `json:"total_pages"`
	CurrentPage     int           `json:"current_page"`
	PageSize        int           `json:"page_limit"`
	SortDirection   SortDirection `json:"sort_direction"`
	HasNextPage     bool          `json:"has_next_page"`
	HasPreviousPage bool          `json:"has_previous_page"`
}

type Page[T any] struct {
	Data     T            `json:"data"`
	Metadata PageMetadata `json:"metadata"`
}

func NewPage[T any](data T, metadata PageMetadata) *Page[T] {
	return &Page[T]{
		Data:     data,
		Metadata: metadata,
	}
}

func EmptyPage[T any]() Page[T] {
	return Page[T]{
		Data:     *new(T),
		Metadata: PageMetadata{},
	}
}

func GetPageMetadata(totalItems int, page PageData) *PageMetadata {
	var totalPages int
	if page.PageSize > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(page.PageSize)))
	} else {
		totalPages = 1
	}

	return &PageMetadata{
		TotalCount:      totalItems,
		TotalPages:      totalPages,
		CurrentPage:     page.PageNumber,
		PageSize:        page.PageSize,
		SortDirection:   page.SortDirection,
		HasNextPage:     page.PageNumber < totalPages,
		HasPreviousPage: page.PageNumber > 1,
	}
}
