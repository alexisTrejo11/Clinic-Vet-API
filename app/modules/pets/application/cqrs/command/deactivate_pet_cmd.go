package command

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type DeactivatePetCommand struct {
	petID      valueobject.PetID
	customerID *valueobject.CustomerID
}

func NewDeactivatePetCommand(petID uint, optCustomerID *uint) *DeactivatePetCommand {
	var customerID *valueobject.CustomerID
	if optCustomerID != nil {
		custID := valueobject.NewCustomerID(*optCustomerID)
		customerID = &custID
	}

	return &DeactivatePetCommand{
		petID:      valueobject.NewPetID(petID),
		customerID: customerID,
	}
}

func (h *petCommandHandler) DeactivatePet(ctx context.Context, cmd DeactivatePetCommand) cqrs.CommandResult {
	pet, err := h.getPet(ctx, cmd.petID, cmd.customerID)
	if err != nil {
		return *cqrs.FailureResult("Error Getting Pet", err)
	}

	if err := pet.Deactivate(); err != nil {
		return *cqrs.FailureResult("Error Deactivating Pet", err)
	}

	if _, err := h.petRepository.Save(ctx, pet); err != nil {
		return *cqrs.FailureResult("Error Updating Pet", err)
	}

	return *cqrs.SuccessResult(pet.ID().String(), "Pet Deactivated Successfully")
}
