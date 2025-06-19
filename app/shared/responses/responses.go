package responses

// ErrorResponse represents a generic error response
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SuccessResponse represents a generic success response
type SuccessResponse struct {
	Message string `json:"message"`
}
