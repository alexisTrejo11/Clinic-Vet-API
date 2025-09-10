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
	id := valueobject.NewAppointmentID(uint(row.ID))
	vetID := valueobject.NewEmployeeID(uint(row.EmployeeID.Int32))
	petID := valueobject.NewPetID(uint(row.PetID))
	statusEnum, err := enum.ParseAppointmentStatus(string(row.Status))
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	ownerID := valueobject.NewCustomerID(uint(row.CustomerID))
	reason, err := enum.ParseVisitReason(row.Reason)
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

	return appointment.NewAppointment(
		id, petID, ownerID,
		appointment.WithEmployeeID(&vetID),
		appointment.WithStatus(statusEnum),
		appointment.WithScheduledDate(row.ScheduleDate.Time),
		appointment.WithReason(reason),
		appointment.WithNotes(notes))
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
		CustomerID:   int32(appointment.CustomerID().Value()),
		EmployeeID:   pgtype.Int4{Int32: int32(appointment.EmployeeID().Value()), Valid: appointment.EmployeeID().IsZero()},
		PetID:        int32(appointment.PetID().Value()),
		ScheduleDate: pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Notes:        pgtype.Text{String: *appointment.Notes(), Valid: appointment.Notes() != nil},
		Status:       models.AppointmentStatus(string(appointment.Status())),
	}
}

func domainToUpdateParams(appointment *appointment.Appointment) sqlc.UpdateAppoinmentParams {
	return sqlc.UpdateAppoinmentParams{
		ID:           int32(appointment.ID().Value()),
		CustomerID:   int32(appointment.CustomerID().Value()),
		EmployeeID:   pgtype.Int4{Int32: int32(appointment.EmployeeID().Value()), Valid: appointment.EmployeeID().IsZero()},
		PetID:        int32(appointment.PetID().Value()),
		ScheduleDate: pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Notes:        pgtype.Text{String: *appointment.Notes(), Valid: appointment.Notes() != nil},
		Status:       models.AppointmentStatus(string(appointment.Status())),
	}
}
