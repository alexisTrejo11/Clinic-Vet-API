package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type AppointmentRepository interface {
	CreateAppointment(params sqlc.CreateAppointmentParams) (*sqlc.Appointment, error)
	GetAppointmentByID(appointmentId int32) (*sqlc.Appointment, error)
	GetAppointmentByOwnerID(ownerID int32) ([]sqlc.Appointment, error)
	UpdateAppointment(updateParams sqlc.UpdateAppointmentParams) error
	UpdateAppointmentStatus(appointmentID int32, status string) error
	DeleteAppointment(appointmentId int32) error
}

type appointmentRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewAppointmentRepository(queries *sqlc.Queries) AppointmentRepository {
	return &appointmentRepositoryImpl{
		queries: queries,
	}
}

func (ar appointmentRepositoryImpl) CreateAppointment(params sqlc.CreateAppointmentParams) (*sqlc.Appointment, error) {
	appointment, err := ar.queries.CreateAppointment(context.Background(), params)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (ar appointmentRepositoryImpl) GetAppointmentByID(appointmentID int32) (*sqlc.Appointment, error) {
	appointment, err := ar.queries.GetAppointmentByID(context.Background(), appointmentID)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (ar appointmentRepositoryImpl) GetAppointmentByOwnerID(ownerID int32) ([]sqlc.Appointment, error) {
	appointmentList, err := ar.queries.ListAppointmentsByOwnerID(context.Background(), ownerID)
	if err != nil {
		return nil, err
	}

	return appointmentList, nil
}

func (ar appointmentRepositoryImpl) UpdateAppointment(updateParams sqlc.UpdateAppointmentParams) error {
	err := ar.queries.UpdateAppointment(context.Background(), updateParams)
	if err != nil {
		return err
	}

	return nil
}

func (ar appointmentRepositoryImpl) DeleteAppointment(appointmentID int32) error {
	if err := ar.queries.DeleteAppointment(context.Background(), appointmentID); err != nil {
		return err
	}

	return nil
}

// TODO: MAKE ENUMS VALIDATIONS
func (ar appointmentRepositoryImpl) UpdateAppointmentStatus(appointmentID int32, status string) error {
	updateParams := sqlc.UpdateAppointmentStatusParams{
		ID:     appointmentID,
		Status: status,
	}

	if err := ar.queries.UpdateAppointmentStatus(context.Background(), updateParams); err != nil {
		return err
	}

	return nil
}
