package repositrories

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type AppointmentRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewAppointmentRepository(queries *sqlc.Queries) *AppointmentRepositoryImpl {
	return &AppointmentRepositoryImpl{
		queries: queries,
	}
}

func (r AppointmentRepositoryImpl) Create(params sqlc.CreateAppointmentParams) (*sqlc.Appointment, error) {
	appointment, err := r.queries.CreateAppointment(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (r AppointmentRepositoryImpl) GetByID(appointmentID int32) (*sqlc.Appointment, error) {
	appointment, err := r.queries.GetAppointmentByID(context.Background(), appointmentID)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (r AppointmentRepositoryImpl) GetByOwnerID(ownerID int32) ([]sqlc.Appointment, error) {
	appointmentList, err := r.queries.ListAppointmentsByOwnerID(context.Background(), ownerID)
	if err != nil {
		return nil, err
	}

	return appointmentList, nil
}

func (r appointmentRepositoryImpl) Update(updateParams sqlc.UpdateAppointmentParams) error {
	err := r.queries.UpdateAppointment(context.Background(), updateParams)
	if err != nil {
		return err
	}

	return nil
}

func (r appointmentRepositoryImpl) Delete(appointmentID int32) error {
	if err := r.queries.DeleteAppointment(context.Background(), appointmentID); err != nil {
		return err
	}

	return nil
}

// TODO: MAKE ENUMS VALIDATIONS
func (r AppointmentRepositoryImpl) UpdateAppointmentStatus(appointmentID int32, status string) error {
	updateParams := sqlc.UpdateAppointmentStatusParams{
		ID:     appointmentID,
		Status: status,
	}

	if err := r.queries.UpdateAppointmentStatus(context.Background(), updateParams); err != nil {
		return err
	}

	return nil
}

func (r AppointmentRepositoryImpl) RequestAppointment(params sqlc.RequestAppointmentParams) (*sqlc.Appointment, error) {
	appointment, err := r.queries.RequestAppointment(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (r AppointmentRepositoryImpl) UpdateOwnerAppointment(updateParams sqlc.UpdateOwnerAppointmentParams) error {
	err := r.queries.UpdateOwnerAppointment(context.Background(), updateParams)
	if err != nil {
		return err
	}

	return nil
}
