package appointDomain

import (
	"time"

	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
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
