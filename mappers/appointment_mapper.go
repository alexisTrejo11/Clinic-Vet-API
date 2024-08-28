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

func (AppointmentMappers) MapRequestInsertDTOToInsertParams(appointmentRequestInsertDTO DTOs.AppointmentRequestInsertDTO, ownerID int32) sqlc.RequestAppointmentParams {
	return sqlc.RequestAppointmentParams{
		PetID:   appointmentRequestInsertDTO.PetID,
		OwnerID: ownerID,
		Service: appointmentRequestInsertDTO.Service,
		Status:  "pending",
		Date:    pgtype.Timestamp{Time: appointmentRequestInsertDTO.Date, Valid: true},
	}
}

func (AppointmentMappers) MapInsertDTOToInsertParams(appointmentInsertDTO DTOs.AppointmentInsertDTO) sqlc.CreateAppointmentParams {
	return sqlc.CreateAppointmentParams{
		PetID:   appointmentInsertDTO.PetID,
		OwnerID: appointmentInsertDTO.OwnerID,
		Service: appointmentInsertDTO.Service,
		VetID:   pgtype.Int4{Int32: appointmentInsertDTO.VetID, Valid: true},
		Status:  appointmentInsertDTO.Status,
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
func (AppointmentMappers) MapUpdateDTOToUpdateParams(appointmentUpdateDTO DTOs.AppointmentUpdateDTO, appointment sqlc.Appointment) sqlc.UpdateAppointmentParams {
	updateAppointmentParams := sqlc.UpdateAppointmentParams{
		ID: appointment.ID,
	}
	if appointmentUpdateDTO.PetID != 0 {
		updateAppointmentParams.PetID = appointmentUpdateDTO.PetID
	} else {
		updateAppointmentParams.PetID = appointment.PetID
	}

	if appointmentUpdateDTO.VetID != 0 {
		updateAppointmentParams.VetID = pgtype.Int4{Int32: appointmentUpdateDTO.VetID, Valid: true}
	} else {
		updateAppointmentParams.VetID = appointment.VetID
	}

	if appointmentUpdateDTO.Service != "" {
		updateAppointmentParams.Service = appointmentUpdateDTO.Service
	} else {
		updateAppointmentParams.Service = appointment.Service
	}

	if appointmentUpdateDTO.OwnerID != 0 {
		updateAppointmentParams.OwnerID = appointmentUpdateDTO.OwnerID
	} else {
		updateAppointmentParams.OwnerID = appointment.OwnerID
	}

	if !appointmentUpdateDTO.Date.IsZero() {
		updateAppointmentParams.Date = pgtype.Timestamp{Time: appointmentUpdateDTO.Date, Valid: true}
	} else {
		updateAppointmentParams.Date = appointment.Date
	}

	if appointmentUpdateDTO.Status != "" {
		updateAppointmentParams.Status = appointmentUpdateDTO.Status
	} else {
		updateAppointmentParams.Status = appointment.Status
	}

	return updateAppointmentParams
}

func (AppointmentMappers) MapUpdateDTOToUpdateParams2(appointmentUpdateDTO DTOs.AppointmentUpdateDTO, ownerID int32) sqlc.UpdateOwnerAppointmentParams {
	return sqlc.UpdateOwnerAppointmentParams{
		ID:      appointmentUpdateDTO.Id,
		PetID:   appointmentUpdateDTO.PetID,
		Service: appointmentUpdateDTO.Service,
		Date:    pgtype.Timestamp{Time: appointmentUpdateDTO.Date, Valid: true},
	}
}

func (AppointmentMappers) MapUpdateOwnerDTOToUpdateOwnerParams(appointmentUpdateDTO DTOs.AppointmentOwnerUpdateDTO) sqlc.UpdateOwnerAppointmentParams {
	return sqlc.UpdateOwnerAppointmentParams{
		ID:      appointmentUpdateDTO.Id,
		PetID:   appointmentUpdateDTO.PetID,
		Service: appointmentUpdateDTO.Service,
		Date:    pgtype.Timestamp{Time: appointmentUpdateDTO.Date, Valid: true},
	}
}
