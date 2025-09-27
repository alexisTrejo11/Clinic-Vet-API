package command

import (
	"context"
	"errors"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
)

type DewormDeleteCommand struct {
	ID    valueobject.DewormID
	PetID *valueobject.PetID
}

func NewDewormDeleteCommand(id uint, petID *uint) DewormDeleteCommand {
	var petIDPtr *valueobject.PetID
	if petID != nil {
		pid := valueobject.NewPetID(*petID)
		petIDPtr = &pid
	}

	return DewormDeleteCommand{ID: valueobject.NewDewormID(id), PetID: petIDPtr}
}

func (h *DewormCommandHandler) HandleDelete(ctx context.Context, cmd DewormDeleteCommand) cqrs.CommandResult {
	if err := h.validateDewormExistence(ctx, cmd.ID, cmd.PetID); err != nil {
		return *cqrs.FailureResult("entity validation error", err)
	}

	if err := h.dewormRepo.Delete(ctx, cmd.ID, true); err != nil {
		return *cqrs.FailureResult("failed to delete deworming record", err)
	}

	return *cqrs.SuccessResult(cmd.ID.String(), "deworming record deleted successfully")
}

func (cmd *DewormDeleteCommand) validate() error {
	if cmd.ID.IsZero() {
		return errors.New("invalid deworm ID")
	}
	if cmd.PetID != nil && cmd.PetID.IsZero() {
		return errors.New("invalid pet ID")
	}
	return nil
}
