package command

import (
	"clinic-vet-api/app/core/domain/entity/pet"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
	"fmt"
	"strings"
)

type UpdatePetCommand struct {
	PetID              valueobject.PetID       `json:"pet_id"`
	Name               *string                 `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Photo              *string                 `json:"photo,omitempty" validate:"omitempty,url"`
	Species            *string                 `json:"species,omitempty" validate:"omitempty,min=2,max=50"`
	Breed              *string                 `json:"breed,omitempty" validate:"omitempty,min=2,max=50"`
	Age                *int                    `json:"age,omitempty" validate:"omitempty,min=0"`
	Gender             *string                 `json:"gender,omitempty" validate:"omitempty,oneof=Male Female Unknown"`
	Weight             *float64                `json:"weight,omitempty" validate:"omitempty,gt=0,lte=1000"`
	Color              *string                 `json:"color,omitempty" validate:"omitempty,min=2,max=50"`
	Microchip          *string                 `json:"microchip,omitempty" validate:"omitempty,len=15,numeric"`
	IsNeutered         *bool                   `json:"is_neutered,omitempty"`
	CustomerID         *valueobject.CustomerID `json:"customer_id,omitempty" validate:"omitempty,gt=0"`
	Allergies          *string                 `json:"allergies,omitempty" validate:"omitempty,max=500"`
	CurrentMedications *string                 `json:"current_medications,omitempty" validate:"omitempty,max=500"`
	SpecialNeeds       *string                 `json:"special_needs,omitempty" validate:"omitempty,max=500"`
	IsActive           *bool                   `json:"is_active,omitempty"`
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

	if cmd.Weight != nil {
		options = append(options, pet.UpdateWeight(cmd.Weight))
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

	if cmd.Allergies != nil {
		options = append(options, pet.UpdateAllergies(cmd.Allergies))
	}

	if cmd.CurrentMedications != nil {
		options = append(options, pet.UpdateMedications(cmd.CurrentMedications))
	}

	if cmd.SpecialNeeds != nil {
		options = append(options, pet.UpdateSpecialNeeds(cmd.SpecialNeeds))
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
