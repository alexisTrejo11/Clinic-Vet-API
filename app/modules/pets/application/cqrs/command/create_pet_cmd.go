package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
	"context"
	"fmt"
)

type CreatePetCommand struct {
	Name       string
	Photo      *string
	Species    string
	Breed      *string
	Age        *int
	Gender     *string
	Color      *string
	Microchip  *string
	Tattoo     *string
	BloodType  *string
	IsNeutered *bool
	CustomerID valueobject.CustomerID
	IsActive   bool
}

func CustomerNotFoundError(customerID valueobject.CustomerID) error {
	return apperror.EntityNotFoundValidationError("customer", "id", customerID.String())
}

func (h *petCommandHandler) CreatePet(ctx context.Context, cmd CreatePetCommand) cqrs.CommandResult {
	if err := h.validateCustomer(ctx, cmd.CustomerID); err != nil {
		return *cqrs.FailureResult("Customer validation failed", err)
	}

	newPet, err := cmd.ToEntity()
	if err != nil {
		return *cqrs.FailureResult("Failed to create pet entity", err)
	}

	petCreated, err := h.petRepository.Save(ctx, newPet)
	if err != nil {
		return *cqrs.FailureResult("Failed to save new pet", err)
	}

	return *cqrs.SuccessResult(petCreated.ID().String(), "Pet created successfully")
}

func (h *petCommandHandler) validateCustomer(ctx context.Context, customerID valueobject.CustomerID) error {
	if exists, err := h.customerRepository.ExistsByID(ctx, customerID); err != nil {
		return err
	} else if !exists {
		return CustomerNotFoundError(customerID)
	}
	return nil
}

func (h *petCommandHandler) getPet(ctx context.Context, petID valueobject.PetID, customerID *valueobject.CustomerID) (pet.Pet, error) {
	if customerID != nil {
		return h.petRepository.FindByIDAndCustomerID(ctx, petID, *customerID)
	}
	return h.petRepository.FindByID(ctx, petID)
}

func (cmd *CreatePetCommand) ToEntity() (pet.Pet, error) {
	opts := []pet.PetOption{
		pet.WithName(cmd.Name),
		pet.WithSpecies(cmd.Species),
		pet.WithIsActive(cmd.IsActive),
	}

	// Agregar options para campos opcionales
	if cmd.Photo != nil {
		opts = append(opts, pet.WithPhoto(cmd.Photo))
	}

	if cmd.Breed != nil {
		opts = append(opts, pet.WithBreed(cmd.Breed))
	}

	if cmd.Age != nil {
		opts = append(opts, pet.WithAge(cmd.Age))
	}

	if cmd.Gender != nil {
		gender, err := enum.ParsePetGender(*cmd.Gender)
		if err != nil {
			return pet.Pet{}, fmt.Errorf("invalid gender: %w", err)
		}
		opts = append(opts, pet.WithGender(&gender))
	}

	if cmd.Tattoo != nil {
		opts = append(opts, pet.WithTattoo(cmd.Tattoo))
	}

	if cmd.BloodType != nil {
		opts = append(opts, pet.WithBloodType(cmd.BloodType))
	}

	if cmd.Color != nil {
		opts = append(opts, pet.WithColor(cmd.Color))
	}

	if cmd.Microchip != nil {
		opts = append(opts, pet.WithMicrochip(cmd.Microchip))
	}

	if cmd.IsNeutered != nil {
		opts = append(opts, pet.WithIsNeutered(cmd.IsNeutered))
	}

	petEntity, err := pet.CreatePet(cmd.CustomerID, opts...)
	if err != nil {
		return pet.Pet{}, err
	}

	return *petEntity, nil
}
