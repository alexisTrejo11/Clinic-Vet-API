package dto

import (
	"clinic-vet-api/app/modules/users/application/usecase/query"
	"time"
)

// UserResponse represents a user API response
// @Description User information response structure
type UserResponse struct {
	// Unique identifier for the user
	// Required: true
	// Example: "123456"
	ID uint `json:"id"`

	// User's phone number in E.164 format
	// Required: false
	// Example: "+1234567890"
	PhoneNumber string `json:"phone_number,omitempty"`

	// User's email address
	// Required: true
	// Format: email
	// Example: "user@example.com"
	Email string `json:"email"`

	// User's role in the system
	// Required: true
	// Enum: admin,employee,customer
	// Example: "customer"
	Role string `json:"role"`

	// Current status of the user account
	// Required: true
	// Enum: active,inactive,pending,suspended
	// Example: "active"
	Status string `json:"status"`

	// Date and time when the user joined the system
	// Required: true
	// Format: date-time
	// Example: "2023-01-15T10:30:00Z"
	JoinedAt time.Time `json:"joined_at"`

	// Date and time of the user's last login
	// Required: false
	// Format: date-time
	// Example: "2023-12-01T15:45:00Z"
	LastLoginAt *time.Time `json:"last_login_at,omitempty"`

	CustomerID *uint `json:"customer_id,omitempty"`
	EmployeeID *uint `json:"employee_id,omitempty"`
}

func UserResultToResponse(userResult *query.UserResult) *UserResponse {
	if userResult == nil {
		return nil
	}

	return &UserResponse{
		ID:          userResult.ID,
		PhoneNumber: userResult.PhoneNumber,
		Email:       userResult.Email,
		Role:        userResult.Role,
		Status:      userResult.Status,
		JoinedAt:    userResult.JoinedAt,
		LastLoginAt: userResult.LastLoginAt,
		CustomerID:  userResult.CustomerID,
		EmployeeID:  userResult.EmployeeID,
	}
}

func UserResultsToResponses(userResults []query.UserResult) []UserResponse {
	if len(userResults) < 1 {
		return []UserResponse{}
	}

	userResponses := make([]UserResponse, 0, len(userResults))
	for _, userResult := range userResults {
		userResponses = append(userResponses, *UserResultToResponse(&userResult))
	}

	return userResponses
}
