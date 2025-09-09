package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/appointment"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type CompleteApptCommand struct {
	id    valueobject.AppointmentID
	vetID *valueobject.VetID
	notes *string
	ctx   context.Context
}

func NewCompleteApptCommand(ctx context.Context, id uint, vetIDInt *uint, notes string) *CompleteApptCommand {
	cmd := &CompleteApptCommand{
		id:    valueobject.NewAppointmentID(id),
		notes: &notes,
		ctx:   ctx,
	}

	if vetIDInt != nil {
		vetID := valueobject.NewVetID(*vetIDInt)
		cmd.vetID = &vetID
	}
	return cmd
}

type CompleteApptHandler struct {
	apptRepository repository.AppointmentRepository
}

func NewCompleteApptHandler(apptRepository repository.AppointmentRepository) cqrs.CommandHandler {
	return &CompleteApptHandler{
		apptRepository: apptRepository,
	}
}

func (h *CompleteApptHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command, valid := cmd.(CompleteApptCommand)
	if !valid {
		return cqrs.FailureResult(ErrInvalidCommandType, nil)
	}

	appointment, err := h.getAppt(command)
	if err != nil {
		return cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Complete(); err != nil {
		return cqrs.FailureResult("failed to complete appointment", err)
	}

	if err := h.apptRepository.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to save completed appointment", err)
	}

	return cqrs.SuccessResult(appointment.ID().String(), "appointment completed successfully")
}

func (h *CompleteApptHandler) getAppt(command CompleteApptCommand) (appointment.Appointment, error) {
	if command.vetID != nil {
		appointment, err := h.apptRepository.GetByIDAndVetID(command.ctx, command.id, *command.vetID)
		if err != nil {
			return appointment, err
		}
		return appointment, nil
	}

	appointment, err := h.apptRepository.GetByID(command.ctx, command.id)
	if err != nil {
		return appointment, err
	}
	return appointment, nil
}
