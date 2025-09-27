package command

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"errors"
)

type UpdateApptCommand struct {
	appointmentID valueobject.AppointmentID
	status        *enum.AppointmentStatus
	notes         *string
	service       *enum.ClinicService
}

func NewUpdateApptCommand(
	appointIDInt uint, status *string, notes *string, service *string,
) *UpdateApptCommand {
	var statusEnum *enum.AppointmentStatus
	if status != nil {
		parsedStatus := enum.AppointmentStatus(*status)
		statusEnum = &parsedStatus
	}

	var serviceEnum *enum.ClinicService
	if service != nil {
		parsedService := enum.ClinicService(*service)
		serviceEnum = &parsedService
	}

	return &UpdateApptCommand{
		appointmentID: valueobject.NewAppointmentID(appointIDInt),
		service:       serviceEnum,
		status:        statusEnum,
		notes:         notes,
	}
}

func (h *apptCommandHandler) UpdateAppointment(ctx context.Context, cmd UpdateApptCommand) cqrs.CommandResult {
	if err := cmd.Validate(); err != nil {
		return *cqrs.FailureResult(ErrInvalidCommand, err)
	}

	appointment, err := h.apptRepository.FindByID(ctx, cmd.appointmentID)
	if err != nil {
		return *cqrs.FailureResult(ErrApptNotFound, err)
	}

	if err := appointment.Update(ctx, cmd.notes, cmd.service); err != nil {
		return *cqrs.FailureResult(ErrUpdateApptFailed, err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult(ErrSaveApptFailed, err)
	}

	return *cqrs.SuccessResult(appointment.ID().String(), SuccessApptUpdated)
}

func (c *UpdateApptCommand) Validate() error {
	if c.appointmentID.IsZero() {
		return errors.New("appointment ID required")
	}

	if c.status != nil {
		_, err := enum.ParseAppointmentStatus(c.status.String())
		if err != nil {
			return err
		}
	}

	if c.service != nil {
		_, err := enum.ParseClinicService(c.service.String())
		if err != nil {
			return err
		}
	}

	return nil
}
