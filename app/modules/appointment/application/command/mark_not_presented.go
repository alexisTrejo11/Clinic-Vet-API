package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type NotAttendApptCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.EmployeeID
}

func NewNotAttendApptCommand(ctx context.Context, id uint, vetIDUint *uint) *NotAttendApptCommand {
	var vetID *valueobject.EmployeeID
	if vetIDUint != nil {
		vetIDVal := valueobject.NewEmployeeID(*vetIDUint)
		vetID = &vetIDVal
	}

	return &NotAttendApptCommand{
		ctx:           ctx,
		appointmentID: valueobject.NewAppointmentID(id),
		vetID:         vetID,
	}
}

type NotAttendApptHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewNotAttendApptHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &NotAttendApptHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *NotAttendApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(*NotAttendApptCommand)
	if !ok {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.getAppointment(command)
	if err != nil {
		return cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.MarkAsNotPresented(); err != nil {
		return cqrs.FailureResult(ErrMarkAsNotPresentedFailed, err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult(ErrSaveApptFailed, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessMarkedAsNotPresented)
}

func (h *NotAttendApptHandler) getAppointment(command *NotAttendApptCommand) (appointment.Appointment, error) {
	if command.vetID != nil {
		return h.appointmentRepo.FindByIDAndEmployeeID(command.ctx, command.appointmentID, *command.vetID)
	}

	return h.appointmentRepo.FindByID(command.ctx, command.appointmentID)
}
