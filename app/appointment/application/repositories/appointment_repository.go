package repositories

import "example.com/at/backend/api-vet/app/container/sqlc"

type AppointmentRepository interface {
	Create(params sqlc.CreateAppointmentParams) (*sqlc.Appointment, error)
	GetByID(appointmentId int32) (*sqlc.Appointment, error)
	Update(updateParams sqlc.UpdateAppointmentParams) error
	Delete(appointmentId int32) error

	GetByOwnerID(ownerID int32) ([]sqlc.Appointment, error)
	UpdateStatus(appointmentID int32, status string) error
	Request(params sqlc.RequestAppointmentParams) (*sqlc.Appointment, error)
	UpdateOwner(updateParams sqlc.UpdateOwnerAppointmentParams) error
}
