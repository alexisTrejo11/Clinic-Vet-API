// Package appointment defines the Appointment entity and its business logic.
package appointment

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/base"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type Appointment struct {
	base.Entity[valueobject.AppointmentID]
	service       enum.ClinicService
	scheduledDate time.Time
	status        enum.AppointmentStatus
	reason        enum.VisitReason
	notes         *string
	customerID    valueobject.CustomerID
	employeeID    *valueobject.EmployeeID
	petID         valueobject.PetID
}

func (a *Appointment) ID() valueobject.AppointmentID {
	return a.Entity.ID()
}

func (a *Appointment) PetID() valueobject.PetID {
	return a.petID
}

func (a *Appointment) CustomerID() valueobject.CustomerID {
	return a.customerID
}

func (a *Appointment) EmployeeID() *valueobject.EmployeeID {
	return a.employeeID
}

func (a *Appointment) Service() enum.ClinicService {
	return a.service
}

func (a *Appointment) ScheduledDate() time.Time {
	return a.scheduledDate
}

func (a *Appointment) Status() enum.AppointmentStatus {
	return a.status
}

func (a *Appointment) Reason() enum.VisitReason {
	return a.reason
}

func (a *Appointment) Notes() *string {
	return a.notes
}
