package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	apperror "clinic-vet-api/app/shared/error/application"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlRowToAppointment(row sqlc.Appointment) (*appointment.Appointment, error) {
	errorMessage := make([]string, 0)
	id := valueobject.NewAppointmentID(uint(row.ID))
	vetID := valueobject.NewEmployeeID(uint(row.EmployeeID.Int32))
	petID := valueobject.NewPetID(uint(row.PetID))
	statusEnum, err := enum.ParseAppointmentStatus(string(row.Status))
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	customerID := valueobject.NewCustomerID(uint(row.CustomerID))
	reason, err := enum.ParseVisitReason(row.Reason)
	if err != nil {
		errorMessage = append(errorMessage, err.Error())
	}

	var notes *string
	if row.Notes.Valid {
		notes = &row.Notes.String
	}

	if len(errorMessage) > 0 {
		return &appointment.Appointment{}, apperror.MappingError(errorMessage, "sql", "appointmentEntity", "appointment")
	}

	return appointment.NewAppointment(
		id, petID, customerID,
		appointment.WithEmployeeID(&vetID),
		appointment.WithStatus(statusEnum),
		appointment.WithScheduledDate(row.ScheduleDate.Time),
		appointment.WithReason(reason),
		appointment.WithNotes(notes))
}

func sqlRowsToAppointments(rows []sqlc.Appointment) ([]appointment.Appointment, error) {
	var appointments []appointment.Appointment

	if len(rows) == 0 {
		return []appointment.Appointment{}, nil
	}

	for _, row := range rows {
		appointment, err := sqlRowToAppointment(row)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, *appointment)
	}
	return appointments, nil
}

func appointmentToCreateParams(appointment *appointment.Appointment) sqlc.CreateAppointmentParams {
	return sqlc.CreateAppointmentParams{
		CustomerID:   int32(appointment.CustomerID().Value()),
		EmployeeID:   pgtype.Int4{Int32: int32(appointment.EmployeeID().Value()), Valid: appointment.EmployeeID().IsZero()},
		PetID:        int32(appointment.PetID().Value()),
		ScheduleDate: pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Notes:        pgtype.Text{String: *appointment.Notes(), Valid: appointment.Notes() != nil},
		Status:       models.AppointmentStatus(string(appointment.Status())),
	}
}

func appointmentToUpdateParams(appointment *appointment.Appointment) sqlc.UpdateAppointmentParams {
	return sqlc.UpdateAppointmentParams{
		ID:           int32(appointment.ID().Value()),
		CustomerID:   int32(appointment.CustomerID().Value()),
		EmployeeID:   pgtype.Int4{Int32: int32(appointment.EmployeeID().Value()), Valid: appointment.EmployeeID().IsZero()},
		PetID:        int32(appointment.PetID().Value()),
		ScheduleDate: pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Notes:        pgtype.Text{String: *appointment.Notes(), Valid: appointment.Notes() != nil},
		Status:       models.AppointmentStatus(string(appointment.Status())),
	}
}
