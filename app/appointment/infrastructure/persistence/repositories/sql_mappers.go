package appointmentRepository

import (
	appointDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/appointment/domain"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func sqlRowToDomain(row sqlc.Appoinment) (*appointDomain.Appointment, error) {
	builder := &appointDomain.AppointmentBuilder{}
	id, err := appointDomain.NewAppointmentId(row.ID)
	if err != nil {
		return nil, err
	}

	vetId, err := vetDomain.NewVeterinarianId(int(row.VeterinarianID))
	if err != nil {
		return nil, err
	}
	petId, err := petDomain.NewPetId(int(row.PetID))
	if err != nil {
		return nil, err
	}

	statusEnum, err := appointDomain.NewAppointmentStatus(string(row.Status))
	if err != nil {
		return nil, err
	}

	builder.WithID(id)
	builder.WithOwnerID(int(row.OwnerID))
	builder.WithVetID(&vetId)
	builder.WithPetID(petId)
	builder.WithStatus(statusEnum)
	builder.WithScheduledDate(row.ScheduleDate.Time)
	builder.WithTimestamps(row.CreatedAt.Time, row.UpdatedAt.Time)
	builder.WithReason(row.Reason)

	if row.Notes.Valid {
		builder.WithNotes(&row.Notes.String)
	}

	return builder.Build()
}

func sqlRowsToDomainList(rows []sqlc.Appoinment) ([]appointDomain.Appointment, error) {
	var appointments []appointDomain.Appointment
	for _, row := range rows {
		appointment, err := sqlRowToDomain(row)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, *appointment)
	}
	return appointments, nil
}

func domainToCreateParams(appointment *appointDomain.Appointment) sqlc.CreateAppoinmentParams {
	return sqlc.CreateAppoinmentParams{
		OwnerID:        int32(appointment.GetOwnerId()),
		VeterinarianID: int32(appointment.GetId().GetValue()),
		PetID:          int32(appointment.GetPetId().GetValue()),
		ScheduleDate:   pgtype.Timestamptz{Time: appointment.GetScheduledDate(), Valid: true},
		Notes:          pgtype.Text{String: *appointment.GetNotes(), Valid: appointment.GetNotes() != nil},
		Status:         models.AppointmentStatus(string(appointment.GetStatus())),
	}
}

func domainToUpdateParams(appointment *appointDomain.Appointment) sqlc.UpdateAppoinmentParams {
	return sqlc.UpdateAppoinmentParams{
		ID:             int32(appointment.GetId().GetValue()),
		OwnerID:        int32(appointment.GetOwnerId()),
		VeterinarianID: int32(appointment.GetVetId().GetValue()),
		PetID:          int32(appointment.GetPetId().GetValue()),
		ScheduleDate:   pgtype.Timestamptz{Time: appointment.GetScheduledDate(), Valid: true},
		Notes:          pgtype.Text{String: *appointment.GetNotes(), Valid: appointment.GetNotes() != nil},
		Status:         models.AppointmentStatus(string(appointment.GetStatus())),
	}
}
