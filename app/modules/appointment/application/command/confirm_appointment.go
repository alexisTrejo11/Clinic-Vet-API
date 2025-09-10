package command

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type ConfirmApptCommand struct {
	id         valueobject.AppointmentID
	employeeID valueobject.EmployeeID
	ctx        context.Context
}

func NewConfirmAppointmentCommand(ctx context.Context, appointIDInt, vetIDInt uint) *ConfirmApptCommand {
	return &ConfirmApptCommand{
		ctx:        ctx,
		id:         valueobject.NewAppointmentID(appointIDInt),
		employeeID: valueobject.NewEmployeeID(vetIDInt),
	}
}

type ConfirmApptHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewConfirmApptHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &ConfirmApptHandler{appointmentRepo: appointmentRepo}
}

func (h *ConfirmApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, valid := cmd.(ConfirmApptCommand)
	if !valid {
		return cqrs.FailureResult("invalid command type", errors.New("invalid command type"))
	}

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.id)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	if err := appointment.Confirm(command.employeeID); err != nil {
		return cqrs.FailureResult("failed to confirm appointment", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save confirmed appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment confirmed successfully")
}
