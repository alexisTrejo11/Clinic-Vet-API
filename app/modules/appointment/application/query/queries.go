package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"

	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/app/shared/page"
)

type FindApptByIDQuery struct {
	appointmentID valueobject.AppointmentID
	customerID    *valueobject.CustomerID
	employeeID    *valueobject.EmployeeID
}

func NewFindApptByIDQuery(id uint, employeeID *uint, customerID *uint) *FindApptByIDQuery {
	var empID *valueobject.EmployeeID
	if employeeID != nil {
		val := valueobject.NewEmployeeID(*employeeID)
		empID = &val
	}

	var custID *valueobject.CustomerID
	if customerID != nil {
		val := valueobject.NewCustomerID(*customerID)
		custID = &val
	}

	return &FindApptByIDQuery{
		appointmentID: valueobject.NewAppointmentID(id),
		customerID:    custID,
		employeeID:    empID,
	}
}

type FindApptsByDateRangeQuery struct {
	startDate  time.Time
	endDate    time.Time
	pagination specification.Pagination
}

func NewFindApptsByDateRangeQuery(startDate, endDate time.Time, pagInput page.PaginationRequest) (FindApptsByDateRangeQuery, error) {
	qry := &FindApptsByDateRangeQuery{
		startDate:  startDate,
		endDate:    endDate,
		pagination: pagInput.ToSpecPagination(),
	}

	if startDate.IsZero() {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("startDate", "zero", "startDate can't be zero")
	}

	if endDate.IsZero() {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("endDate", "zero", "endDate can't be zero")
	}

	if startDate.Before(endDate) {
		return FindApptsByDateRangeQuery{}, apperror.FieldValidationError("date-range", "", "startDate can't be before endDate")
	}

	return *qry, nil
}

type FindApptsByCustomerIDQuery struct {
	customerID valueobject.CustomerID
	petID      *valueobject.PetID
	pagination specification.Pagination
}

func NewFindApptsByCustomerIDQuery(
	pagination page.PaginationRequest,
	customerId uint,
	petID *uint, status *string,
) *FindApptsByCustomerIDQuery {
	var petIDvo *valueobject.PetID
	if petID != nil {
		val := valueobject.NewPetID(*petID)
		petIDvo = &val
	}

	return &FindApptsByCustomerIDQuery{
		customerID: valueobject.NewCustomerID(customerId),
		pagination: pagination.ToSpecPagination(),
		petID:      petIDvo,
	}
}

type FindApptsByPetQuery struct {
	petID      valueobject.PetID
	pagination specification.Pagination
}

func NewFindApptsByPetQuery(employeeID uint, pagination page.PaginationRequest) *FindApptsByPetQuery {
	return &FindApptsByPetQuery{petID: valueobject.NewPetID(employeeID), pagination: pagination.ToSpecPagination()}
}

type FindApptsByEmployeeIDQuery struct {
	employeeID valueobject.EmployeeID
	pagination specification.Pagination
}

func NewFindApptsByEmployeeIDQuery(employeeID uint, pagination page.PaginationRequest) *FindApptsByEmployeeIDQuery {
	return &FindApptsByEmployeeIDQuery{
		employeeID: valueobject.NewEmployeeID(employeeID),
		pagination: pagination.ToSpecPagination(),
	}
}

type FindApptsBySpecQuery struct {
	spec specification.ApptSearchSpecification
}

func NewFindApptsBySpecQuery(spec specification.ApptSearchSpecification) *FindApptsBySpecQuery {
	return &FindApptsBySpecQuery{spec: spec}
}
