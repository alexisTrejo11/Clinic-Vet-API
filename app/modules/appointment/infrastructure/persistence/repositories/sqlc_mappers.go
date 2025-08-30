package sqlcrepository

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func sqlRowToDomain(row sqlc.Appoinment) (*entity.Appointment, error) {
	builder := &entity.AppointmentBuilder{}
	id, err := valueobject.NewAppointmentID(int(row.ID))
	if err != nil {
		return nil, err
	}

	vetID, err := valueobject.NewVetID(int(row.VeterinarianID.Int32))
	if err != nil {
		return nil, err
	}
	petID, err := valueobject.NewPetID(int(row.PetID))
	if err != nil {
		return nil, err
	}

	statusEnum, err := enum.NewAppointmentStatus(string(row.Status))
	if err != nil {
		return nil, err
	}

	builder.WithID(id)
	builder.WithOwnerID(int(row.OwnerID))
	builder.WithVetID(&vetID)
	builder.WithPetID(petID)
	builder.WithStatus(statusEnum)
	builder.WithScheduledDate(row.ScheduleDate.Time)
	builder.WithTimestamps(row.CreatedAt.Time, row.UpdatedAt.Time)
	builder.WithReason(row.Reason)

	if row.Notes.Valid {
		builder.WithNotes(&row.Notes.String)
	}

	return builder.Build()
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
		OwnerID:        int32(appointment.GetOwnerID()),
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
		OwnerID:        int32(appointment.GetOwnerID()),
		VeterinarianID: pgtype.Int4{Int32: int32(appointment.GetVetID().GetValue()), Valid: appointment.GetVetID().IsZero()},
		PetID:          int32(appointment.GetPetID().GetValue()),
		ScheduleDate:   pgtype.Timestamptz{Time: appointment.GetScheduledDate(), Valid: true},
		Notes:          pgtype.Text{String: *appointment.GetNotes(), Valid: appointment.GetNotes() != nil},
		Status:         models.AppointmentStatus(string(appointment.GetStatus())),
	}
}
