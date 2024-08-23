package services

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
)

type AppointmentService interface {
	CreateAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO) error
	GetAppointmentById(appointmentId int32) (*DTOs.AppointmentDTO, error)
}

type appointmentServiceImpl struct {
	appointMappers        mappers.AppointmentMappers
	appointmentRepository repository.AppointmentRepository
}

func NewAppoinmentService(appointmentRepository repository.AppointmentRepository) AppointmentService {
	return &appointmentServiceImpl{
		appointmentRepository: appointmentRepository,
	}
}

func (as *appointmentServiceImpl) CreateAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO) error {
	appointmentParams := as.appointMappers.MapInsertDTOToInsertParams(appointmentInsertDTO)
	if err := as.appointmentRepository.CreateAppointment(appointmentParams); err != nil {
		return err
	}

	return nil
}

func (as *appointmentServiceImpl) GetAppointmentById(appointmentId int32) (*DTOs.AppointmentDTO, error) {
	appointment, err := as.appointmentRepository.GetAppointmentByID(appointmentId)
	if err != nil {
		return nil, err
	}

	appointmentDTO := as.appointMappers.MapSqlcEntityToToDTO(*appointment)

	return &appointmentDTO, nil
}

// Todo: Handle Optinal Fields
func (as *appointmentServiceImpl) UpdateAppointment(appointmentId int32) (*DTOs.AppointmentDTO, error) {
	appointment, err := as.appointmentRepository.GetAppointmentByID(appointmentId)
	if err != nil {
		return nil, err
	}

	appointmentDTO := as.appointMappers.MapSqlcEntityToToDTO(*appointment)

	return &appointmentDTO, nil
}
