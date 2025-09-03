// Package mapper contains all the operations to map domain entity to output dtos
package mapper

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/dto"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type AppointmentMapper struct{}

func NewAppointmentMapper() *AppointmentMapper {
	return &AppointmentMapper{}
}

// ToAppointmentResponse converts domain appointment to response
func (m *AppointmentMapper) ToAppointmentResponse(appointment *entity.Appointment) dto.AppointmentResponse {
	var vetID *int
	if appointment.GetVetID() != nil {
		v := appointment.GetVetID().GetValue()
		vetID = &v
	}

	return dto.AppointmentResponse{
		ID:            appointment.GetID().GetValue(),
		PetID:         appointment.GetPetID().GetValue(),
		OwnerID:       appointment.GetOwnerID().GetValue(),
		VetID:         vetID,
		Service:       appointment.GetService(),
		ScheduledDate: appointment.GetScheduledDate(),
		Status:        appointment.GetStatus(),
		CreatedAt:     appointment.GetCreatedAt(),
		UpdatedAt:     appointment.GetUpdatedAt(),
	}
}

// ToAppointmentDetail converts domain appointment to detailed response
func (m *AppointmentMapper) ToAppointmentDetail(
	appointment *entity.Appointment,
	pet *dto.PetSummary,
	owner *dto.OwnerSummary,
	vet *dto.VetSummary,
) dto.AppointmentDetail {
	return dto.AppointmentDetail{
		ID:            appointment.GetID().GetValue(),
		Pet:           pet,
		Owner:         owner,
		Veterinarian:  vet,
		Service:       appointment.GetService(),
		ScheduledDate: appointment.GetScheduledDate(),
		Status:        appointment.GetStatus(),
		CreatedAt:     appointment.GetCreatedAt(),
		UpdatedAt:     appointment.GetUpdatedAt(),
	}
}

// RequestToDomain converts request  to domain appointment
func (m *AppointmentMapper) RequestToDomain(dto dto.AppointmentCreate) (entity.Appointment, error) {
	errorsMessages := make([]string, 0)
	petID, err := valueobject.NewPetID(dto.PetID)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	var vetID *valueobject.VetID
	if dto.VetID != nil {
		id, err := valueobject.NewVetID(*dto.VetID)
		if err != nil {
			errorsMessages = append(errorsMessages, err.Error())
		}
		vetID = &id
	}

	ownerID, err := valueobject.NewOwnerID(dto.OwnerID)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	now := time.Now()

	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	if len(errorsMessages) > 0 {
		return entity.Appointment{}, apperror.MappingError(errorsMessages, "createDTO", "domain", "appointment")
	}

	return *entity.
		NewAppointmentBuilder().
		WithPetID(petID).
		WithNotes(dto.Notes).
		WithOwnerID(ownerID).
		WithReason(*dto.Notes).
		WithTimestamps(now, now).
		WithScheduledDate(dto.ScheduledDate).
		WithVetID(vetID).
		WithService(dto.Service).
		Build(), nil
}

// ToCreateAppointmentResponse creates a response for appointment creation
func (m *AppointmentMapper) ToCreateAppointmentResponse(appointment *entity.Appointment) dto.CreateAppointmentResponse {
	return dto.CreateAppointmentResponse{
		Appointment: m.ToAppointmentResponse(appointment),
		Message:     "Appointment requested successfully",
	}
}

// ToCancelAppointmentResponse creates a response for appointment cancellation
func (m *AppointmentMapper) ToCancelAppointmentResponse(appointmentID valueobject.AppointmentID) dto.CancelAppointmentResponse {
	return dto.CancelAppointmentResponse{
		AppointmentID: appointmentID.GetValue(),
		Status:        string(enum.StatusCancelled),
		Message:       "App",
		CancelledAt:   time.Now(),
	}
}

// ToSearchCriteria converts search  to repository search criteria
func (m *AppointmentMapper) ToSearchCriteria(dto dto.AppointmentSearch) map[string]interface{} {
	criteria := make(map[string]interface{})

	if dto.OwnerID != nil {
		criteria["owner_id"] = *dto.OwnerID
	}
	if dto.PetID != nil {
		criteria["pet_id"] = *dto.PetID
	}
	if dto.VetID != nil {
		criteria["vet_id"] = *dto.VetID
	}
	if dto.Status != nil {
		criteria["status"] = *dto.Status
	}
	if dto.Service != nil {
		criteria["service"] = *dto.Service
	}
	if dto.StartDate != nil {
		criteria["start_date"] = *dto.StartDate
	}
	if dto.EndDate != nil {
		criteria["end_date"] = *dto.EndDate
	}

	return criteria
}
