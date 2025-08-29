package controller

// PaginationMetadata represents metadata for paginated responses
type PaginationMetadata struct {
	CurrentPage  int `json:"current_page" example:"1"`
	PageSize     int `json:"page_size" example:"10"`
	TotalPages   int `json:"total_pages" example:"5"`
	TotalRecords int `json:"total_records" example:"42"`
}
