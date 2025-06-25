package vetDomain

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type Veterinarian struct {
	ID               uint
	Name             shared.PersonName
	Photo            *string
	LicenseNumber    string
	Specialty        *string
	UserID           *uint
	YearsExperience  *int
	ConsultationFee  *shared.Money
	WorkDaysSchedule *[]WorkDaySchedule
	IsActive         bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type WorkDaySchedule struct {
	Day              time.Weekday
	WorkDayHourRange string
	BreakHourRange   string
}
