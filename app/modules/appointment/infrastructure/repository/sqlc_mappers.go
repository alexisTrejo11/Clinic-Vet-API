package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func sqlRowToAppointment(row sqlc.Appointment) (*appointment.Appointment, error) {
	id := valueobject.NewAppointmentID(uint(row.ID))
	petID := valueobject.NewPetID(uint(row.PetID))
	statusEnum := enum.AppointmentStatus(string(row.Status))
	var vetID *valueobject.EmployeeID

	if row.EmployeeID.Valid {
		vetIDValue := valueobject.NewEmployeeID(uint(row.EmployeeID.Int32))
		vetID = &vetIDValue
	}

	customerID := valueobject.NewCustomerID(uint(row.CustomerID))

	var notes *string
	if row.Notes.Valid {
		notes = &row.Notes.String
	}

	return appointment.NewAppointment(
		id, petID, customerID,
		appointment.WithEmployeeID(vetID),
		appointment.WithStatus(statusEnum),
		appointment.WithScheduledDate(row.ScheduledDate.Time),
		appointment.WithNotes(notes),
		appointment.WithTimestamps(row.CreatedAt.Time, row.UpdatedAt.Time),
	)
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

func appointmentToUpdateParams(appointment *appointment.Appointment) sqlc.UpdateAppointmentParams {
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

func (r *SqlcAppointmentRepository) toDomainEntity(row sqlc.FindAppointmentsBySpecRow) (appointment.Appointment, error) {
	id := valueobject.NewAppointmentID(uint(row.ID))
	petID := valueobject.NewPetID(uint(row.PetID))
	customerID := valueobject.NewCustomerID(uint(row.CustomerID))

	var vetID *valueobject.EmployeeID
	if row.EmployeeID.Valid {
		vetIDObj := valueobject.NewEmployeeID(uint(row.EmployeeID.Int32))
		vetID = &vetIDObj
	}

	statusEnum := enum.AppointmentStatus(row.Status)

	var notes *string
	if row.Notes.Valid {
		notes = &row.Notes.String
	}

	appt, err := appointment.NewAppointment(
		id, petID, customerID,
		appointment.WithEmployeeID(vetID),
		appointment.WithStatus(statusEnum),
		appointment.WithScheduledDate(row.ScheduledDate.Time),
		appointment.WithNotes(notes),
		appointment.WithTimestamps(row.CreatedAt.Time, row.UpdatedAt.Time),
	)
	if err != nil {
		return appointment.Appointment{}, r.wrapConversionError(err)
	}
	return *appt, nil
}

func (r *SqlcAppointmentRepository) ToDomainEntities(rows []sqlc.FindAppointmentsBySpecRow) ([]appointment.Appointment, error) {
	if len(rows) == 0 {
		return []appointment.Appointment{}, nil
	}

	appointments := make([]appointment.Appointment, len(rows))
	for i, row := range rows {
		appt, err := r.toDomainEntity(row)
		if err != nil {
			return nil, r.wrapConversionError(err)
		}
		appointments[i] = appt
	}
	return appointments, nil
}
