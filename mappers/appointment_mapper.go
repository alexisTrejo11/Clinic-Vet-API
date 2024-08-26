package mappers

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type AppointmentMappers struct{}

func (AppointmentMappers) MapInsertDTOToInsertParams(appointmentInsertDTO DTOs.AppointmentInsertDTO, ownerID int32) sqlc.CreateAppointmentParams {
	return sqlc.CreateAppointmentParams{
		PetID:   appointmentInsertDTO.PetID,
		VetID:   appointmentInsertDTO.VetID,
		OwnerID: ownerID,
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
		OwnerID: appointment.OwnerID,
		Date:    appointment.CreatedAt.Time,
	}
}

func (AppointmentMappers) MapUpdateDTOToUpdateParams(appointmentUpdateDTO DTOs.AppointmentUpdateDTO, ownerID int32) sqlc.UpdateAppointmentParams {
	return sqlc.UpdateAppointmentParams{
		ID:      appointmentUpdateDTO.Id,
		PetID:   appointmentUpdateDTO.PetID,
		VetID:   appointmentUpdateDTO.VetID,
		Service: appointmentUpdateDTO.Service,
		OwnerID: ownerID,
		Date:    pgtype.Timestamp{Time: appointmentUpdateDTO.Date, Valid: true},
	}
}
