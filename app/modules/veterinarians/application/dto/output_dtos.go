package dto

import "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"

// @Description Represents the response structure for a veterinarian.
type VetResponse struct {
	// The unique ID of the veterinarian.
	ID int `json:"id"`
	// The first name of the veterinarian.
	FirstName string `json:"first_name"`
	// The last name of the veterinarian.
	LastName string `json:"last_name"`
	// The URL of the veterinarian's photo.
	Photo string `json:"photo"`
	// The license number of the veterinarian.
	LicenseNumber string `json:"license_number"`
	// The years of experience of the veterinarian.
	YearsExperience int `json:"years_experience"`
	// The specialty of the veterinarian.
	Specialty string `json:"specialty"`
	// The consultation fee charged by the veterinarian.
	ConsultationFee *valueobject.Money `json:"consultation_fee"`
	// The working schedule of the veterinarian.
	LaboralSchedule *[]ScheduleData `json:"laboral_schedule"`
}
