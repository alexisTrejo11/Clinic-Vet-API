package services

import (
	"errors"

	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type ClientAppointmentService interface {
	RequestAnAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO, ownerID int32) (*DTOs.AppointmentDTO, error)
	GetAppointmentById(appointmentID int32) (*DTOs.AppointmentDTO, error)
	GetAppointmentByOwnerId(ownerID int32) ([]DTOs.AppointmentDTO, error)
	UpdateAppointment(appointmentUpdateDTO DTOs.AppointmentUpdateDTO, ownerID int32) error
	CancelAppointmentById(appointmentID int32) error
}

type clientAppointmentServiceImpl struct {
	appointMappers        mappers.AppointmentMappers
	appointmentRepository repository.AppointmentRepository
}

func NewAppointmentService(appointmentRepository repository.AppointmentRepository) ClientAppointmentService {
	return &clientAppointmentServiceImpl{
		appointmentRepository: appointmentRepository,
	}
}

func (as clientAppointmentServiceImpl) RequestAnAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO, ownerID int32) (*DTOs.AppointmentDTO, error) {
	appointmentParams := as.appointMappers.MapInsertDTOToInsertParams(appointmentInsertDTO, ownerID)
	appointment, err := as.appointmentRepository.CreateAppointment(appointmentParams)
	if err != nil {
		return nil, err
	}

	appointmentDTO := as.appointMappers.MapSqlcEntityToToDTO(*appointment)
	return &appointmentDTO, nil
}

func (as clientAppointmentServiceImpl) GetAppointmentById(appointmentID int32) (*DTOs.AppointmentDTO, error) {
	appointment, err := as.appointmentRepository.GetAppointmentByID(appointmentID)
	if err != nil {
		return nil, err
	}

	appointmentDTO := as.appointMappers.MapSqlcEntityToToDTO(*appointment)

	return &appointmentDTO, nil
}

func (as clientAppointmentServiceImpl) GetAppointmentByOwnerId(ownerID int32) ([]DTOs.AppointmentDTO, error) {
	appointments, err := as.appointmentRepository.GetAppointmentByOwnerID(ownerID)
	if err != nil {
		return nil, err
	}

	var appointmentsDTOs []DTOs.AppointmentDTO
	for _, appointment := range appointments {
		appointmentDTO := as.appointMappers.MapSqlcEntityToToDTO(appointment)
		appointmentsDTOs = append(appointmentsDTOs, appointmentDTO)
	}

	return appointmentsDTOs, nil
}

func (as clientAppointmentServiceImpl) UpdateAppointment(appointmentUpdateDTO DTOs.AppointmentUpdateDTO, ownerID int32) error {
	updateParams := as.appointMappers.MapUpdateDTOToUpdateParams(appointmentUpdateDTO, ownerID)

	if err := as.appointmentRepository.UpdateAppointment(updateParams); err != nil {
		return err
	}

	return nil
}

func (as clientAppointmentServiceImpl) CancelAppointmentById(appointmentID int32) error {
	appointment, _ := as.appointmentRepository.GetAppointmentByID(appointmentID)

	COMPLETED := pgtype.Text{String: "completed", Valid: true}
	CANCELED := pgtype.Text{String: "canceled", Valid: true}
	NO_SHOW := pgtype.Text{String: "no_show", Valid: true}

	if appointment.Status == COMPLETED || appointment.Status == NO_SHOW {
		return errors.New("appointment not allowed to be canceled")
	} else if appointment.Status == CANCELED {
		return errors.New("appointment already cancelled")
	} else {
		as.appointmentRepository.UpdateAppointmentStatus(appointmentID, CANCELED)
		return nil
	}

}
