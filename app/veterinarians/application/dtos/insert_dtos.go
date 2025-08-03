package vetDtos

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type VetCreate struct {
	FirstName       string                 `json:"first_name" validate:"required"`
	LastName        string                 `json:"last_name" validate:"required"`
	Photo           string                 `json:"photo"`
	LicenseNumber   string                 `json:"license_number" validate:"required"`
	YearsExperience int                    `json:"years_experience"`
	IsActive        bool                   `json:"is_active"`
	Specialty       vetDomain.VetSpecialty `json:"specialty"`
	ConsultationFee *shared.Money          `json:"consultation_fee"`
	LaboralSchedule []ScheduleInsert       `json:"laboral_schedule"`
}

type ScheduleInsert struct {
	Day           time.Weekday
	EntryTime     int
	DepartureTime int
	StartBreak    int
	EndBreak      int
}

type VetUpdate struct {
	FirstName       *string                 `json:"first_name"`
	LastName        *string                 `json:"last_name"`
	Photo           *string                 `json:"photo"`
	LicenseNumber   *string                 `json:"license_number"`
	YearsExperience *int                    `json:"years_experience"`
	Specialty       *vetDomain.VetSpecialty `json:"specialty"`
	IsActive        *bool                   `json:"is_active"`
	ConsultationFee *shared.Money           `json:"consultation_fee"`
	LaboralSchedule *[]ScheduleInsert       `json:"laboral_schedule"`
}

type VetSearchParams struct {
	shared.PageInput
	Filters VetFilters `json:"filters"`
	OrderBy VetOrderBy
}

type VetFilters struct {
	Name            *string                 `json:"name"`
	LicenseNumber   *string                 `json:"license_number"`
	Specialty       *vetDomain.VetSpecialty `json:"specialty"`
	YearsExperience *struct {
		Min *int `json:"min"`
		Max *int `json:"max"`
	} `json:"years_experience"`
	IsActive *bool `json:"is_active"`
}

type VetOrderBy string

const (
	OrderByName            VetOrderBy = "name"
	OrderBySpecialty       VetOrderBy = "specialty"
	OrderByYearsExperience VetOrderBy = "years_experience"
	OrderByCreatedAt       VetOrderBy = "created_at"
)
