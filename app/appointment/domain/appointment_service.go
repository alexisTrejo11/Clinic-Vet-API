package appointDomain

/*
type AppointmentService interface {
	//CRUD
	CreateAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO) error
	GetAppointmentById(appointmentID int32) (*DTOs.AppointmentDTO, error)
	UpdateAppointment(appointmentUpdateDTO DTOs.AppointmentUpdateDTO) error
	DeleteAppointmentById(appointmentID int32) error

	// Owner Functions
	RequestAnAppointment(AppointmentRequestInsertDTO DTOs.AppointmentRequestInsertDTO, ownerID int32) (*DTOs.AppointmentNamedDTO, error)
	UpdateOwnerAppointment(appointmentUpdateDTO DTOs.AppointmentOwnerUpdateDTO, ownerID int32) error
	CancelAppointmentById(appointmentID int32) error
	GetAppointmentByOwnerId(ownerID int32) ([]DTOs.AppointmentDTO, error)
}

type AppointmentServiceImpl struct {
	appointMappers        *mappers.AppointmentMappers
	appointmentRepository repository.AppointmentRepository
}

func NewAppointmentService(appointmentRepository repository.AppointmentRepository) AppointmentService {
	return &AppointmentServiceImpl{
		appointmentRepository: appointmentRepository,
		appointMappers:        mappers.NewAppointmentMappers(),
	}
}

func (as AppointmentServiceImpl) CreateAppointment(appointmentInsertDTO DTOs.AppointmentInsertDTO) error {
	appointmentParams := as.appointMappers.MapInsertDTOToInsertParams(appointmentInsertDTO)
	_, err := as.appointmentRepository.CreateAppointment(appointmentParams)
	if err != nil {
		return err
	}

	return nil
}

func (as AppointmentServiceImpl) GetAppointmentById(appointmentID int32) (*DTOs.AppointmentDTO, error) {
	appointment, err := as.appointmentRepository.GetAppointmentByID(appointmentID)
	if err != nil {
		return nil, err
	}

	appointmentDTO := as.appointMappers.MapSqlcEntityToDTO(*appointment)
	return &appointmentDTO, nil
}

func (as AppointmentServiceImpl) UpdateAppointment(appointmentUpdateDTO DTOs.AppointmentUpdateDTO) error {
	appointment, _ := as.appointmentRepository.GetAppointmentByID(appointmentUpdateDTO.Id)
	updateParams := as.appointMappers.MapUpdateDTOToUpdateParams(appointmentUpdateDTO, *appointment)

	if err := as.appointmentRepository.UpdateAppointment(updateParams); err != nil {
		return err
	}

	return nil
}

func (as AppointmentServiceImpl) DeleteAppointmentById(appointmentID int32) error {
	if err := as.appointmentRepository.DeleteAppointment(appointmentID); err != nil {
		return err
	}

	return nil
}

func (as AppointmentServiceImpl) RequestAnAppointment(AppointmentRequestInsertDTO DTOs.AppointmentRequestInsertDTO, ownerID int32) (*DTOs.AppointmentNamedDTO, error) {
	appointmentParams := as.appointMappers.MapRequestInsertDTOToInsertParams(AppointmentRequestInsertDTO, ownerID)
	appointment, err := as.appointmentRepository.RequestAppointment(appointmentParams)
	if err != nil {
		return nil, err
	}

	appointmentNamesDTO := as.appointMappers.MapSqlcEntityToNamedDTO(*appointment)
	return &appointmentNamesDTO, nil
}

func (as AppointmentServiceImpl) GetAppointmentByOwnerId(ownerID int32) ([]DTOs.AppointmentDTO, error) {
	appointments, err := as.appointmentRepository.GetAppointmentByOwnerID(ownerID)
	if err != nil {
		return nil, err
	}

	var appointmentsDTOs []DTOs.AppointmentDTO
	for _, appointment := range appointments {
		appointmentDTO := as.appointMappers.MapSqlcEntityToDTO(appointment)
		appointmentsDTOs = append(appointmentsDTOs, appointmentDTO)
	}

	return appointmentsDTOs, nil
}

func (as AppointmentServiceImpl) UpdateOwnerAppointment(appointmentUpdateDTO DTOs.AppointmentOwnerUpdateDTO, ownerID int32) error {
	params := as.appointMappers.MapUpdateOwnerDTOToUpdateOwnerParams(appointmentUpdateDTO)
	if err := as.appointmentRepository.UpdateOwnerAppointment(params); err != nil {
		return err
	}

	return nil
}

func (as AppointmentServiceImpl) CancelAppointmentById(appointmentID int32) error {
	appointment, _ := as.appointmentRepository.GetAppointmentByID(appointmentID)

	if appointment.Status == "completed" || appointment.Status == "no_show" {
		return errors.New("appointment not allowed to be canceled")
	} else if appointment.Status == "cancelled" {
		return errors.New("appointment already cancelled")
	} else {
		as.appointmentRepository.UpdateAppointmentStatus(appointmentID, "cancelled")
		return nil
	}

}
*/
