package command

import (
	"context"
	"errors"
	"time"

	appt "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CreateApptCommand struct {
	ctx      context.Context
	ownerID  valueobject.OwnerID
	petID    valueobject.PetID
	vetID    *valueobject.VetID
	service  enum.ClinicService
	datetime time.Time
	status   *enum.AppointmentStatus
	reason   enum.VisitReason
	notes    *string
}

func NewCreateApptCommand(
	ctx context.Context,
	ownerIDInt,
	petIDInt uint,
	vetIDInt *uint,
	service enum.ClinicService,
	dateTime time.Time,
	status enum.AppointmentStatus,
	reason enum.VisitReason,
	notes *string,
) *CreateApptCommand {
	var vetID *valueobject.VetID
	if vetIDInt != nil {
		vetIDObj := valueobject.NewVetID(*vetIDInt)
		vetID = &vetIDObj
	}

	return &CreateApptCommand{
		ownerID:  valueobject.NewOwnerID(ownerIDInt),
		petID:    valueobject.NewPetID(petIDInt),
		vetID:    vetID,
		datetime: dateTime,
		notes:    notes,
		reason:   reason,
		status:   &status,
		service:  service,
	}
}

type CreateApptHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewCreateApptHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &CreateApptHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *CreateApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, valid := cmd.(CreateApptCommand)
	if !valid {
		return cqrs.FailureResult("invalid command type", errors.New("expected CreateApptCommand"))
	}

	appointment, err := h.commandToDomain(command)
	if err != nil {
		return cqrs.FailureResult("failed to create appointment domain", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, appointment); err != nil {
		return cqrs.FailureResult("failed to save appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment created successfully")
}

func (h *CreateApptHandler) commandToDomain(command CreateApptCommand) (*appt.Appointment, error) {
	status := enum.AppointmentStatusPending
	if command.status != nil {
		status = *command.status
	}

	appointmentEntity, err := appt.CreateAppointment(
		command.petID,
		command.ownerID,
		appt.WithVetID(command.vetID),
		appt.WithService(command.service),
		appt.WithScheduledDate(command.datetime),
		appt.WithReason(command.reason),
		appt.WithNotes(command.notes),
		appt.WithStatus(status),
	)
	if err != nil {
		return nil, apperror.MappingError([]string{err.Error()}, "constructor", "domain", "Appointment")
	}
	return appointmentEntity, nil
}
