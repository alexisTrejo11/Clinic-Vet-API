// Package entity contains all the domain entities to handle all the buissness logic
package entity

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
)

const (
	MinAllowedDaysToSchedule = 3
	MaxAllowedDaysToSchedule = 30
)

type Appointment struct {
	id            valueobject.AppointmentID
	service       enum.ClinicService
	scheduledDate time.Time
	status        enum.AppointmentStatus
	reason        string
	notes         *string
	ownerID       valueobject.OwnerID
	vetID         *valueobject.VetID
	petID         valueobject.PetID
	createdAt     time.Time
	updatedAt     time.Time
}

func NewAppointment(
	id valueobject.AppointmentID,
	petID valueobject.PetID,
	ownerID valueobject.OwnerID,
	vetID *valueobject.VetID,
	service enum.ClinicService,
	scheduledDate time.Time,
	status enum.AppointmentStatus,
	createdAt,
	updatedAt time.Time,
) *Appointment {
	return &Appointment{
		id:            id,
		petID:         petID,
		ownerID:       ownerID,
		vetID:         vetID,
		service:       service,
		scheduledDate: scheduledDate,
		status:        status,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

func (a *Appointment) GetID() valueobject.AppointmentID {
	return a.id
}

func (a *Appointment) GetPetID() valueobject.PetID {
	return a.petID
}

func (a *Appointment) GetOwnerID() valueobject.OwnerID {
	return a.ownerID
}

func (a *Appointment) GetVetID() *valueobject.VetID {
	return a.vetID
}

func (a *Appointment) GetService() enum.ClinicService {
	return a.service
}

func (a *Appointment) GetScheduledDate() time.Time {
	return a.scheduledDate
}

func (a *Appointment) GetStatus() enum.AppointmentStatus {
	return a.status
}

func (a *Appointment) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *Appointment) GetUpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Appointment) GetReason() string {
	return a.reason
}

func (a *Appointment) GetNotes() *string {
	return a.notes
}

func (a *Appointment) ValidateFields() error {
	if a.ownerID.IsZero() {
		return errors.New("owner ID must be greater than zero")
	}

	if a.GetScheduledDate().IsZero() {
		return domainerr.AppointmentScheduleDateZeroErr()
	}

	if err := a.ValidateRequestSchedule(); err != nil {
		return err
	}
	return nil
}

func (a *Appointment) Update(notes *string, vetID *valueobject.VetID, service *enum.ClinicService, reason *string) error {
	if notes != nil {
		a.notes = notes
	}

	if vetID != nil {
		a.vetID = vetID
	}

	if service != nil {
		a.service = *service
	}

	if reason != nil {
		a.reason = *reason
	}

	return nil
}

func (a *Appointment) RescheduleAppointment(newDate time.Time) error {
	if newDate.Before(time.Now()) {
		return domainerr.AppointmentScheduleDateRuleErr("appointment date must be in the future")
	}

	a.scheduledDate = newDate
	a.status = enum.StatusRescheduled
	a.updatedAt = time.Now()

	return nil
}

func (a *Appointment) Cancel() error {
	if err := a.validateStatusTranstion(enum.StatusCancelled); err != nil {
		return err
	}

	a.status = enum.StatusCancelled
	a.updatedAt = time.Now()
	return nil
}

func (a *Appointment) Complete() error {
	if err := a.validateStatusTranstion(enum.StatusCompleted); err != nil {
		return err
	}

	a.status = enum.StatusCompleted
	a.updatedAt = time.Now()

	return nil
}

func (a *Appointment) NotPresented() error {
	if err := a.validateStatusTranstion(enum.StatusNotPresented); err != nil {
		return err
	}

	a.status = enum.StatusNotPresented
	a.updatedAt = time.Now()
	return nil
}

