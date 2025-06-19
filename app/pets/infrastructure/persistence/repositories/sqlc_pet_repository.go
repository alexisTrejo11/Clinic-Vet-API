package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type petRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewPetRepository(queries *sqlc.Queries) PetRepository {
	return &petRepositoryImpl{
		queries: queries,
	}
}

func (r *petRepositoryImpl) Create(args sqlc.CreatePetParams) (*sqlc.Pet, error) {
	pet, err := r.queries.CreatePet(context.Background(), args)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (r *petRepositoryImpl) GetById(petId int32) (*sqlc.Pet, error) {
	pet, err := r.queries.GetPetByID(context.Background(), petId)
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (r *petRepositoryImpl) GetByOwnerID(petId int32) ([]sqlc.Pet, error) {
	pets, err := r.queries.ListPetsByOwnerByID(context.Background(), petId)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (r *petRepositoryImpl) Update(params sqlc.UpdatePetParams) error {
	err := r.queries.UpdatePet(context.Background(), params)
	if err != nil {
		return err
	}
	return nil
}

func (r *petRepositoryImpl) Delete(petId int32) error {
	err := r.queries.DeletePet(context.Background(), petId)
	if err != nil {
		return err
	}
	return nil
}
