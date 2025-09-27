package repository

import (
	appt "clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlcToEntity(row sqlc.Appointment) *appt.Appointment {
	var employeeID *valueobject.EmployeeID
	if row.EmployeeID.Valid {
		employeeIDValue := valueobject.NewEmployeeID(uint(row.EmployeeID.Int32))
		employeeID = &employeeIDValue
	}

	var notes *string
	if row.Notes.Valid {
		notes = &row.Notes.String
	}

	return appt.NewAppointmentBuilder().
		WithID(valueobject.NewAppointmentID(uint(row.ID))).
		WithPetID(valueobject.NewPetID(uint(row.PetID))).
		WithCustomerID(valueobject.NewCustomerID(uint(row.CustomerID))).
		WithEmployeeID(employeeID).
		WithStatus(enum.AppointmentStatus(string(row.Status))).
		WithScheduledDate(row.ScheduledDate.Time).
		WithNotes(notes).
		WithTimestamps(row.CreatedAt.Time, row.UpdatedAt.Time).
		Build()
}

func sqlcRowsToEntities(rows []sqlc.Appointment) []appt.Appointment {
	var appointments []appt.Appointment
	if len(rows) == 0 {
		return []appt.Appointment{}
	}

	for _, row := range rows {
		appointments = append(appointments, *sqlcToEntity(row))
	}
	return appointments
}

func appointmentToCreateParams(appointment *appt.Appointment) sqlc.CreateAppointmentParams {
	params := sqlc.CreateAppointmentParams{
		CustomerID:    int32(appointment.CustomerID().Value()),
		PetID:         int32(appointment.PetID().Value()),
		ScheduledDate: pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Status:        models.AppointmentStatus(string(appointment.Status())),
		ClinicService: models.ClinicService(string(appointment.Service().String())),
	}

	if appointment.EmployeeID() != nil {
		params.EmployeeID = pgtype.Int4{
			Int32: int32(appointment.EmployeeID().Value()),
			Valid: true,
		}
	} else {
		params.EmployeeID = pgtype.Int4{Valid: false}
	}

	if appointment.Notes() != nil {
		params.Notes = pgtype.Text{
			String: *appointment.Notes(),
			Valid:  true,
		}
	} else {
		params.Notes = pgtype.Text{Valid: false}
	}

	return params
}

func appointmentToUpdateParams(appointment *appt.Appointment) sqlc.UpdateAppointmentParams {
	return sqlc.UpdateAppointmentParams{
		ID:            int32(appointment.ID().Value()),
		CustomerID:    int32(appointment.CustomerID().Value()),
		EmployeeID:    pgtype.Int4{Int32: int32(appointment.EmployeeID().Value()), Valid: appointment.EmployeeID().IsZero()},
		PetID:         int32(appointment.PetID().Value()),
		ScheduledDate: pgtype.Timestamptz{Time: appointment.ScheduledDate(), Valid: true},
		Notes:         pgtype.Text{String: *appointment.Notes(), Valid: appointment.Notes() != nil},
		Status:        models.AppointmentStatus(string(appointment.Status())),
	}
}

func toEntity(row sqlc.FindAppointmentsBySpecRow) appt.Appointment {
	var vetID *valueobject.EmployeeID
	if row.EmployeeID.Valid {
		vetIDObj := valueobject.NewEmployeeID(uint(row.EmployeeID.Int32))
		vetID = &vetIDObj
	}

	var notes *string
	if row.Notes.Valid {
		notes = &row.Notes.String
	}

	appt := appt.NewAppointmentBuilder().
		WithID(valueobject.NewAppointmentID(uint(row.ID))).
		WithPetID(valueobject.NewPetID(uint(row.PetID))).
		WithCustomerID(valueobject.NewCustomerID(uint(row.CustomerID))).
		WithEmployeeID(vetID).
		WithStatus(enum.AppointmentStatus(row.Status)).
		WithScheduledDate(row.ScheduledDate.Time).
		WithNotes(notes).
		WithTimestamps(row.CreatedAt.Time, row.UpdatedAt.Time).
		Build()

	return *appt
}

func toEntities(rows []sqlc.FindAppointmentsBySpecRow) []appt.Appointment {
	if len(rows) == 0 {
		return []appt.Appointment{}
	}

	appointments := make([]appt.Appointment, len(rows))
	for i, row := range rows {
		appointments[i] = toEntity(row)
	}
	return appointments
}