func (a *Appointment) ValidateRequestSchedule() error {
	now := time.Now()

	if a.GetScheduledDate().IsZero() {
		return domainerr.AppointmentScheduleDateZeroErr()
	}

	if a.GetScheduledDate().Before(now) {
		return domainerr.AppointmentScheduleDateRuleErr("scheduled date cannot be in the past")
	}

	if a.GetScheduledDate().Before(now.AddDate(0, 0, MinAllowedDaysToSchedule)) {
		return domainerr.AppointmentScheduleDateRuleErr("appointments must be scheduled at least 3 days in advance")
	}

	if a.GetScheduledDate().Weekday() == time.Saturday || a.GetScheduledDate().Weekday() == time.Sunday {
		return domainerr.AppointmentScheduleDateRuleErr("appointments cannot be scheduled on weekends")
	}

	return nil
}

func (a *Appointment) Confirm(vetID *valueobject.VetID) error {
	if err := a.validateStatusTranstion(enum.StatusConfirmed); err != nil {
		return err
	}

	a.vetID = vetID
	a.status = enum.StatusConfirmed
	a.updatedAt = time.Now()
	return nil
}

func (a *Appointment) validateStatusTranstion(toStatus enum.AppointmentStatus) error {
	return nil
}

func (a *Appointment) SetID(id valueobject.AppointmentID) {
	a.id = id
}

type AppointmentBuilder struct {
	id            valueobject.AppointmentID
	petID         valueobject.PetID
	ownerID       valueobject.OwnerID
	vetID         *valueobject.VetID
	service       enum.ClinicService
	scheduledDate time.Time
	status        enum.AppointmentStatus
	reason        string
	notes         *string
	createdAt     time.Time
	updatedAt     time.Time
}

// NewAppointmentBuilder creates and returns a new AppointmentBuilder
func NewAppointmentBuilder() *AppointmentBuilder {
	return &AppointmentBuilder{}
}

// WithID sets the appointment ID.
func (ab *AppointmentBuilder) WithID(id valueobject.AppointmentID) *AppointmentBuilder {
	ab.id = id
	return ab
}

// WithPetID sets the pet ID.
func (ab *AppointmentBuilder) WithPetID(petID valueobject.PetID) *AppointmentBuilder {
	ab.petID = petID
	return ab
}

// WithOwnerID sets the owner ID.
func (ab *AppointmentBuilder) WithOwnerID(ownerID valueobject.OwnerID) *AppointmentBuilder {
	ab.ownerID = ownerID
	return ab
}

// WithVetID sets the veterinarian ID.
func (ab *AppointmentBuilder) WithVetID(vetID *valueobject.VetID) *AppointmentBuilder {
	ab.vetID = vetID
	return ab
}

// WithService sets the clinic service.
func (ab *AppointmentBuilder) WithService(service enum.ClinicService) *AppointmentBuilder {
	ab.service = service
	return ab
}

// WithScheduledDate sets the scheduled date and time.
func (ab *AppointmentBuilder) WithScheduledDate(scheduledDate time.Time) *AppointmentBuilder {
	ab.scheduledDate = scheduledDate
	return ab
}

// WithStatus sets the appointment status
func (ab *AppointmentBuilder) WithStatus(status enum.AppointmentStatus) *AppointmentBuilder {
	ab.status = status
	return ab
}

// WithReason sets the reason for the appointment.
func (ab *AppointmentBuilder) WithReason(reason string) *AppointmentBuilder {
	ab.reason = reason
	return ab
}

// WithNotes sets the notes for the appointment.
func (ab *AppointmentBuilder) WithNotes(notes *string) *AppointmentBuilder {
	ab.notes = notes
	return ab
}

// WithTimestamps sets the creation and update timestamps.
func (ab *AppointmentBuilder) WithTimestamps(createdAt, updatedAt time.Time) *AppointmentBuilder {
	ab.createdAt = createdAt
	ab.updatedAt = updatedAt
	return ab
}

func (ab *AppointmentBuilder) Build() *Appointment {
	return &Appointment{
		id:            ab.id,
		petID:         ab.petID,
		ownerID:       ab.ownerID,
		vetID:         ab.vetID,
		service:       ab.service,
		scheduledDate: ab.scheduledDate,
		status:        ab.status,
		reason:        ab.reason,
		notes:         ab.notes,
		createdAt:     ab.createdAt,
		updatedAt:     ab.updatedAt,
	}
}
