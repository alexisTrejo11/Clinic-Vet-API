package specification

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	vo "clinic-vet-api/app/modules/core/domain/valueobject"
	"time"
)

// Base specification interface
type ApptSearchSpecification interface {
	ToSQLCParams() ApptSearchParams
	And(spec ApptSearchSpecification) ApptSearchSpecification
	Or(spec ApptSearchSpecification) ApptSearchSpecification
	WithPagination(pagination Pagination) ApptSearchSpecification
}

// Individual specifications
type ApptByIDSpec struct {
	ApptID vo.AppointmentID
}

type ApptByCustomerSpec struct {
	CustomerID vo.CustomerID
}

type ApptByEmployeeSpec struct {
	EmployeeID vo.EmployeeID
}

type ApptByPetSpec struct {
	PetID vo.PetID
}

type ApptByServiceSpec struct {
	Service enum.ClinicService
}

type ApptByStatusSpec struct {
	Status enum.AppointmentStatus
}

type ApptByClinicSpec struct {
	ClinicService enum.ClinicService
}

type ApptByDateRangeSpec struct {
	StartDate time.Time
	EndDate   time.Time
}

type ApptByScheduledDateSpec struct {
	ScheduledDate time.Time
}

// Composite
type ApptCompositeSpec struct {
	specs      []ApptSearchSpecification
	operator   string // "AND" or "OR"
	pagination Pagination
}

type ApptSearchParams struct {
	ApptID        *int32
	CustomerID    *int32
	EmployeeID    *int32
	PetID         *int32
	Service       *string
	Status        *string
	StartDate     *time.Time
	EndDate       *time.Time
	ScheduledDate *time.Time
	Limit         int32
	Offset        int32
}

func (s ApptByIDSpec) ToSQLCParams() ApptSearchParams {
	id := int32(s.ApptID.Value())
	return ApptSearchParams{ApptID: &id}
}

func (s ApptByCustomerSpec) ToSQLCParams() ApptSearchParams {
	id := int32(s.CustomerID.Value()) // Convert vo.CustomerID to int32
	return ApptSearchParams{CustomerID: &id}
}

func (s ApptByEmployeeSpec) ToSQLCParams() ApptSearchParams {
	id := int32(s.EmployeeID.Value()) // Convert vo.EmployeeID to int32
	return ApptSearchParams{EmployeeID: &id}
}

func (s ApptByPetSpec) ToSQLCParams() ApptSearchParams {
	id := int32(s.PetID.Value()) // Convert vo.PetID to int32
	return ApptSearchParams{PetID: &id}
}

func (s ApptByServiceSpec) ToSQLCParams() ApptSearchParams {
	serviceStr := string(s.Service)
	return ApptSearchParams{Service: &serviceStr}
}

func (s ApptByStatusSpec) ToSQLCParams() ApptSearchParams {
	statusStr := string(s.Status)
	return ApptSearchParams{Status: &statusStr}
}

func (s ApptByClinicSpec) ToSQLCParams() ApptSearchParams {
	// Remove this since you don't have ClinicService field
	return ApptSearchParams{}
}

func (s ApptByDateRangeSpec) ToSQLCParams() ApptSearchParams {
	return ApptSearchParams{
		StartDate: &s.StartDate,
		EndDate:   &s.EndDate,
	}
}

func (s ApptByScheduledDateSpec) ToSQLCParams() ApptSearchParams {
	return ApptSearchParams{ScheduledDate: &s.ScheduledDate}
}

func (s ApptCompositeSpec) ToSQLCParams() ApptSearchParams {
	params := ApptSearchParams{
		Limit:  int32(s.pagination.Limit),
		Offset: int32(s.pagination.Offset),
	}

	// Merge all specifications
	for _, spec := range s.specs {
		specParams := spec.ToSQLCParams()

		if specParams.ApptID != nil && params.ApptID == nil {
			params.ApptID = specParams.ApptID
		}
		if specParams.CustomerID != nil && params.CustomerID == nil {
			params.CustomerID = specParams.CustomerID
		}
		if specParams.EmployeeID != nil && params.EmployeeID == nil {
			params.EmployeeID = specParams.EmployeeID
		}
		if specParams.PetID != nil && params.PetID == nil {
			params.PetID = specParams.PetID
		}
		if specParams.Service != nil && params.Service == nil {
			params.Service = specParams.Service
		}
		if specParams.Status != nil && params.Status == nil {
			params.Status = specParams.Status
		}
		if specParams.StartDate != nil && params.StartDate == nil {
			params.StartDate = specParams.StartDate
		}
		if specParams.EndDate != nil && params.EndDate == nil {
			params.EndDate = specParams.EndDate
		}
		if specParams.ScheduledDate != nil && params.ScheduledDate == nil {
			params.ScheduledDate = specParams.ScheduledDate
		}
	}

	return params
}

