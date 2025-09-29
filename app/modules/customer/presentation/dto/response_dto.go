package dto

import (
	"time"

	"clinic-vet-api/app/modules/customer/application/handler"
)

// CustomerResponse represents the customer response
// @Description Customer information response
type CustomerResponse struct {
	// Customer ID
	ID uint `json:"id" example:"cust_123456"`

	// Customer's first name
	FirstName string `json:"first_name" example:"John"`

	// Customer's last name
	LastName string `json:"last_name" example:"Doe"`

	// Customer's gender
	Gender string `json:"gender" example:"male"`

	// Customer's date of birth
	DateOfBirth string `json:"date_of_birth" example:"1990-01-15T00:00:00Z"`

	PetCount int `json:"pet_count" example:"3"`

	// Whether the customer is active
	IsActive bool `json:"is_active" example:"true"`

	// Creation timestamp
	CreatedAt string `json:"created_at" example:"2024-01-01T12:00:00Z"`

	// Last update timestamp
	UpdatedAt string `json:"updated_at" example:"2024-01-15T14:30:00Z"`
}

func FromResult(result handler.CustomerResult) *CustomerResponse {
	return &CustomerResponse{
		ID:          result.ID.Value(),
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		Gender:      result.Gender.DisplayName(),
		DateOfBirth: result.DateOfBirth.Format(time.DateOnly),
		PetCount:    result.PetsCount,
		IsActive:    result.IsActive,
		CreatedAt:   result.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   result.UpdatedAt.Format(time.RFC3339),
	}
}

func FromResultList(results []handler.CustomerResult) []CustomerResponse {
	if len(results) == 0 {
		return []CustomerResponse{}
	}

	response := make([]CustomerResponse, len(results))
	for i, res := range results {
		response[i] = *FromResult(res)
	}
	return response
}
