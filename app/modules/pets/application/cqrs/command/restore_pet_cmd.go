package command

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type RestorePetCommand struct {
	petID valueobject.PetID
}

func NewRestorePetCommand(petID uint) *RestorePetCommand {
	return &RestorePetCommand{
		petID: valueobject.NewPetID(petID),
	}
}

func (h *petCommandHandler) RestorePet(ctx context.Context, cmd RestorePetCommand) cqrs.CommandResult {
	pet, err := h.getPet(ctx, cmd.petID, nil)
	if err != nil {
		return *cqrs.FailureResult("Error Getting Pet", err)
	}

	if err := pet.Restore(); err != nil {
		return *cqrs.FailureResult("Error Restoring Pet", err)
	}

	if _, err := h.petRepository.Save(ctx, pet); err != nil {
		return *cqrs.FailureResult("Error Updating Pet", err)

	}

	return *cqrs.SuccessResult(pet.ID().String(), "Pet Restored Successfully")
}
