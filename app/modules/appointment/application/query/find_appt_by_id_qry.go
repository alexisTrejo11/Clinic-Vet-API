package query

import "clinic-vet-api/app/modules/core/domain/valueobject"

type FindApptByIDQuery struct {
	appointmentID valueobject.AppointmentID
	customerID    *valueobject.CustomerID
	employeeID    *valueobject.EmployeeID
}

func NewFindApptByIDQuery(id uint, employeeID *uint, customerID *uint) FindApptByIDQuery {
	return FindApptByIDQuery{
		appointmentID: valueobject.NewAppointmentID(id),
		customerID:    valueobject.NewOptCustomerID(customerID),
		employeeID:    valueobject.NewOptEmployeeID(employeeID),
	}
}

func (q FindApptByIDQuery) AppointmentID() valueobject.AppointmentID { return q.appointmentID }
func (q FindApptByIDQuery) CustomerID() *valueobject.CustomerID      { return q.customerID }
func (q FindApptByIDQuery) EmployeeID() *valueobject.EmployeeID      { return q.employeeID }
