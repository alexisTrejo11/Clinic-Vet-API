package dto

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/customer/application/query"
	"clinic-vet-api/app/shared/page"
	"time"
)

// CustomerSearchQuery represents query parameters for customer search
// @Description Query parameters for filtering and paginating customers
type CustomerSearchQuery struct {
	// Filter by customer name (partial match)
	Name *string `form:"name" example:"John"`

	// Filter by customer last name (partial match)
	LastName *string `form:"last_name" example:"Doe"`

	// Filter by customer email
	Email *string `form:"email" example:"john.doe@example.com"`

	// Filter by phone number (partial match)
	PhoneNumber *string `form:"phone_number" example:"+1234567890"`

	// Filter by gender
	Gender *enum.PersonGender `form:"gender" example:"male"`

	// Filter by date of birth range - start date (ISO 8601)
	DateOfBirthFrom *time.Time `form:"date_of_birth_from" example:"1990-01-01T00:00:00Z"`

	// Filter by date of birth range - end date (ISO 8601)
	DateOfBirthTo *time.Time `form:"date_of_birth_to" example:"2000-12-31T23:59:59Z"`

	// Filter by active status (true/false)
	IsActive *bool `form:"is_active" example:"true"`

	// Filter by creation date range - start date (ISO 8601)
	CreatedAtFrom *time.Time `form:"created_at_from" example:"2024-01-01T00:00:00Z"`

	// Filter by creation date range - end date (ISO 8601)
	CreatedAtTo *time.Time `form:"created_at_to" example:"2024-12-31T23:59:59Z"`

	// Filter by minimum number of pets
	MinPets *int `form:"min_pets" example:"1"`

	// Filter by maximum number of pets
	MaxPets *int `form:"max_pets" example:"5"`

	// Filter by customer ID
	CustomerID *uint `form:"id" example:"123"`

	// Filter by associated user ID
	UserID *uint `form:"user_id" example:"456"`
	page.PaginationRequest
}

func (q *CustomerSearchQuery) ToSpecification() specification.CustomerSpecification {
	var ID *valueobject.CustomerID
	if q.CustomerID != nil {
		id := valueobject.NewCustomerID(*q.CustomerID)
		ID = &id
	}

	var userID *valueobject.UserID
	if q.UserID != nil {
		uid := valueobject.NewUserID(*q.UserID)
		userID = &uid
	}

	spec := &specification.CustomerSpecification{
		ID:              ID,
		Name:            q.Name,
		LastName:        q.LastName,
		Email:           q.Email,
		PhoneNumber:     q.PhoneNumber,
		Gender:          q.Gender,
		DateOfBirthFrom: q.DateOfBirthFrom,
		DateOfBirthTo:   q.DateOfBirthTo,
		IsActive:        q.IsActive,
		CreatedAtFrom:   q.CreatedAtFrom,
		CreatedAtTo:     q.CreatedAtTo,
		MinPets:         q.MinPets,
		MaxPets:         q.MaxPets,
		UserID:          userID,
	}

	return *spec
}

func (q *CustomerSearchQuery) ToQuery() (query.FindCustomerBySpecificationQuery, error) {
	spec := q.ToSpecification()
	return query.NewFindCustomerBySpecificationQuery(spec)
}
