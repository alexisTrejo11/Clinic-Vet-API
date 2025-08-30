package entity

import (
	"errors"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
)

type Appointment struct {
	id            valueobject.AppointmentID
	service       enum.ClinicService
	scheduledDate time.Time
	status        enum.AppointmentStatus
	reason        string
	notes         *string
	ownerID       int
	vetID         *valueobject.VetID
	petID         valueobject.PetID
	createdAt     time.Time
	updatedAt     time.Time
}

func NewAppointment(
	id valueobject.AppointmentID,
	petID valueobject.PetID,
	ownerID int,
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

func (a *Appointment) GetOwnerID() int {
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

func (a *Appointment) SetVetID(vetID *valueobject.VetID) {
	a.vetID = vetID
}

func (a *Appointment) SetService(service enum.ClinicService) {
	a.service = service
}

func (a *Appointment) SetScheduledDate(scheduledDate time.Time) {
	a.scheduledDate = scheduledDate
}

func (a *Appointment) SetStatus(status enum.AppointmentStatus) {
	a.status = status
}

func (a *Appointment) SetUpdatedAt(updatedAt time.Time) {
	a.updatedAt = updatedAt
}

func (a *Appointment) SetReason(reason string) {
	a.reason = reason
}

func (a *Appointment) SetNotes(notes *string) {
	if notes != nil && *notes == "" {
		a.notes = nil
	} else {
		a.notes = notes
	}
}

func (a *Appointment) SetID(id valueobject.AppointmentID) {
	a.id = id
}

type AppointmentBuilder struct {
	id            valueobject.AppointmentID
	petID         valueobject.PetID
	ownerID       int
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
func (ab *AppointmentBuilder) WithOwnerID(ownerID int) *AppointmentBuilder {
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

func (ab *AppointmentBuilder) Build() (*Appointment, error) {
	// Simple validation, you can add more complex logic here.
	if ab.ownerID <= 0 {
		return nil, errors.New("owner ID must be greater than zero")
	}

	appointment := &Appointment{
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

	return appointment, nil
}
