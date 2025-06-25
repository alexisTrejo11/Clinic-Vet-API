package vetDtos

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type VetCreate struct {
	FirstName       string            `json:"first_name" validate:"required"`
	LastName        string            `json:"last_name" validate:"required"`
	Photo           string            `json:"photo"`
	LicenseNumber   string            `json:"license_number" validate:"required"`
	YearsExperience uint              `json:"years_experience"`
	IsActive        bool              `json:"is_active"`
	Specialty       *string           `json:"specialty"`
	ConsultationFee *shared.Money     `json:"consultation_fee"`
	LaboralSchedule *[]ScheduleInsert `json:"laboral_schedule"`
}

type ScheduleInsert struct {
	Day           time.Weekday
	EntryTime     int
	DepartureTime int
	StartBreak    int
	EndBreak      int
}

type VetUpdate struct {
	FirstName       *string           `json:"first_name"`
	LastName        *string           `json:"last_name"`
	Photo           *string           `json:"photo"`
	LicenseNumber   *string           `json:"license_number"`
	YearsExperience *uint             `json:"years_experience"`
	Specialty       *string           `json:"specialty"`
	IsActive        *bool             `json:"is_active"`
	ConsultationFee *shared.Money     `json:"consultation_fee"`
	LaboralSchedule *[]ScheduleInsert `json:"laboral_schedule"`
}
