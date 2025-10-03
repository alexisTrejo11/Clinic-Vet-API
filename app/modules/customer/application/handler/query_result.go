package handler

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/customer"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type CustomerResult struct {
	ID          valueobject.CustomerID
	FirstName   string
	LastName    string
	Gender      enum.PersonGender
	DateOfBirth time.Time
	Photo       string
	UserID      *valueobject.UserID
	IsActive    bool
	PetsCount   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func customerToResult(customer customer.Customer) CustomerResult {
	result := &CustomerResult{
		ID:          customer.ID(),
		PetsCount:   len(customer.Pets()),
		IsActive:    customer.IsActive(),
		Gender:      customer.Gender(),
		UserID:      customer.UserID(),
		Photo:       customer.Photo(),
		FirstName:   customer.FirstName(),
		LastName:    customer.LastName(),
		DateOfBirth: customer.DateOfBirth(),
		UpdatedAt:   customer.UpdatedAt(),
		CreatedAt:   customer.CreatedAt(),
	}
	return *result
}
