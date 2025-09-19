package cqrs

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/cqrs/command"
	"clinic-vet-api/app/modules/pets/application/cqrs/query"
)

type PetServiceBus interface {
	command.PetCommandHandler
	query.PetQueryHandler
}

type petServiceBus struct {
	command.PetCommandHandler
	query.PetQueryHandler
}

func NewPetServiceBus(petRepo repository.PetRepository, customerRepo repository.CustomerRepository) PetServiceBus {
	cmdHandler := command.NewPetCommandHandler(petRepo, customerRepo)
	qryHandler := query.NewPetQueryHandler(petRepo, customerRepo)

	return &petServiceBus{
		PetCommandHandler: cmdHandler,
		PetQueryHandler:   qryHandler,
	}
}
