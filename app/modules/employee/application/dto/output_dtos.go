package dto

import "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"

// @Description Represents the response structure for a Employee.
type EmployeeResponse struct {
	// The unique ID of the employee.
	ID uint `json:"id"`
	// The first name of the employee.
	FirstName string `json:"first_name"`
	// The last name of the employee.
	LastName string `json:"last_name"`
	// The URL of the employee's photo.
	Photo string `json:"photo"`
	// The license number of the employee.
	LicenseNumber string `json:"license_number"`
	// The years of experience of the employee.
	YearsExperience int `json:"years_experience"`
	// The specialty of the employee.
	Specialty string `json:"specialty"`
	// The consultation fee charged by the employee.
	ConsultationFee *valueobject.Money `json:"consultation_fee"`
	// The working schedule of the employee.
	LaboralSchedule *[]ScheduleData `json:"laboral_schedule"`
}
