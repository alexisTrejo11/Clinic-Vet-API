package query

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/customer"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

// CustomerResult represents the customer data from repository/domain layer
// @Description Customer data structure for intermediate mapping between domain and presentation layers
type CustomerResult struct {
	ID          valueobject.CustomerID `json:"id"`
	FirstName   string                 `json:"first_name"`
	LastName    string                 `json:"last_name"`
	Gender      enum.PersonGender      `json:"gender"`
	DateOfBirth time.Time              `json:"date_of_birth"`
	Photo       string                 `json:"photo,omitempty"`
	UserID      *valueobject.UserID    `json:"user_id"`
	IsActive    bool                   `json:"is_active"`
	PetsCount   int                    `json:"pets_count"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

// ToResponse convierte CustomerResult a CustomerResponse
func FromEntityToResult(customer customer.Customer) *CustomerResult {
	return &CustomerResult{
		ID:          customer.ID(),
		PetsCount:   len(customer.Pets()),
		IsActive:    customer.IsActive(),
		CreatedAt:   customer.CreatedAt(),
		UpdatedAt:   customer.UpdatedAt(),
		UserID:      customer.UserID(),
		Photo:       customer.Photo(),
		FirstName:   customer.FullName().FirstName,
		LastName:    customer.FullName().LastName,
		DateOfBirth: customer.DateOfBirth(),
	}
}

func FromEntityListToResultList(customers []customer.Customer) []CustomerResult {
	results := make([]CustomerResult, len(customers))
	for i, cust := range customers {
		results[i] = *FromEntityToResult(cust)
	}
	return results
}
