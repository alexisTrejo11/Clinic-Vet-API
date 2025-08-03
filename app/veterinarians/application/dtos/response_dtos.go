package vetDtos

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type VetResponse struct {
	Id              int               `json:"id"`
	FirstName       string            `json:"first_name"`
	LastName        string            `json:"last_name"`
	Photo           string            `json:"photo"`
	LicenseNumber   string            `json:"license_number"`
	YearsExperience int               `json:"years_experience"`
	Specialty       string            `json:"specialty"`
	ConsultationFee *shared.Money     `json:"consultation_fee"`
	LaboralSchedule *[]ScheduleInsert `json:"laboral_schedule"`
}
