// Package mapper contains all the operations to map domain entity to output dtos
package mapper

import (
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/appointment/application/dto"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type AppointmentMapper struct{}

func NewAppointmentMapper() *AppointmentMapper {
	return &AppointmentMapper{}
}

// ToAppointmentResponse converts domain appointment to response
func (m *AppointmentMapper) ToAppointmentResponse(appointment *appointment.Appointment) dto.AppointmentResponse {
	var vetID *int
	if appointment.VetID() != nil {
		v := appointment.VetID().Value()
		vetID = &v
	}

	return dto.AppointmentResponse{
		ID:            appointment.ID().Value(),
		PetID:         appointment.PetID().Value(),
		OwnerID:       appointment.OwnerID().Value(),
		VetID:         vetID,
		Service:       appointment.Service(),
		ScheduledDate: appointment.ScheduledDate(),
		Status:        appointment.Status(),
		CreatedAt:     appointment.CreatedAt(),
		UpdatedAt:     appointment.UpdatedAt(),
	}
}

// ToAppointmentDetail converts domain appointment to detailed response
func (m *AppointmentMapper) ToAppointmentDetail(
	appointment *appointment.Appointment,
	pet *dto.PetSummary,
	owner *dto.OwnerSummary,
	vet *dto.VetSummary,
) dto.AppointmentDetail {
	return dto.AppointmentDetail{
		ID:            appointment.ID().Value(),
		Pet:           pet,
		Owner:         owner,
		Veterinarian:  vet,
		Service:       appointment.Service(),
		ScheduledDate: appointment.ScheduledDate(),
		Status:        appointment.Status(),
		CreatedAt:     appointment.CreatedAt(),
		UpdatedAt:     appointment.UpdatedAt(),
	}
}

// RequestToDomain converts request to domain appointment
func (m *AppointmentMapper) RequestToDomain(dto dto.AppointmentCreate) (appointment.Appointment, error) {
	errorsMessages := make([]string, 0)

	petID, err := valueobject.NewPetID(dto.PetID)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	ownerID, err := valueobject.NewOwnerID(dto.OwnerID)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	var vetID *valueobject.VetID
	if dto.VetID != nil {
		id, err := valueobject.NewVetID(*dto.VetID)
		if err != nil {
			errorsMessages = append(errorsMessages, err.Error())
		} else {
			vetID = &id
		}
	}

	appointmentID, err := valueobject.NewAppointmentID(0)
	if err != nil {
		errorsMessages = append(errorsMessages, err.Error())
	}

	if len(errorsMessages) > 0 {
		return appointment.Appointment{}, apperror.MappingError(errorsMessages, "createDTO", "domain", "appointment")
	}

	appt, err := appointment.NewAppointment(
		appointmentID,
		petID,
		ownerID,
		appointment.WithVetID(vetID),
		appointment.WithService(dto.Service),
		appointment.WithScheduledDate(dto.ScheduledDate),
		appointment.WithNotes(dto.Notes),
	)
	if err != nil {
		return appointment.Appointment{}, fmt.Errorf("failed to create appointment: %w", err)
	}

	return *appt, nil
}

// ToCreateAppointmentResponse creates a response for appointment creation
func (m *AppointmentMapper) ToCreateAppointmentResponse(appointment *appointment.Appointment) dto.CreateAppointmentResponse {
	return dto.CreateAppointmentResponse{
		Appointment: m.ToAppointmentResponse(appointment),
		Message:     "Appointment requested successfully",
	}
}

// ToCancelAppointmentResponse creates a response for appointment cancellation
func (m *AppointmentMapper) ToCancelAppointmentResponse(appointmentID valueobject.AppointmentID) dto.CancelAppointmentResponse {
	return dto.CancelAppointmentResponse{
		AppointmentID: appointmentID.Value(),
		Status:        enum.AppointmentStatusCancelled.DisplayName(),
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
