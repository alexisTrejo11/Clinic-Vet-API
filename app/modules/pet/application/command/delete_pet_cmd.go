package command

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type DeletePetCommand struct {
	petID        valueobject.PetID
	customer     *valueobject.CustomerID
	isSoftDelete bool
}

func NewDeletePetCommand(petIDUint uint, customerIDUint *uint, isSoftDelete bool) *DeletePetCommand {
	var customerID *valueobject.CustomerID
	if customerIDUint != nil {
		customerIDVal := valueobject.NewCustomerID(*customerIDUint)
		customerID = &customerIDVal
	}

	return &DeletePetCommand{
		petID:        valueobject.NewPetID(petIDUint),
		customer:     customerID,
		isSoftDelete: isSoftDelete,
	}
}

func (h *petCommandHandler) DeletePet(ctx context.Context, cmd DeletePetCommand) cqrs.CommandResult {
	pet, err := h.getPet(ctx, cmd.petID, cmd.customer)
	if err != nil {
		return *cqrs.FailureResult("Error Getting Pet", err)
	}

	if err := h.petRepository.Delete(ctx, pet.ID(), false); err != nil {
		return *cqrs.FailureResult("Error Deleting Pet", err)
	}

	return *cqrs.SuccessResult("Pet Deleted Successfully")
}
