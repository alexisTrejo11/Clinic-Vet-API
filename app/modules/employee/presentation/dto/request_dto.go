// Package dto contains data transfer objects for the employee module.
package dto

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/employee/application/cqrs/command"
	"clinic-vet-api/app/shared/page"
	"time"
)

// TODO: CREATE PERSONAL DATA DTO
// CreateEmployeeRequest represents the request to create a new employee
// @Description Request body for creating a new employee
type CreateEmployeeRequest struct {
	// Customer's first name
	// Required: true
	// Minimum length: 2, Maximum length: 50
	FirstName string `json:"first_name" binding:"required,min=2,max=50" example:"John"`

	// Customer's last name
	// Required: true
	// Minimum length: 2, Maximum length: 50
	LastName string `json:"last_name" binding:"required,min=2,max=50" example:"Doe"`

	// Customer's gender
	// Required: true
	// Enum: male, female, not_specified
	Gender string `json:"gender" binding:"required" example:"male"`

	// Customer's date of birth
	// Required: true
	DateOfBirth time.Time `json:"date_of_birth" binding:"required" example:"1990-01-15T00:00:00Z"`

	// URL or path to employee's photo
	// Required: false
	Photo string `json:"photo" example:"https://example.com/photo.jpg"`

	// Professional license number
	// Required: true
	LicenseNumber string `json:"license_number" validate:"required" example:"VET123456"`

	// Years of professional experience
	// Required: false
	YearsExperience int `json:"years_experience" example:"5"`

	// Indicates if the employee is currently active
	// Required: false
	IsActive bool `json:"is_active" example:"true"`

	// Veterinary specialty
	// Required: false
	Specialty string `json:"specialty" example:"CARDIOLOGY"`

	// Consultation fee information
	// Required: false
	ConsultationFee *valueobject.Money `json:"consultation_fee"`

	// Weekly work schedule
	// Required: false
	LaboralSchedule []ScheduleRequest `json:"laboral_schedule"`
}

// ScheduleRequest represents an employee's work schedule for a specific day
// @Description Work schedule details for a specific day
type ScheduleRequest struct {
	// Day of the week (e.g., Monday, Tuesday, etc.)
	// Required: true
	Day string `json:"day" example:"Monday"`

	// Entry time in minutes from midnight
	// Required: true
	EntryTime int `json:"entry_time" example:"540"` // 9:00 AM

	// Departure time in minutes from midnight
	// Required: true
	DepartureTime int `json:"departure_time" example:"1020"` // 5:00 PM

	// Break start time in minutes from midnight
	// Required: false
	StartBreak int `json:"start_break" example:"720"` // 12:00 PM

	// Break end time in minutes from midnight
	// Required: false
	EndBreak int `json:"end_break" example:"780"` // 1:00 PM
}

// UpdateEmployeeRequest represents the request to update an existing employee
// @Description Request body for updating an employee
type UpdateEmployeeRequest struct {
	// First name of the employee
	// Required: false
	FirstName *string `json:"first_name" example:"Jane"`

	// Last name of the employee
	// Required: false
	LastName *string `json:"last_name" example:"Smith"`

	// URL or path to employee's photo
	// Required: false
	Photo *string `json:"photo" example:"https://example.com/new-photo.jpg"`

	// Professional license number
	// Required: false
	LicenseNumber *string `json:"license_number" example:"VET654321"`

	// Veterinary specialty
	// Required: false
	Specialty *enum.VetSpecialty `json:"specialty" example:"DERMATOLOGY"`

	// Indicates if the employee is currently active
	// Required: false
	IsActive *bool `json:"is_active" example:"true"`

	// Consultation fee information
	// Required: false
	ConsultationFee map[string]string `json:"consultation_fee"`

	// Years of professional experience
	// Required: false
	YearsExperience *int `json:"years_experience" example:"7"`

	// Weekly work schedule
	// Required: false
	LaboralSchedule *[]ScheduleRequest `json:"laboral_schedule"`
}

// EmployeeSearchParams represents the parameters for searching employees
// @Description Search parameters for employee queries
type EmployeeSearchParams struct {
	page.PaginationRequest
	Filters EmployeeFilters `json:"filters"`
	OrderBy EmployeeOrderBy
}

// EmployeeFilters represents the filter criteria for employee searches
// @Description Filter criteria for employee search
type EmployeeFilters struct {
	// Filter by employee name (partial match)
	// Required: false
	Name *string `json:"name" example:"John"`

	// Filter by license number (exact match)
	// Required: false
	LicenseNumber *string `json:"license_number" example:"VET123456"`

	// Filter by veterinary specialty
	// Required: false
	Specialty string `json:"specialty" example:"SURGERY"`

	// Filter by years of experience range
	// Required: false
	YearsExperience *struct {
		// Minimum years of experience
		// Required: false
		Min *int `json:"min" example:"2"`

		// Maximum years of experience
		// Required: false
		Max *int `json:"max" example:"10"`
	} `json:"years_experience"`

	// Filter by active status
	// Required: false
	IsActive *bool `json:"is_active" example:"true"`
}

// EmployeeOrderBy represents the available ordering options for employee results
type EmployeeOrderBy string

const (
	// OrderByName orders results by employee name
	OrderByName EmployeeOrderBy = "name"

	// OrderBySpecialty orders results by veterinary specialty
	OrderBySpecialty EmployeeOrderBy = "specialty"

	// OrderByYearsExperience orders results by years of experience
	OrderByYearsExperience EmployeeOrderBy = "years_experience"

	// OrderByCreatedAt orders results by creation date
	OrderByCreatedAt EmployeeOrderBy = "created_at"
)

func (r *CreateEmployeeRequest) ToCommand() (command.CreateEmployeeCommand, error) {
	cmd := &command.CreateEmployeeCommand{
		Photo:           r.Photo,
		LicenseNumber:   r.LicenseNumber,
		YearsExperience: r.YearsExperience,
		IsActive:        r.IsActive,
		ConsultationFee: r.ConsultationFee,
		DateOfBirth:     r.DateOfBirth,
	}

	personName, err := valueobject.NewPersonName(r.FirstName, r.LastName)
	if err != nil {
		return command.CreateEmployeeCommand{}, err
	}
	cmd.Name = personName

	gender, err := enum.ParseGender(r.Gender)
	if err != nil {
		return command.CreateEmployeeCommand{}, err
	}
	cmd.Gender = gender

	if r.Specialty != "" {
		specialty, err := enum.ParseVetSpecialty(string(r.Specialty))
		if err != nil {
			return command.CreateEmployeeCommand{}, err
		}
		cmd.Specialty = specialty
	}

	if len(r.LaboralSchedule) > 0 {
		var scheduleData []command.ScheduleData
		for _, s := range r.LaboralSchedule {
			scheduleData = append(scheduleData, command.ScheduleData{
				Day:           s.Day,
				EntryTime:     s.EntryTime,
				DepartureTime: s.DepartureTime,
				StartBreak:    s.StartBreak,
				EndBreak:      s.EndBreak,
			})
		}

		cmd.LaboralSchedule = scheduleData
	}

	return *cmd, nil
}

func (r *UpdateEmployeeRequest) ToCommand(employeeIDUint uint) *command.UpdateEmployeeCommand {
	return &command.UpdateEmployeeCommand{
		EmployeeID:      valueobject.NewEmployeeID(employeeIDUint),
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
		IsActive: r.IsActive,
		//ConsultationFee: r.ConsultationFee,
	}
}
