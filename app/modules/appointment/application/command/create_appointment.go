package command

import (
	"context"
	"time"

	appt "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
)

type CreateAppointmentCommand struct {
	ctx         context.Context
	ownerID     valueobject.OwnerID
	petID       valueobject.PetID
	vetID       *valueobject.VetID
	service     enum.ClinicService
	datetime    time.Time
	status      *enum.AppointmentStatus
	reason      enum.VisitReason
	notes       *string
	RequestByID valueobject.UserID
}

func NewCreateAppointCmd(
	ctx context.Context,
	ownerIDInt,
	petIDInt int,
	vetIDInt *int,
	service string,
	dateTime time.Time,
	status string,
	reason string,
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

	cliniService, err := enum.ParseClinicService(service)
	if err != nil {
		return nil, err
	}

	appointmentStatus, err := enum.ParseAppointmentStatus(status)
	if err != nil {
		return nil, err
	}

	visitReason, err := enum.ParseVisitReason(reason)
	if err != nil {
		return nil, err
	}

	return &CreateAppointmentCommand{
		ownerID:  ownerID,
		petID:    petID,
		vetID:    vetID,
		datetime: dateTime,
		notes:    notes,
		reason:   visitReason,
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

	if err := h.appointmentRepo.Save(command.ctx, appointment); err != nil {
		return cqrs.FailureResult("failed to save appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment created successfully")
}

func (h *CreateAppointmentHandler) commandToDomain(command CreateAppointmentCommand) (*appt.Appointment, error) {
	appointmentID, err := valueobject.NewAppointmentID(0)
	if err != nil {
		return nil, err
	}

	status := enum.AppointmentStatusPending
	if command.status != nil {
		status = *command.status
	}

	appointmentEntity, err := appt.NewAppointment(
		appointmentID,
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
