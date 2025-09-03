package repositoryimpl

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func sqlRowToDomain(row sqlc.Appoinment) (*entity.Appointment, error) {
	errorMessage := make([]string, 0)

	id, err := valueobject.NewAppointmentID(int(row.ID))
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	vetID, err := valueobject.NewVetID(int(row.VeterinarianID.Int32))
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}
	petID, err := valueobject.NewPetID(int(row.PetID))
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	statusEnum, err := enum.NewAppointmentStatus(string(row.Status))
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	ownerID, err := valueobject.NewOwnerID(int(row.OwnerID))
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	var notes *string
	if row.Notes.Valid {
		notes = &row.Notes.String
	}

	if len(errorMessage) > 0 {
		return &entity.Appointment{}, apperror.MappingError(errorMessage, "sql", "domainEntity", "appointment")
	}

	return entity.NewAppointmentBuilder().
		WithID(id).
		WithOwnerID(ownerID).
		WithVetID(&vetID).
		WithPetID(petID).
		WithStatus(statusEnum).
		WithScheduledDate(row.ScheduleDate.Time).
		WithTimestamps(row.CreatedAt.Time, row.UpdatedAt.Time).
		WithReason(row.Reason).
		WithNotes(notes).
		Build(), nil
}

func sqlRowsToDomainList(rows []sqlc.Appoinment) ([]entity.Appointment, error) {
	var appointments []entity.Appointment

	if len(rows) == 0 {
		return []entity.Appointment{}, nil
	}

	for _, row := range rows {
		appointment, err := sqlRowToDomain(row)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, *appointment)
	}
	return appointments, nil
}

func domainToCreateParams(appointment *entity.Appointment) sqlc.CreateAppoinmentParams {
	return sqlc.CreateAppoinmentParams{
		OwnerID:        int32(appointment.GetOwnerID().GetValue()),
		VeterinarianID: pgtype.Int4{Int32: int32(appointment.GetVetID().GetValue()), Valid: appointment.GetVetID().IsZero()},
		PetID:          int32(appointment.GetPetID().GetValue()),
		ScheduleDate:   pgtype.Timestamptz{Time: appointment.GetScheduledDate(), Valid: true},
		Notes:          pgtype.Text{String: *appointment.GetNotes(), Valid: appointment.GetNotes() != nil},
		Status:         models.AppointmentStatus(string(appointment.GetStatus())),
	}
}

func domainToUpdateParams(appointment *entity.Appointment) sqlc.UpdateAppoinmentParams {
	return sqlc.UpdateAppoinmentParams{
		ID:             int32(appointment.GetID().GetValue()),
		OwnerID:        int32(appointment.GetOwnerID().GetValue()),
		VeterinarianID: pgtype.Int4{Int32: int32(appointment.GetVetID().GetValue()), Valid: appointment.GetVetID().IsZero()},
		PetID:          int32(appointment.GetPetID().GetValue()),
		ScheduleDate:   pgtype.Timestamptz{Time: appointment.GetScheduledDate(), Valid: true},
		Notes:          pgtype.Text{String: *appointment.GetNotes(), Valid: appointment.GetNotes() != nil},
		Status:         models.AppointmentStatus(string(appointment.GetStatus())),
	}
}
