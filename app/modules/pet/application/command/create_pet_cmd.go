package command

import (
	"clinic-vet-api/app/modules/core/domain/entity/pet"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
	"context"
)

type CreatePetCommand struct {
	Name       string
	Photo      *string
	Species    enum.PetSpecies
	Breed      *string
	Age        *int
	Gender     enum.PetGender
	Color      *string
	Microchip  *string
	Tattoo     *string
	BloodType  *string
	IsNeutered *bool
	CustomerID valueobject.CustomerID
	IsActive   bool
}

func (h *petCommandHandler) CreatePet(ctx context.Context, cmd CreatePetCommand) cqrs.CommandResult {
	if err := h.validateCustomer(ctx, cmd.CustomerID); err != nil {
		return cqrs.FailureResult("Customer validation failed", err)
	}

	newPet := cmd.ToEntity()
	petCreated, err := h.petRepository.Save(ctx, newPet)
	if err != nil {
		return cqrs.FailureResult("Failed to save new pet", err)
	}

	return cqrs.SuccessCreateResult(petCreated.ID().String(), "Pet created successfully")
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

func (cmd *CreatePetCommand) ToEntity() pet.Pet {
	return *pet.NewPetBuilder().
		WithName(cmd.Name).
		WithSpecies(cmd.Species).
		WithIsActive(cmd.IsActive).
		WithPhoto(cmd.Photo).
		WithBreed(cmd.Breed).
		WithAge(cmd.Age).
		WithGender(cmd.Gender).
		WithCustomerID(cmd.CustomerID).
		WithTattoo(cmd.Tattoo).
		WithBloodType(cmd.BloodType).
		WithColor(cmd.Color).
		WithMicrochip(cmd.Microchip).
		WithIsNeutered(cmd.IsNeutered).
		Build()
}

func (cmd *CreatePetCommand) Validate() error {
	if cmd.Name == "" {
		return apperror.ValidationError("pet name is required")
	}

	if cmd.CustomerID.IsZero() {
		return apperror.ValidationError("customer_id is required")
	}

	if err := cmd.Species.Validate(); err != nil {
		return err
	}

	if err := cmd.Gender.Validate(); err != nil {
		return err
	}

	return nil
}

func CustomerNotFoundError(customerID valueobject.CustomerID) error {
	return apperror.EntityNotFoundValidationError("customer", "id", customerID.String())
}
