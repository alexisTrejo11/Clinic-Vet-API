// Package dto contains data transfer objects for the veterinarians module.
package dto

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// CreateVetRequest represents the data transfer object for creating a new veterinarian.
type CreateVetRequest struct {
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
	Specialty enum.VetSpecialty `json:"specialty"`
	// The fee for a consultation.
	ConsultationFee *valueobject.Money `json:"consultation_fee"`
	// The working schedule of the veterinarian.
	LaboralSchedule []ScheduleRequest `json:"laboral_schedule"`
}

func (r *CreateVetRequest) ToCreateData() (dto.CreateVetData, error) {
	return dto.CreateVetData{
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		Photo:           r.Photo,
		LicenseNumber:   r.LicenseNumber,
		YearsExperience: r.YearsExperience,
		IsActive:        r.IsActive,
		Specialty:       string(r.Specialty),
		ConsultationFee: r.ConsultationFee,
	}, nil
}

// ScheduleRequest represents a working day schedule for a veterinarian.
type ScheduleRequest struct {
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

// UpdatePetRequest represents the data transfer object for updating a veterinarian.
type UpdateVetRequest struct {
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
	Specialty *enum.VetSpecialty `json:"specialty"`
	// Indicates if the veterinarian is currently active.
	IsActive *bool `json:"is_active"`
	// The fee for a consultation.
	ConsultationFee *valueobject.Money `json:"consultation_fee"`
	// The working schedule of the veterinarian.
	LaboralSchedule *[]ScheduleRequest `json:"laboral_schedule"`
}

func (r *UpdateVetRequest) ToUpdateData(vetIDInt int) (dto.UpdateVetData, error) {
	vetID, err := valueobject.NewVetID(vetIDInt)
	if err != nil {
		return dto.UpdateVetData{}, err
	}

	return dto.UpdateVetData{
		VetID:           vetID,
		FirstName:       r.FirstName,
		LastName:        r.LastName,
		Photo:           r.Photo,
		LicenseNumber:   r.LicenseNumber,
		YearsExperience: r.YearsExperience,
		Specialty: func() *string {
			if r.Specialty != nil {
				s := string(*r.Specialty)
				return &s
			}
			return nil
		}(),
		IsActive:        r.IsActive,
		ConsultationFee: r.ConsultationFee,
	}, nil
}

// VetSearchParams defines the parameters for searching and filtering veterinarians.
type VetSearchParams struct {
	// Embedded page data for pagination.
	page.PageInput
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
	Specialty *enum.VetSpecialty `json:"specialty"`
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
