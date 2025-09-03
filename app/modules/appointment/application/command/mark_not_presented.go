package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type NotAttendAppointmentCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
}

func NewNotAttendAppointmentCommand(ctx context.Context, id int) (*NotAttendAppointmentCommand, error) {
	appointmentID, err := valueobject.NewAppointmentID(id)
	if err != nil {
		return nil, err
	}
	return &NotAttendAppointmentCommand{
		ctx:           ctx,
		appointmentID: appointmentID,
	}, nil
}

type NotAttendAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewNotAttendAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &NotAttendAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *NotAttendAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(NotAttendAppointmentCommand)
	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.NotPresented(); err != nil {
		return cqrs.FailureResult("failed to mark appointment as not presented", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment marked as not presented successfully")
}
