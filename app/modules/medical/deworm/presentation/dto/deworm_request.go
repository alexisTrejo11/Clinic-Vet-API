package dto

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/medical/deworm/application/command"
	"clinic-vet-api/app/modules/medical/deworm/application/query"
	"clinic-vet-api/app/shared/page"
	"time"
)

// CreateDewormRequest represents the payload for creating a new deworming record
// @Description Request payload for creating a new deworming treatment record with comprehensive validation rules
type CreateDewormRequest struct {
	PetID            uint       `json:"petId" binding:"required,min=1" example:"5" description:"ID of the pet receiving the deworming treatment. Must be a positive integer representing an existing pet."`
	MedicationName   string     `json:"medicationName" binding:"required,min=2,max=255" example:"Drontal Plus" description:"Name of the deworming medication. Must be between 2 and 255 characters. Common medications: Drontal, Revolution, Stronghold, Advantage Multi"`
	AdministeredDate time.Time  `json:"administeredDate" binding:"required,datetime=2006-01-02" example:"2024-01-15" description:"Date when the deworming treatment was administered. Format: YYYY-MM-DD. Cannot be a future date."`
	NextDueDate      *time.Time `json:"nextDueDate,omitempty" binding:"omitempty,datetime=2006-01-02" example:"2024-04-15" description:"Optional next due date for deworming. If not provided, system will calculate based on pet's age and medication type. Format: YYYY-MM-DD. Must be after administered date."`
	AdministeredBy   uint       `json:"administeredBy" binding:"required,min=1" example:"3" description:"ID of the employee/veterinarian who administered the treatment. Must be a positive integer representing an existing employee."`
	Notes            *string    `json:"notes,omitempty" binding:"omitempty,max=1000" example:"Administered with food. No adverse reactions. Pet weight: 8.5kg" description:"Optional observations about the treatment. Maximum 1000 characters."`
}

// UpdateDewormRequest represents the payload for updating an existing deworming record
// @Description Request payload for updating an existing deworming treatment record. All fields are optional except ID, but at least one field must be provided.
type UpdateDewormRequest struct {
	MedicationName   *string    `json:"medicationName,omitempty" binding:"omitempty,min=2,max=255" example:"Revolution" description:"Updated medication name. Must be between 2 and 255 characters if provided."`
	AdministeredDate *time.Time `json:"administeredDate,omitempty" binding:"omitempty,datetime=2006-01-02" example:"2024-01-16" description:"Updated administration date. Format: YYYY-MM-DD. Cannot be a future date."`
	NextDueDate      *time.Time `json:"nextDueDate,omitempty" binding:"omitempty,datetime=2006-01-02" example:"2024-04-16" description:"Updated next due date. Format: YYYY-MM-DD. Must be after administered date if both are provided."`
	AdministeredBy   *uint      `json:"administeredBy,omitempty" binding:"omitempty,min=1" example:"4" description:"Updated employee ID who administered the treatment. Must be a positive integer if provided."`
	Notes            *string    `json:"notes,omitempty" binding:"omitempty,max=1000" example:"Updated notes: Pet tolerated treatment well." description:"Updated observations. Maximum 1000 characters."`
}

// BulkCreateDewormRequest represents the payload for creating multiple deworming records
// @Description Request payload for creating multiple deworming records in a single operation. Useful for batch processing.
type BulkCreateDewormRequest struct {
	Dewormings []CreateDewormRequest `json:"dewormings" binding:"required,min=1,max=100,dive" example:"[{'petId': 5, 'medicationName': 'Drontal Plus', 'administeredDate': '2024-01-15', 'administeredBy': 3}]" description:"Array of deworming records to create. Maximum 100 records per request."`
}

// DewormQueryParams represents query parameters for filtering deworming records
// @Description Query parameters for filtering and paginating deworming records
type DewormQueryParams struct {
	PetID                *uint   `form:"petId" binding:"omitempty,min=1" example:"5" description:"Filter by pet ID. Must be positive integer."`
	AdministeredBy       *uint   `form:"administeredBy" binding:"omitempty,min=1" example:"3" description:"Filter by employee ID who administered the treatment."`
	MedicationName       *string `form:"medicationName" binding:"omitempty,min=1,max=255" example:"Drontal" description:"Filter by medication name (partial match, case insensitive)."`
	AdministeredDateFrom *string `form:"administeredDateFrom" binding:"omitempty,datetime=2006-01-02" example:"2024-01-01" description:"Filter by administered date range start. Format: YYYY-MM-DD."`
	AdministeredDateTo   *string `form:"administeredDateTo" binding:"omitempty,datetime=2006-01-02" example:"2024-01-31" description:"Filter by administered date range end. Format: YYYY-MM-DD."`
	NextDueDateFrom      *string `form:"nextDueDateFrom" binding:"omitempty,datetime=2006-01-02" example:"2024-04-01" description:"Filter by next due date range start. Format: YYYY-MM-DD."`
	NextDueDateTo        *string `form:"nextDueDateTo" binding:"omitempty,datetime=2006-01-02" example:"2024-04-30" description:"Filter by next due date range end. Format: YYYY-MM-DD."`
	Pagination           page.PaginationRequest
}

// FindDewormsByDateRangeRequest represents the payload for querying deworming records within a date range
// @Description Request payload for querying deworming records within a specific date range with pagination support
type FindDewormsByDateRangeRequest struct {
	StartDate  time.Time `form:"startDate" binding:"required,datetime=2006-01-02" example:"2024-01-01" description:"Start date of the range to filter deworming records (format: YYYY-MM-DD)"`
	EndDate    time.Time `form:"endDate" binding:"required,datetime=2006-01-02" example:"2024-01-31" description:"End date of the range to filter deworming records (format: YYYY-MM-DD)"`
	Pagination page.PaginationRequest
}

func (r FindDewormsByDateRangeRequest) ToQuery() *query.FindDewormsByDateRangeQuery {
	query := &query.FindDewormsByDateRangeQuery{
		StartDate:  r.StartDate,
		EndDate:    r.EndDate,
		Pagination: r.Pagination,
	}
	return query
}

func (r *CreateDewormRequest) ToCommand() command.DewormCreateCommand {
	return command.NewDewormCreateCommand(
		r.PetID,
		r.AdministeredBy,
		r.MedicationName,
		r.AdministeredDate,
		r.NextDueDate,
		r.Notes,
	)

}

func (r *UpdateDewormRequest) ToCommand(dewormID uint) command.DewormUpdateCommand {
	return command.DewormUpdateCommand{
		ID:               valueobject.NewDewormID(dewormID),
		MedicationName:   r.MedicationName,
		AdministeredDate: r.AdministeredDate,
		NextDueDate:      r.NextDueDate,
		AdministeredBy:   valueobject.NewOptEmployeeID(r.AdministeredBy),
		Notes:            r.Notes,
	}
}
