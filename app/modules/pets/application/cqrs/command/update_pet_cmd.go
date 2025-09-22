package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"fmt"
	"strings"
)

type UpdatePetCommand struct {
	PetID      valueobject.PetID
	Name       *string
	Photo      *string
	Species    *string
	Breed      *string
	Age        *int
	BloodType  *string
	Tattoo     *string
	Gender     *string
	Color      *string
	Microchip  *string
	IsNeutered *bool
	CustomerID *valueobject.CustomerID
	IsActive   *bool
}

func (h *petCommandHandler) UpdatePet(ctx context.Context, cmd UpdatePetCommand) cqrs.CommandResult {
	pet, err := h.getPet(ctx, cmd.PetID, cmd.CustomerID)
	if err != nil {
		return *cqrs.FailureResult("Error finding pet", err)
	}

	if err := appyUpdates(ctx, &pet, cmd); err != nil {
		return *cqrs.FailureResult("Error updating pet", err)
	}

	if _, err := h.petRepository.Save(ctx, pet); err != nil {
		return *cqrs.FailureResult("Error saving pet", err)
	}

	return *cqrs.SuccessResult(pet.ID().String(), "Pet updated successfully")
}

func appyUpdates(ctx context.Context, p *pet.Pet, cmd UpdatePetCommand) error {
	var options []pet.PetUpdateOption
	var errors []error

	if cmd.Name != nil {
		options = append(options, pet.UpdateName(*cmd.Name))
	}

	if cmd.Photo != nil {
		options = append(options, pet.UpdatePhoto(cmd.Photo))
	}

	if cmd.Species != nil {
		options = append(options, pet.UpdateSpecies(*cmd.Species))
	}

	if cmd.Breed != nil {
		options = append(options, pet.UpdateBreed(cmd.Breed))
	}

	if cmd.Age != nil {
		options = append(options, pet.UpdateAge(cmd.Age))
	}

	if cmd.Gender != nil {
		petGender, err := enum.ParsePetGender(*cmd.Gender)
		if err != nil {
			errors = append(errors, fmt.Errorf("gender: %w", err))
		} else {
			options = append(options, pet.UpdateGender(&petGender))
		}
	}

	if cmd.Color != nil {
		options = append(options, pet.UpdateColor(cmd.Color))
	}

	if cmd.Microchip != nil {
		options = append(options, pet.UpdateMicrochip(cmd.Microchip))
	}

	if cmd.IsNeutered != nil {
		options = append(options, pet.UpdateIsNeutered(cmd.IsNeutered))
	}

	if cmd.CustomerID != nil {
		options = append(options, pet.UpdateCustomerID(*cmd.CustomerID))
	}

	if cmd.IsActive != nil {
		options = append(options, pet.UpdateIsActive(*cmd.IsActive))
	}

	if len(errors) > 0 {
		var errorMsgs []string
		for _, err := range errors {
			errorMsgs = append(errorMsgs, err.Error())
		}
		return fmt.Errorf("parsing errors: %v", strings.Join(errorMsgs, "; "))
	}

	if err := p.Update(ctx, options...); err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	return nil
}
