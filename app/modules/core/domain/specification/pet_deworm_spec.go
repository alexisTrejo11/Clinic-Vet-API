package specification

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

// SortOrder represents the sorting order
type SortOrder string

const (
	SortByAdministeredDateDesc SortOrder = "administered_date_desc"
	SortByAdministeredDateAsc  SortOrder = "administered_date_asc"
	SortByNextDueDateDesc      SortOrder = "next_due_date_desc"
	SortByNextDueDateAsc       SortOrder = "next_due_date_asc"
	SortByCreatedAtDesc        SortOrder = "created_at_desc"
	SortByCreatedAtAsc         SortOrder = "created_at_asc"
)

// PetDewormSpecification encapsulates query criteria for pet dewormings
type PetDewormSpecification struct {
	ID                    *valueobject.DewormID
	PetID                 *valueobject.PetID
	AdministeredBy        *valueobject.EmployeeID
	MedicationName        *string
	AdministeredDateFrom  *time.Time
	AdministeredDateTo    *time.Time
	AdministeredDateExact *time.Time
	NextDueDateFrom       *time.Time
	NextDueDateTo         *time.Time
	NextDueDateExact      *time.Time
	CreatedAtFrom         *time.Time
	CreatedAtTo           *time.Time
	SortBy                SortOrder
	Limit                 *int32
	Offset                *int32
}

// NewPetDewormSpecification creates a new specification instance
func NewPetDewormSpecification() *PetDewormSpecification {
	return &PetDewormSpecification{}
}

func NewPetDewormSpecificationWithDefaults() *PetDewormSpecification {
	return &PetDewormSpecification{
		SortBy: SortByAdministeredDateDesc,
		Limit:  intPtr(50),
		Offset: intPtr(0),
	}
}

func intPtr(i int32) *int32 {
	return &i
}

// WithID sets the deworming ID filter
func (s *PetDewormSpecification) WithID(id valueobject.DewormID) *PetDewormSpecification {
	s.ID = &id
	return s
}

// WithPetID sets the pet ID filter
func (s *PetDewormSpecification) WithPetID(petID valueobject.PetID) *PetDewormSpecification {
	s.PetID = &petID
	return s
}

// WithAdministeredBy sets the administered by employee filter
func (s *PetDewormSpecification) WithAdministeredBy(employeeID valueobject.EmployeeID) *PetDewormSpecification {
	s.AdministeredBy = &employeeID
	return s
}

// WithMedicationName sets the medication name filter (partial match)
func (s *PetDewormSpecification) WithMedicationName(medicationName string) *PetDewormSpecification {
	s.MedicationName = &medicationName
	return s
}

// WithAdministeredDateRange sets the administered date range filter
func (s *PetDewormSpecification) WithAdministeredDateRange(from, to time.Time) *PetDewormSpecification {
	s.AdministeredDateFrom = &from
	s.AdministeredDateTo = &to
	return s
}

// WithAdministeredDateFrom sets the administered date from filter
func (s *PetDewormSpecification) WithAdministeredDateFrom(from time.Time) *PetDewormSpecification {
	s.AdministeredDateFrom = &from
	return s
}

// WithAdministeredDateTo sets the administered date to filter
func (s *PetDewormSpecification) WithAdministeredDateTo(to time.Time) *PetDewormSpecification {
	s.AdministeredDateTo = &to
	return s
}

// WithAdministeredDateExact sets the exact administered date filter
func (s *PetDewormSpecification) WithAdministeredDateExact(date time.Time) *PetDewormSpecification {
	s.AdministeredDateExact = &date
	return s
}

// WithNextDueDateRange sets the next due date range filter
func (s *PetDewormSpecification) WithNextDueDateRange(from, to time.Time) *PetDewormSpecification {
	s.NextDueDateFrom = &from
	s.NextDueDateTo = &to
	return s
}

// WithNextDueDateFrom sets the next due date from filter
func (s *PetDewormSpecification) WithNextDueDateFrom(from time.Time) *PetDewormSpecification {
	s.NextDueDateFrom = &from
	return s
}

// WithNextDueDateTo sets the next due date to filter
func (s *PetDewormSpecification) WithNextDueDateTo(to time.Time) *PetDewormSpecification {
	s.NextDueDateTo = &to
	return s
}

// WithNextDueDateExact sets the exact next due date filter
func (s *PetDewormSpecification) WithNextDueDateExact(date time.Time) *PetDewormSpecification {
	s.NextDueDateExact = &date
	return s
}

// WithCreatedAtRange sets the created at date range filter
func (s *PetDewormSpecification) WithCreatedAtRange(from, to time.Time) *PetDewormSpecification {
	s.CreatedAtFrom = &from
	s.CreatedAtTo = &to
	return s
}

// WithCreatedAtFrom sets the created at from filter
func (s *PetDewormSpecification) WithCreatedAtFrom(from time.Time) *PetDewormSpecification {
	s.CreatedAtFrom = &from
	return s
}

// WithCreatedAtTo sets the created at to filter
func (s *PetDewormSpecification) WithCreatedAtTo(to time.Time) *PetDewormSpecification {
	s.CreatedAtTo = &to
	return s
}

// WithSort sets the sort order
func (s *PetDewormSpecification) WithSort(sortBy SortOrder) *PetDewormSpecification {
	s.SortBy = sortBy
	return s
}

// WithLimit sets the limit for pagination
func (s *PetDewormSpecification) WithLimit(limit int32) *PetDewormSpecification {
	s.Limit = &limit
	return s
}

// WithOffset sets the offset for pagination
func (s *PetDewormSpecification) WithOffset(offset int32) *PetDewormSpecification {
	s.Offset = &offset
	return s
}

// WithPagination sets both limit and offset for pagination
func (s *PetDewormSpecification) WithPagination(limit, offset int32) *PetDewormSpecification {
	s.Limit = &limit
	s.Offset = &offset
	return s
}
