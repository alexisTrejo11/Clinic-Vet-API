package repositoryimpl

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func sqlRowToDomain(row sqlc.Appoinment) (*appointment.Appointment, error) {
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

	statusEnum, err := enum.ParseAppointmentStatus(string(row.Status))
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
		return &appointment.Appointment{}, apperror.MappingError(errorMessage, "sql", "domainEntity", "appointment")
	}

	return appointment.NewAppointmentBuilder().
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

func sqlRowsToDomainList(rows []sqlc.Appoinment) ([]appointment.Appointment, error) {
	var appointments []appointment.Appointment

	if len(rows) == 0 {
		return []appointment.Appointment{}, nil
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

func domainToCreateParams(appointment *appointment.Appointment) sqlc.CreateAppoinmentParams {
	return sqlc.CreateAppoinmentParams{
		OwnerID:        int32(appointment.OwnerID().Value()),
		VeterinarianID: pgtype.Int4{Int32: int32(appointment.VetID().Value()), Valid: appointment.VetID().IsZero()},
		PetID:          int32(appointment.PetID().Value()),
		ScheduleDate:   pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Notes:          pgtype.Text{String: *appointment.Notes(), Valid: appointment.Notes() != nil},
		Status:         models.AppointmentStatus(string(appointment.Status())),
	}
}

func domainToUpdateParams(appointment *appointment.Appointment) sqlc.UpdateAppoinmentParams {
	return sqlc.UpdateAppoinmentParams{
		ID:             int32(appointment.ID().Value()),
		OwnerID:        int32(appointment.OwnerID().Value()),
		VeterinarianID: pgtype.Int4{Int32: int32(appointment.VetID().Value()), Valid: appointment.VetID().IsZero()},
		PetID:          int32(appointment.PetID().Value()),
		ScheduleDate:   pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Notes:          pgtype.Text{String: *appointment.Notes(), Valid: appointment.Notes() != nil},
		Status:         models.AppointmentStatus(string(appointment.Status())),
	}
}
