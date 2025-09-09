package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type UpdateApptCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.VetID
	status        *enum.AppointmentStatus
	reason        *string
	notes         *string
	service       *enum.ClinicService
}

func NewUpdateApptCommand(ctx context.Context, appointIDInt uint, vetIDInt *uint, status string, reason, notes *string, service *enum.ClinicService) *UpdateApptCommand {
	var vetID *valueobject.VetID
	if vetIDInt != nil {
		vetIDObj := valueobject.NewVetID(*vetIDInt)
		vetID = &vetIDObj
	}

	return &UpdateApptCommand{
		ctx:           ctx,
		appointmentID: valueobject.NewAppointmentID(appointIDInt),
		vetID:         vetID,
		service:       service,
		reason:        reason,
		notes:         notes,
	}
}

type UpdateApptHandler struct {
	apptRepository repository.AppointmentRepository
}

func NewUpdateApptHandler(apptRepository repository.AppointmentRepository) cqrs.CommandHandler {
	return &UpdateApptHandler{
		apptRepository: apptRepository,
	}
}

func (h *UpdateApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, ok := cmd.(*UpdateApptCommand)
	if !ok {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.apptRepository.GetByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Update(command.notes, command.vetID, command.service, command.reason); err != nil {
		return cqrs.FailureResult(ErrUpdateApptFailed, err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult(ErrSaveApptFailed, err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), SuccessApptUpdated)
}