// ==========================================
// 3. SPECIFICATION BUILDER METHODS
// ==========================================

func (s ApptByIDSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByIDSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByIDSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByCustomerSpec methods
func (s ApptByCustomerSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByCustomerSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByCustomerSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByEmployeeSpec methods
func (s ApptByEmployeeSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByEmployeeSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByEmployeeSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByPetSpec methods
func (s ApptByPetSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByPetSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByPetSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByServiceSpec methods
func (s ApptByServiceSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByServiceSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByServiceSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByStatusSpec methods
func (s ApptByStatusSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByStatusSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByStatusSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByClinicSpec methods
func (s ApptByClinicSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByClinicSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByClinicSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByDateRangeSpec methods
func (s ApptByDateRangeSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByDateRangeSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByDateRangeSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

// ApptByScheduledDateSpec methods
func (s ApptByScheduledDateSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptByScheduledDateSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptByScheduledDateSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s}, operator: "AND", pagination: pagination}
}

func (s ApptCompositeSpec) And(spec ApptSearchSpecification) ApptSearchSpecification {
	if s.operator == "AND" {
		return &ApptCompositeSpec{
			specs:      append(s.specs, spec),
			operator:   "AND",
			pagination: s.pagination,
		}
	}
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "AND"}
}

func (s ApptCompositeSpec) Or(spec ApptSearchSpecification) ApptSearchSpecification {
	if s.operator == "OR" {
		return &ApptCompositeSpec{
			specs:      append(s.specs, spec),
			operator:   "OR",
			pagination: s.pagination,
		}
	}
	return &ApptCompositeSpec{specs: []ApptSearchSpecification{s, spec}, operator: "OR"}
}

func (s ApptCompositeSpec) WithPagination(pagination Pagination) ApptSearchSpecification {
	return &ApptCompositeSpec{
		specs:      s.specs,
		operator:   s.operator,
		pagination: pagination,
	}
}

func NewApptSearch() ApptSearchSpecification {
	return &ApptCompositeSpec{
		specs:      []ApptSearchSpecification{},
		operator:   "AND",
		pagination: Pagination{Limit: 50, Offset: 0},
	}
}

func ApptByID(id vo.AppointmentID) ApptSearchSpecification {
	return ApptByIDSpec{ApptID: id}
}

func ApptByCustomer(customerID vo.CustomerID) ApptSearchSpecification {
	return ApptByCustomerSpec{CustomerID: customerID}
}

func ApptByEmployee(employeeID vo.EmployeeID) ApptSearchSpecification {
	return ApptByEmployeeSpec{EmployeeID: employeeID}
}

func ApptByPet(petID vo.PetID) ApptSearchSpecification {
	return ApptByPetSpec{PetID: petID}
}

func ApptByService(service enum.ClinicService) ApptSearchSpecification {
	return ApptByServiceSpec{Service: service}
}

func ApptByStatus(status enum.AppointmentStatus) ApptSearchSpecification {
	return ApptByStatusSpec{Status: status}
}

func ApptByClinicService(ClinicService enum.ClinicService) ApptSearchSpecification {
	return ApptByClinicSpec{ClinicService: ClinicService}
}

func ApptByDateRange(start, end time.Time) ApptSearchSpecification {
	return ApptByDateRangeSpec{StartDate: start, EndDate: end}
}

func ApptByScheduledDate(date time.Time) ApptSearchSpecification {
	return ApptByScheduledDateSpec{ScheduledDate: date}
}
