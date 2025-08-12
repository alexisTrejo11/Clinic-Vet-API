package appointDomain

import (
	"errors"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type AppointmentBuilder struct {
	id            AppointmentId
	petId         petDomain.PetId
	ownerId       int
	vetId         *vetDomain.VetId
	service       ClinicService
	scheduledDate time.Time
	status        AppointmentStatus
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
func (ab *AppointmentBuilder) WithID(id AppointmentId) *AppointmentBuilder {
	ab.id = id
	return ab
}

// WithPetID sets the pet ID.
func (ab *AppointmentBuilder) WithPetID(petId petDomain.PetId) *AppointmentBuilder {
	ab.petId = petId
	return ab
}

// WithOwnerID sets the owner ID.
func (ab *AppointmentBuilder) WithOwnerID(ownerId int) *AppointmentBuilder {
	ab.ownerId = ownerId
	return ab
}

// WithVetID sets the veterinarian ID.
func (ab *AppointmentBuilder) WithVetID(vetId *vetDomain.VetId) *AppointmentBuilder {
	ab.vetId = vetId
	return ab
}

// WithService sets the clinic service.
func (ab *AppointmentBuilder) WithService(service ClinicService) *AppointmentBuilder {
	ab.service = service
	return ab
}

// WithScheduledDate sets the scheduled date and time.
func (ab *AppointmentBuilder) WithScheduledDate(scheduledDate time.Time) *AppointmentBuilder {
	ab.scheduledDate = scheduledDate
	return ab
}

// WithStatus sets the appointment status.
func (ab *AppointmentBuilder) WithStatus(status AppointmentStatus) *AppointmentBuilder {
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
	if ab.petId.Equals(petDomain.PetId{}) {
		return nil, errors.New("pet ID cannot be nil")
	}

	if ab.ownerId <= 0 {
		return nil, errors.New("owner ID must be greater than zero")
	}

	appointment := &Appointment{
		id:            ab.id,
		petId:         ab.petId,
		ownerId:       ab.ownerId,
		vetId:         ab.vetId,
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
