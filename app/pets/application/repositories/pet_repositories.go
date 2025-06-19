package repositories

import "example.com/at/backend/api-vet/app/container/sqlc"

type PetRepository interface {
	CreatePet(args sqlc.CreatePetParams) (*sqlc.Pet, error)
	GetPetById(petId int32) (*sqlc.Pet, error)
	GetPetByOwnerID(petId int32) ([]sqlc.Pet, error)
	UpdatePetById(params sqlc.UpdatePetParams) error
	DeletePetById(petId int32) error
}
