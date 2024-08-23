package mappers

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type AppointmentMappers struct{}

func (AppointmentMappers) MapInsertDTOToInsertParams(appointmentInsertDTO DTOs.AppointmentInsertDTO) sqlc.CreateAppointmentParams {
	return sqlc.CreateAppointmentParams{
		PetID:   appointmentInsertDTO.PetID,
		VetID:   appointmentInsertDTO.VetID,
		Service: appointmentInsertDTO.Service,
		Date:    pgtype.Timestamp{Time: appointmentInsertDTO.Date, Valid: true},
	}
}

func (AppointmentMappers) MapSqlcEntityToToDTO(appointment sqlc.Appointment) DTOs.AppointmentDTO {
	return DTOs.AppointmentDTO{
		Id:      appointment.ID,
		PetID:   appointment.PetID,
		VetID:   appointment.VetID,
		Service: appointment.Service,
		Date:    appointment.CreatedAt.Time,
	}
}

func (AppointmentMappers) MapUpdateDTOToUpdateParams(appointmentUpdateDTO DTOs.AppointmentUpdateDTO) sqlc.UpdateAppointmentParams {
	return sqlc.UpdateAppointmentParams{
		ID:      appointmentUpdateDTO.Id,
		PetID:   appointmentUpdateDTO.PetID,
		VetID:   appointmentUpdateDTO.VetID,
		Service: appointmentUpdateDTO.Service,
		Date:    pgtype.Timestamp{Time: appointmentUpdateDTO.Date, Valid: true},
	}
}
