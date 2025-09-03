package command

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type UpdateAppointmentCommand struct {
	ctx           context.Context
	appointmentID valueobject.AppointmentID
	vetID         *valueobject.VetID
	status        *enum.AppointmentStatus
	reason        *string
	notes         *string
	service       *enum.ClinicService
}

func NewUpdateAppointmentCommand(ctx context.Context, appointIDInt int, vetIDInt *int, status string, reason, notes, service *string) (*UpdateAppointmentCommand, error) {
	appointmentID, err := valueobject.NewAppointmentID(appointIDInt)
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

	var clinicService *enum.ClinicService
	if service != nil {
		service, err := enum.NewClinicService(*service)
		if err != nil {
			return nil, err
		}
		clinicService = &service
	}

	return &UpdateAppointmentCommand{
		ctx:           ctx,
		appointmentID: appointmentID,
		vetID:         vetID,
		service:       clinicService,
		reason:        reason,
		notes:         notes,
	}, nil
}

type UpdateAppointmentHandler struct {
	appointmentRepo repository.AppointmentRepository
}

func NewUpdateAppointmentHandler(appointmentRepo repository.AppointmentRepository) cqrs.CommandHandler {
	return &UpdateAppointmentHandler{
		appointmentRepo: appointmentRepo,
	}
}

func (h *UpdateAppointmentHandler) Handle(cmd cqrs.Command) cqrs.CommandResult {
	command := cmd.(UpdateAppointmentCommand)

	appointment, err := h.appointmentRepo.GetByID(command.ctx, command.appointmentID)
	if err != nil {
		return cqrs.FailureResult("appointment not found", err)
	}

	appointment.Update(command.notes, command.vetID, command.service, command.reason)

	if err := appointment.ValidateFields(); err != nil {
		return cqrs.FailureResult("appointment validation failed", err)
	}

	if err := h.appointmentRepo.Save(command.ctx, &appointment); err != nil {
		return cqrs.FailureResult("failed to update appointment", err)
	}

	return cqrs.SuccessResult(appointment.GetID().String(), "appointment updated successfully")
}
