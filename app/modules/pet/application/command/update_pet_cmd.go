package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type UpdatePetCommand struct {
	PetID      valueobject.PetID
	Name       *string
	Photo      *string
	Species    *enum.PetSpecies
	Breed      *string
	Age        *int
	BloodType  *string
	Tattoo     *string
	Gender     *enum.PetGender
	Color      *string
	Microchip  *string
	IsNeutered *bool
	CustomerID *valueobject.CustomerID
	IsActive   *bool
}

func (h *petCommandHandler) UpdatePet(ctx context.Context, cmd UpdatePetCommand) cqrs.CommandResult {
	pet, err := h.getPet(ctx, cmd.PetID, cmd.CustomerID)
	if err != nil {
		return cqrs.FailureResult("Error finding pet", err)
	}

	petUpdated := updatePet(pet, cmd)
	if _, err := h.petRepository.Save(ctx, petUpdated); err != nil {
		return cqrs.FailureResult("Error saving pet", err)
	}

	return cqrs.SuccessResult("Pet updated successfully")
}

func updatePet(p pet.Pet, cmd UpdatePetCommand) pet.Pet {
	petBuilder := pet.NewPetBuilder().WithID(p.ID())

	if cmd.Name != nil {
		petBuilder = petBuilder.WithName(*cmd.Name)
	} else {
		petBuilder = petBuilder.WithName(p.Name())
	}

	if cmd.Photo != nil {
		petBuilder = petBuilder.WithPhoto(cmd.Photo)
	} else {
		petBuilder = petBuilder.WithPhoto(p.Photo())
	}

	if cmd.Species != nil {
		petBuilder = petBuilder.WithSpecies(*cmd.Species)
	} else {
		petBuilder = petBuilder.WithSpecies(p.Species())
	}

	if cmd.Breed != nil {
		petBuilder = petBuilder.WithBreed(cmd.Breed)
	} else {
		petBuilder = petBuilder.WithBreed(p.Breed())
	}

	if cmd.Age != nil {
		petBuilder = petBuilder.WithAge(cmd.Age)
	} else {
		petBuilder = petBuilder.WithAge(p.Age())
	}

	if cmd.Gender != nil {
		petBuilder = petBuilder.WithGender(*cmd.Gender)
	} else {
		petBuilder = petBuilder.WithGender(p.Gender())
	}

	if cmd.Color != nil {
		petBuilder = petBuilder.WithColor(cmd.Color)
	} else {
		petBuilder = petBuilder.WithColor(p.Color())
	}

	if cmd.Microchip != nil {
		petBuilder = petBuilder.WithMicrochip(cmd.Microchip)
	} else {
		petBuilder = petBuilder.WithMicrochip(p.Microchip())
	}

	if cmd.IsNeutered != nil {
		petBuilder = petBuilder.WithIsNeutered(cmd.IsNeutered)
	} else {
		petBuilder = petBuilder.WithIsNeutered(p.IsNeutered())
	}

	if cmd.CustomerID != nil {
		petBuilder = petBuilder.WithCustomerID(*cmd.CustomerID)
	} else {
		petBuilder = petBuilder.WithCustomerID(p.CustomerID())
	}

	if cmd.IsActive != nil {
		petBuilder = petBuilder.WithIsActive(*cmd.IsActive)
	} else {
		petBuilder = petBuilder.WithIsActive(p.IsActive())
	}

	return *petBuilder.Build()
}
