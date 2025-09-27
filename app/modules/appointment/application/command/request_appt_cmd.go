package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/appointment"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"time"
)

type RequestApptByCustomerCommand struct {
	petID         valueobject.PetID
	customerID    valueobject.CustomerID
	requestedDate time.Time
	service       enum.ClinicService
	notes         *string
}

func NewRequestApptByCustomerCommand(
	petID uint, customerID uint, requestedDate time.Time, service string, notes *string,
) (*RequestApptByCustomerCommand, error) {
	parsedService, err := enum.ParseClinicService(service)
	if err != nil {
		return nil, err
	}

	return &RequestApptByCustomerCommand{
		petID:         valueobject.NewPetID(petID),
		requestedDate: requestedDate,
		customerID:    valueobject.NewCustomerID(customerID),
		service:       parsedService,
		notes:         notes,
	}, nil
}

func (h *apptCommandHandler) RequestAppointmentByCustomer(ctx context.Context, cmd RequestApptByCustomerCommand) cqrs.CommandResult {
	appointment := appointment.CreateCustomerRequest(
		cmd.petID, cmd.customerID, cmd.service, cmd.requestedDate, cmd.notes,
	)

	if err := appointment.ValidatePersistence(ctx); err != nil {
		return *cqrs.FailureResult("appointment validation failed", err)
	}

	if err := h.apptRepository.Save(ctx, &appointment); err != nil {
		return *cqrs.FailureResult("failed to save appointment", err)
	}

	// Event
	return *cqrs.SuccessResult(appointment.ID().String(), "appointment created successfully")
}
