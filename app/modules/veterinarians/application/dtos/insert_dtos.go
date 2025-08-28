package vetDtos

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

// VetCreate represents the data transfer object for creating a new veterinarian.
type VetCreate struct {
	// The veterinarian's first name.
	FirstName string `json:"first_name" validate:"required"`
	// The veterinarian's last name.
	LastName string `json:"last_name" validate:"required"`
	// URL for the veterinarian's photo.
	Photo string `json:"photo"`
	// The unique license number of the veterinarian.
	LicenseNumber string `json:"license_number" validate:"required"`
	// The number of years of professional experience.
	YearsExperience int `json:"years_experience"`
	// Indicates if the veterinarian is currently active.
	IsActive bool `json:"is_active"`
	// The veterinarian's medical specialty.
	Specialty vetDomain.VetSpecialty `json:"specialty"`
	// The fee for a consultation.
	ConsultationFee *shared.Money `json:"consultation_fee"`
	// The working schedule of the veterinarian.
	LaboralSchedule []ScheduleInsert `json:"laboral_schedule"`
}

// ScheduleInsert represents a working day schedule for a veterinarian.
type ScheduleInsert struct {
	// The day of the week for the schedule. (e.g., "Monday", "Tuesday")
	Day string `json:"day"`
	// The start time of the work day in hours (e.g., 8 for 8:00 AM).
	EntryTime int `json:"entry_time"`
	// The departure time of the work day in hours.
	DepartureTime int `json:"departure_time"`
	// The start time of the break in hours.
	StartBreak int `json:"start_break"`
	// The end time of the break in hours.
	EndBreak int `json:"end_break"`
}

// VetUpdate represents the data transfer object for updating a veterinarian.
type VetUpdate struct {
	// The veterinarian's first name.
	FirstName *string `json:"first_name"`
	// The veterinarian's last name.
	LastName *string `json:"last_name"`
	// URL for the veterinarian's photo.
	Photo *string `json:"photo"`
	// The unique license number of the veterinarian.
	LicenseNumber *string `json:"license_number"`
	// The number of years of professional experience.
	YearsExperience *int `json:"years_experience"`
	// The veterinarian's medical specialty.
	Specialty *vetDomain.VetSpecialty `json:"specialty"`
	// Indicates if the veterinarian is currently active.
	IsActive *bool `json:"is_active"`
	// The fee for a consultation.
	ConsultationFee *shared.Money `json:"consultation_fee"`
	// The working schedule of the veterinarian.
	LaboralSchedule *[]ScheduleInsert `json:"laboral_schedule"`
}

// VetSearchParams defines the parameters for searching and filtering veterinarians.
type VetSearchParams struct {
	// Embedded page data for pagination.
	page.PageData
	// Filters to apply to the search results.
	Filters VetFilters `json:"filters"`
	// The field to order the results by.
	OrderBy VetOrderBy
}

// VetFilters defines the available filters for the veterinarian search.
type VetFilters struct {
	// Filter by the veterinarian's name.
	Name *string `json:"name"`
	// Filter by the veterinarian's license number.
	LicenseNumber *string `json:"license_number"`
	// Filter by the veterinarian's medical specialty.
	Specialty *vetDomain.VetSpecialty `json:"specialty"`
	// Filter by a range of years of experience.
	YearsExperience *struct {
		Min *int `json:"min"`
		Max *int `json:"max"`
	} `json:"years_experience"`
	// Filter by active status.
	IsActive *bool `json:"is_active"`
}

// VetOrderBy is a type for defining the field to order the results.
// @Enum name specialty years_experience created_at
type VetOrderBy string

const (
	// OrderByName orders the results by the veterinarian's name.
	OrderByName VetOrderBy = "name"
	// OrderBySpecialty orders the results by specialty.
	OrderBySpecialty VetOrderBy = "specialty"
	// OrderByYearsExperience orders the results by years of experience.
	OrderByYearsExperience VetOrderBy = "years_experience"
	// OrderByCreatedAt orders the results by creation date.
	OrderByCreatedAt VetOrderBy = "created_at"
)
