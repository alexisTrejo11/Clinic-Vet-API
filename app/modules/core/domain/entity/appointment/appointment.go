// Package appointment defines the Appointment entity and its business logic.
package appointment

import (
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/base"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
)

type Appointment struct {
	base.Entity[valueobject.AppointmentID]
	service       enum.ClinicService
	scheduledDate time.Time
	status        enum.AppointmentStatus
	notes         *string
	customerID    valueobject.CustomerID
	employeeID    *valueobject.EmployeeID
	petID         valueobject.PetID
}

type AppointmentBuilder struct{ appt *Appointment }

func NewAppointmentBuilder() *AppointmentBuilder {
	return &AppointmentBuilder{appt: &Appointment{}}
}

func (b *AppointmentBuilder) Build() *Appointment {
	return b.appt
}

func (b *AppointmentBuilder) WithID(id valueobject.AppointmentID) *AppointmentBuilder {
	b.appt.Entity.SetID(id)
	return b
}

func (b *AppointmentBuilder) WithService(service enum.ClinicService) *AppointmentBuilder {
	b.appt.service = service
	return b
}

func (b *AppointmentBuilder) WithScheduledDate(scheduledDate time.Time) *AppointmentBuilder {
	b.appt.scheduledDate = scheduledDate
	return b
}

func (b *AppointmentBuilder) WithStatus(status enum.AppointmentStatus) *AppointmentBuilder {
	b.appt.status = status
	return b
}

func (b *AppointmentBuilder) WithNotes(notes *string) *AppointmentBuilder {
	b.appt.notes = notes
	return b
}

func (b *AppointmentBuilder) WithEmployeeID(employeeID *valueobject.EmployeeID) *AppointmentBuilder {
	b.appt.employeeID = employeeID
	return b
}

func (b *AppointmentBuilder) WithPetID(petID valueobject.PetID) *AppointmentBuilder {
	b.appt.petID = petID
	return b
}

func (b *AppointmentBuilder) WithCustomerID(customerID valueobject.CustomerID) *AppointmentBuilder {
	b.appt.customerID = customerID
	return b
}

func (b *AppointmentBuilder) WithTimestamps(createdAt, updatedAt time.Time) *AppointmentBuilder {
	b.appt.Entity.SetTimeStamps(createdAt, updatedAt)
	return b
}

func (a *Appointment) ID() valueobject.AppointmentID       { return a.Entity.ID() }
func (a *Appointment) PetID() valueobject.PetID            { return a.petID }
func (a *Appointment) CustomerID() valueobject.CustomerID  { return a.customerID }
func (a *Appointment) EmployeeID() *valueobject.EmployeeID { return a.employeeID }
func (a *Appointment) Service() enum.ClinicService         { return a.service }
func (a *Appointment) ScheduledDate() time.Time            { return a.scheduledDate }
func (a *Appointment) Status() enum.AppointmentStatus      { return a.status }
func (a *Appointment) Notes() *string                      { return a.notes }
