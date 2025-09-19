package command

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type PetCommandHandler interface {
	CreatePet(ctx context.Context, cmd CreatePetCommand) cqrs.CommandResult
	UpdatePet(ctx context.Context, cmd UpdatePetCommand) cqrs.CommandResult
	DeletePet(ctx context.Context, cmd DeletePetCommand) cqrs.CommandResult
	RestorePet(ctx context.Context, cmd RestorePetCommand) cqrs.CommandResult
	DeactivatePet(ctx context.Context, cmd DeactivatePetCommand) cqrs.CommandResult
}

type petCommandHandler struct {
	petRepository      repository.PetRepository
	customerRepository repository.CustomerRepository
}

func NewPetCommandHandler(petRepo repository.PetRepository, customerRepo repository.CustomerRepository) PetCommandHandler {
	return &petCommandHandler{
		petRepository:      petRepo,
		customerRepository: customerRepo,
	}
}
