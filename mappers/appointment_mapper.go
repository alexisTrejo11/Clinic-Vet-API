package mappers

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/db"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type AppointmentMappers struct {
	NamesRepository repository.NamesRepository
}

func NewAppointmentMappers() *AppointmentMappers {
	dbConn := db.InitDb()
	queries := sqlc.New(dbConn)
	namesRepository := repository.NewNamesRepository(queries)

	return &AppointmentMappers{
		NamesRepository: *namesRepository,
	}
}

func (AppointmentMappers) MapInsertDTOToInsertParams(appointmentInsertDTO DTOs.AppointmentInsertDTO, ownerID int32) sqlc.CreateAppointmentParams {
	return sqlc.CreateAppointmentParams{
		PetID:   appointmentInsertDTO.PetID,
		OwnerID: ownerID,
		Service: appointmentInsertDTO.Service,
		Status:  "pending",
		Date:    pgtype.Timestamp{Time: appointmentInsertDTO.Date, Valid: true},
	}
}

func (AppointmentMappers) MapSqlcEntityToDTO(appointment sqlc.Appointment) DTOs.AppointmentDTO {
	return DTOs.AppointmentDTO{
		Id:      appointment.ID,
		PetID:   appointment.PetID,
		VetID:   appointment.VetID.Int32,
		Service: appointment.Service,
		OwnerID: appointment.OwnerID,
		Status:  appointment.Status,
		Date:    appointment.CreatedAt.Time,
	}
}

func (am AppointmentMappers) MapSqlcEntityToNamedDTO(appointment sqlc.Appointment) DTOs.AppointmentNamedDTO {
	appointmentNames, err := am.NamesRepository.GetAppointmentRelationshipNames(appointment)
	if err != nil {
		return DTOs.AppointmentNamedDTO{}
	}

	return DTOs.AppointmentNamedDTO{
		Pet:     appointmentNames.PetName,
		Service: appointment.Service,
		Vet:     appointmentNames.VeterinarianName,
		Owner:   appointmentNames.OwnerFullName,
		Status:  appointment.Status,
		Date:    appointment.CreatedAt.Time,
	}
}

func (AppointmentMappers) MapUpdateDTOToUpdateParams(appointmentUpdateDTO DTOs.AppointmentUpdateDTO, ownerID int32) sqlc.UpdateAppointmentParams {
	return sqlc.UpdateAppointmentParams{
		ID:      appointmentUpdateDTO.Id,
		PetID:   appointmentUpdateDTO.PetID,
		VetID:   pgtype.Int4{Int32: appointmentUpdateDTO.PetID, Valid: true},
		Service: appointmentUpdateDTO.Service,
		OwnerID: ownerID,
		Date:    pgtype.Timestamp{Time: appointmentUpdateDTO.Date, Valid: true},
	}
}
