package specification

import (
	"strings"
	"time"

	"clinic-vet-api/app/core/domain/entity/customer"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
)

type CustomerSpecification struct {
	// Filter by customer ID
	ID *valueobject.CustomerID

	// Filter by customer name (partial match)
	Name *string

	// Filter by customer last name (partial match)
	LastName *string

	// Filter by customer email
	Email *string

	// Filter by phone number (partial match)
	PhoneNumber *string

	// Filter by user ID associated with the customer
	UserID *valueobject.UserID

	// Filter by gender
	Gender *enum.PersonGender

	// Filter by date of birth range - start date
	DateOfBirthFrom *time.Time

	// Filter by date of birth range - end date
	DateOfBirthTo *time.Time

	// Filter by active status
	IsActive *bool

	// Filter by creation date range - start date
	CreatedAtFrom *time.Time

	// Filter by creation date range - end date
	CreatedAtTo *time.Time

	// Filter by minimum number of pets
	MinPets *int

	// Filter by maximum number of pets
	MaxPets *int

	// Pagination parameters
	Pagination
}

func (s *CustomerSpecification) IsSatisfiedBy(entity any) bool {
	customer, ok := entity.(customer.Customer)
	if !ok {
		return false
	}

	// Aplicar todos los filtros
	if s.Name != nil && !strings.Contains(strings.ToLower(customer.Name().FirstName), strings.ToLower(*s.Name)) {
		return false
	}

	if s.LastName != nil && !strings.Contains(strings.ToLower(customer.Name().LastName), strings.ToLower(*s.LastName)) {
		return false
	}

	if s.IsActive != nil && customer.IsActive() != *s.IsActive {
		return false
	}

	if s.Gender != nil && customer.Gender() != *s.Gender {
		return false
	}

	if s.DateOfBirthFrom != nil && customer.DateOfBirth().Before(*s.DateOfBirthFrom) {
		return false
	}

	if s.DateOfBirthTo != nil && customer.DateOfBirth().After(*s.DateOfBirthTo) {
		return false
	}

	if s.CreatedAtFrom != nil && customer.CreatedAt().Before(*s.CreatedAtFrom) {
		return false
	}

	if s.CreatedAtTo != nil && customer.CreatedAt().After(*s.CreatedAtTo) {
		return false
	}

	return true
}

func (s *CustomerSpecification) ToSQL() (string, []any) {
	var conditions []string
	var args []any

	if s.Name != nil {
		conditions = append(conditions, "LOWER(first_name) LIKE LOWER(?)")
		args = append(args, "%"+*s.Name+"%")
	}

	if s.LastName != nil {
		conditions = append(conditions, "LOWER(last_name) LIKE LOWER(?)")
		args = append(args, "%"+*s.LastName+"%")
	}

	if s.Email != nil {
		conditions = append(conditions, "email = ?")
		args = append(args, *s.Email)
	}

	if s.PhoneNumber != nil {
		conditions = append(conditions, "phone_number LIKE ?")
		args = append(args, "%"+*s.PhoneNumber+"%")
	}

	if s.IsActive != nil {
		conditions = append(conditions, "is_active = ?")
		args = append(args, *s.IsActive)
	}

	if s.Gender != nil {
		conditions = append(conditions, "gender = ?")
		args = append(args, string(*s.Gender))
	}

	if s.DateOfBirthFrom != nil {
		conditions = append(conditions, "date_of_birth >= ?")
		args = append(args, *s.DateOfBirthFrom)
	}

	if s.DateOfBirthTo != nil {
		conditions = append(conditions, "date_of_birth <= ?")
		args = append(args, *s.DateOfBirthTo)
	}

	if s.CreatedAtFrom != nil {
		conditions = append(conditions, "created_at >= ?")
		args = append(args, *s.CreatedAtFrom)
	}

	if s.CreatedAtTo != nil {
		conditions = append(conditions, "created_at <= ?")
		args = append(args, *s.CreatedAtTo)
	}

	if s.MinPets != nil {
		conditions = append(conditions, "pet_count >= ?")
		args = append(args, *s.MinPets)
	}

	if s.MaxPets != nil {
		conditions = append(conditions, "pet_count <= ?")
		args = append(args, *s.MaxPets)
	}

	if len(conditions) == 0 {
		return "1 = 1", args // Retorna condición verdadera si no hay filtros
	}

	sql := strings.Join(conditions, " AND ")
	return sql, args
}
