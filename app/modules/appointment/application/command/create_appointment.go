package command

import (
	"context"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type CreateAppointmentCommand struct {
	ctx      context.Context
	ownerID  valueobject.OwnerID
	petID    valueobject.PetID
	vetID    *valueobject.VetID
	service  enum.ClinicService
	datetime time.Time
	status   *enum.AppointmentStatus
	reason   string
	notes    *string
}

func NewCreateAppointCmd(
	ctx context.Context,
	ownerIDInt,
	petIDInt int,
	vetIDInt *int,
	service string,
	dateTime time.Time,
	status string,
	notes *string,
) (*CreateAppointmentCommand, error) {
	ownerID, err := valueobject.NewOwnerID(ownerIDInt)
	if err != nil {
		return nil, err
	}
	petID, err := valueobject.NewPetID(petIDInt)
	if err != nil {
		return nil, err
	}

	var vetID *valueobject.VetID
	if vetIDInt != nil {
		vetIDObj, err := valueobject.NewVetID(*vetIDInt)
		vetID = &vetIDObj

		if err != nil {
			return nil, err
		}
	}

	cliniService, err := enum.NewClinicService(service)
	if err != nil {
		return nil, err
	}

	appointmentStatus, err := enum.NewAppointmentStatus(status)
	if err != nil {
		return nil, err
	}

	return &CreateAppointmentCommand{
		ownerID:  ownerID,
		petID:    petID,
		vetID:    vetID,
		datetime: dateTime,
		notes:    notes,
		status:   &appointmentStatus,
		service:  cliniService,
	}, nil
}

type CreateAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewCreateAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &CreateAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *CreateAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(CreateAppointmentCommand)

	appointment, err := h.commandToDomain(command)
	if err != nil {
		return cqrs.FailureResult("failed to create appointment domain", err)
	}

	if err := appointment.ValidateRequestSchedule(); err != nil {
		return cqrs.FailureResult("validation for schedule failed", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, appointment); err != nil {
		return cqrs.FailureResult("failed to save appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment created successfully")
}

func (h *CreateAppointmentHandler) commandToDomain(command CreateAppointmentCommand) (*entity.Appointment, error) {
	appointmentID, err := valueobject.NewAppointmentID(0)
	if err != nil {
		return nil, err
	}

	status := enum.StatusPending
	if command.status != nil {
		status = *command.status
	}

	now := time.Now()
	appointment := entity.NewAppointment(
		appointmentID,
		command.petID,
		command.ownerID,
		command.vetID,
		command.service,
		command.datetime,
		status,
		now,
		now,
	)

	return appointment, nil
}
