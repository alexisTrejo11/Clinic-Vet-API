package appointmentMapper

import (
	"time"

	appointmentDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/application/dtos"
	appointDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
)

type AppointmentMapper struct{}

func NewAppointmentMapper() *AppointmentMapper {
	return &AppointmentMapper{}
}

// ToAppointmentResponseDTO converts domain appointment to response DTO
func (m *AppointmentMapper) ToAppointmentResponseDTO(appointment *appointDomain.Appointment) appointmentDTOs.AppointmentResponseDTO {
	var vetId *int
	if appointment.GetVetId() != nil {
		id := appointment.GetVetId().GetValue()
		vetId = &id
	}

	return appointmentDTOs.AppointmentResponseDTO{
		Id:            appointment.GetId().GetValue(),
		PetId:         appointment.GetPetId().GetValue(),
		OwnerId:       appointment.GetOwnerId(),
		VetId:         vetId,
		Service:       appointment.GetService(),
		ScheduledDate: appointment.GetScheduledDate(),
		Status:        appointment.GetStatus(),
		CreatedAt:     appointment.GetCreatedAt(),
		UpdatedAt:     appointment.GetUpdatedAt(),
	}
}

// ToAppointmentDetailDTO converts domain appointment to detailed response DTO
func (m *AppointmentMapper) ToAppointmentDetailDTO(appointment *appointDomain.Appointment, pet *appointmentDTOs.PetSummaryDTO, owner *appointmentDTOs.OwnerSummaryDTO, vet *appointmentDTOs.VetSummaryDTO) appointmentDTOs.AppointmentDetailDTO {
	return appointmentDTOs.AppointmentDetailDTO{
		Id:            appointment.GetId().GetValue(),
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

// RequestDTOToDomain converts request DTO to domain appointment
func (m *AppointmentMapper) RequestDTOToDomain(dto appointmentDTOs.AppointmentCreateDTO) (*appointDomain.Appointment, error) {
	appointmentId := appointDomain.NilAppointmentId()

	petId, err := petDomain.NewPetId(dto.PetId)
	if err != nil {
		return nil, err
	}

	var vetId *vetDomain.VetId
	if dto.VetId != nil {
		id, err := vetDomain.NewVeterinarianId(*dto.VetId)
		if err != nil {
			return nil, err
		}
		vetId = &id
	}

	now := time.Now()
	appointment := appointDomain.NewAppointment(
		appointmentId,
		petId,
		dto.PetId,
		vetId,
		dto.Service,
		dto.ScheduledDate,
		appointDomain.StatusPending,
		now,
		now,
	)

	return appointment, nil
}

// UpdateDTOToDomain applies update DTO to existing domain appointment
func (m *AppointmentMapper) UpdateDTOToDomain(appointment *appointDomain.Appointment, dto appointmentDTOs.AppointmentUpdateDTO) error {
	if dto.VetId != nil {
		if *dto.VetId != 0 {
			vetId, err := vetDomain.NewVeterinarianId(*dto.VetId)
			if err != nil {
				return err
			}
			appointment.SetVetId(&vetId)
		} else {
			appointment.SetVetId(nil)
		}
	}

	if dto.Service != nil {
		appointment.SetService(*dto.Service)
	}

	if dto.ScheduledDate != nil {
		appointment.SetScheduledDate(*dto.ScheduledDate)
	}

	appointment.SetUpdatedAt(time.Now())
	return nil
}

// OwnerUpdateDTOToDomain applies owner update DTO to existing domain appointment
func (m *AppointmentMapper) OwnerUpdateDTOToDomain(appointment *appointDomain.Appointment, dto appointmentDTOs.AppointmentOwnerUpdateDTO) error {
	if dto.Service != nil {
		appointment.SetService(*dto.Service)
	}

	if dto.ScheduledDate != nil {
		appointment.SetScheduledDate(*dto.ScheduledDate)
	}

	appointment.SetUpdatedAt(time.Now())
	return nil
}

// VetUpdateDTOToDomain applies vet update DTO to existing domain appointment
func (m *AppointmentMapper) VetUpdateDTOToDomain(appointment *appointDomain.Appointment, dto appointmentDTOs.AppointmentVetUpdateDTO) {
	appointment.SetStatus(dto.Status)
	appointment.SetUpdatedAt(time.Now())
}

// ToCreateAppointmentResponse creates a response for appointment creation
func (m *AppointmentMapper) ToCreateAppointmentResponse(appointment *appointDomain.Appointment) appointmentDTOs.CreateAppointmentResponseDTO {
	return appointmentDTOs.CreateAppointmentResponseDTO{
		Appointment: m.ToAppointmentResponseDTO(appointment),
		Message:     "Appointment requested successfully",
	}
}

// ToCancelAppointmentResponse creates a response for appointment cancellation
func (m *AppointmentMapper) ToCancelAppointmentResponse(appointmentId int) appointmentDTOs.CancelAppointmentResponseDTO {
	return appointmentDTOs.CancelAppointmentResponseDTO{
		AppointmentId: appointmentId,
		Status:        string(appointDomain.StatusCancelled),
		Message:       "Appointment cancelled successfully",
		CancelledAt:   time.Now(),
	}
}

// ToSearchCriteria converts search DTO to repository search criteria
func (m *AppointmentMapper) ToSearchCriteria(dto appointmentDTOs.AppointmentSearchDTO) map[string]interface{} {
	criteria := make(map[string]interface{})

	if dto.OwnerId != nil {
		criteria["owner_id"] = *dto.OwnerId
	}
	if dto.PetId != nil {
		criteria["pet_id"] = *dto.PetId
	}
	if dto.VetId != nil {
		criteria["vet_id"] = *dto.VetId
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
