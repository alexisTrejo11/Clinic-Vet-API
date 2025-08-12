package appointDomain

import (
	"errors"
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

const (
	MIN_DAYS_TO_SCHEDULE = 3
	MAX_DAYS_TO_SCHEDULE = 30
	CLINIC_OPENING_HOUR  = 8
	CLINIC_CLOSING_HOUR  = 20
)

type Appointment struct {
	id            AppointmentId
	service       ClinicService
	scheduledDate time.Time
	status        AppointmentStatus
	reason        string
	notes         *string
	ownerId       int
	vetId         *vetDomain.VetId
	petId         petDomain.PetId
	createdAt     time.Time
	updatedAt     time.Time
}

func NewAppointment(id AppointmentId, petId petDomain.PetId, ownerId int, vetId *vetDomain.VetId, service ClinicService, scheduledDate time.Time, status AppointmentStatus, createdAt, updatedAt time.Time) *Appointment {
	return &Appointment{
		id:            id,
		petId:         petId,
		ownerId:       ownerId,
		vetId:         vetId,
		service:       service,
		scheduledDate: scheduledDate,
		status:        status,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

func NilAppointmentId() AppointmentId {
	return AppointmentId{IntegerId: shared.NilIntegerId()}
}

func (a *Appointment) RescheduleAppointment(newDate time.Time) {
	a.scheduledDate = newDate
	a.status = StatusRescheduled
	a.updatedAt = time.Now()
}

func (a *Appointment) ValidateFields() error {
	if a.id.Equals(shared.NilIntegerId()) {
		return errors.New("appointment ID cannot be nil")
	}

	if a.petId.Equals(petDomain.PetId{}) {
		return errors.New("pet ID cannot be nil")
	}

	if a.ownerId <= 0 {
		return errors.New("owner ID must be greater than zero")
	}

	if a.scheduledDate.IsZero() {
		return errors.New("scheduled date cannot be zero")
	}

	if err := a.ValidateRequestSchedule(); err != nil {
		return err
	}

	return nil
}

func (a *Appointment) Cancel() error {
	a.status = StatusCancelled
	a.updatedAt = time.Now()

	if a.scheduledDate.Before(time.Now()) {
		return errors.New("cannot cancel an appointment that is already in the past")
	}

	if a.status == StatusCompleted || a.status == StatusNotPresented {
		return errors.New("cannot cancel an appointment that is already completed or marked as not presented")
	}

	if a.status == StatusCancelled {
		return errors.New("appointment is already cancelled")
	}

	return nil
}

func (a *Appointment) CompleteAppointment() {
	a.status = StatusCompleted
	a.updatedAt = time.Now()
}

func (a *Appointment) MarkAsNotPresented() {
	a.status = StatusNotPresented
	a.updatedAt = time.Now()
}

func (a *Appointment) ValidateRequestSchedule() error {
	now := time.Now()

	if a.scheduledDate.IsZero() {
		return errors.New("schedule cannot be zero")
	}

	if a.scheduledDate.Before(now) {
		return errors.New("schedule cannot be in the past")
	}

	if a.scheduledDate.Before(now.AddDate(0, 0, MIN_DAYS_TO_SCHEDULE)) {
		return errors.New("schedule must be at least 3 days in advance")
	}

	if a.scheduledDate.Weekday() == time.Saturday || a.scheduledDate.Weekday() == time.Sunday {
		return errors.New("appointments cannot be scheduled on weekends")
	}

	return nil
}

// Get/Set
func (a *Appointment) GetId() AppointmentId {
	return a.id
}

func (a *Appointment) GetPetId() petDomain.PetId {
	return a.petId
}

func (a *Appointment) GetOwnerId() int {
	return a.ownerId
}

func (a *Appointment) GetVetId() *vetDomain.VetId {
	return a.vetId
}

func (a *Appointment) GetService() ClinicService {
	return a.service
}

func (a *Appointment) GetScheduledDate() time.Time {
	return a.scheduledDate
}

func (a *Appointment) GetStatus() AppointmentStatus {
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

func (a *Appointment) SetVetId(vetId *vetDomain.VetId) {
	a.vetId = vetId
}

func (a *Appointment) SetService(service ClinicService) {
	a.service = service
}

func (a *Appointment) SetScheduledDate(scheduledDate time.Time) {
	a.scheduledDate = scheduledDate
}

func (a *Appointment) SetStatus(status AppointmentStatus) {
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

func (a *Appointment) SetId(id AppointmentId) {
	a.id = id
}
