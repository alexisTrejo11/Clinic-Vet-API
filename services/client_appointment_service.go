package services

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
)

type ClientAppointmentService interface {
	RequestAnAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO) (*DTOs.AppointmentDTO, error)
	GetAppointmentById(appointmentID int32) (*DTOs.AppointmentDTO, error)
	GetAppointmentByOnwerId(ownerID int32) ([]*DTOs.AppointmentDTO, error)
	UpdateAppointment(appointmentUpdateDTO DTOs.AppointmentUpdateDTO) (*DTOs.AppointmentDTO, error)
	CancelAppointmentById(appointmentID int32) error
	ValidAppointmentDeleted(appointmentID, owernID int32) error
}

type clientAppointmentServiceImpl struct {
	appointMappers        mappers.AppointmentMappers
	appointmentRepository repository.AppointmentRepository
}

func NewAppoinmentService(appointmentRepository repository.AppointmentRepository) *clientAppointmentServiceImpl {
	return &clientAppointmentServiceImpl{
		appointmentRepository: appointmentRepository,
	}
}

func (as *clientAppointmentServiceImpl) RequestAnAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO) (*DTOs.AppointmentDTO, error) {
	appointmentParams := as.appointMappers.MapInsertDTOToInsertParams(appointmentInsertDTO)
	appointment, err := as.appointmentRepository.CreateAppointment(appointmentParams)
	if err != nil {
		return nil, err
	}

	appointmentDTO := as.appointMappers.MapSqlcEntityToToDTO(*appointment)
	return &appointmentDTO, nil
}

func (as *clientAppointmentServiceImpl) GetAppointmentById(appointmentID int32) (*DTOs.AppointmentDTO, error) {
	appointment, err := as.appointmentRepository.GetAppointmentByID(appointmentID)
	if err != nil {
		return nil, err
	}

	appointmentDTO := as.appointMappers.MapSqlcEntityToToDTO(*appointment)

	return &appointmentDTO, nil
}

// Todo: Handle Optinal Fields
