package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type PetRepositoryInterface interface {
	CreatePet(ctx context.Context, args sqlc.CreatePetParams) (sqlc.Pet, error)
	GetPetById(ctx context.Context, petId int32) (sqlc.Pet, error)
	UpdatePetById(ctx context.Context, params sqlc.UpdatePetParams) error
	DeletePetById(ctx context.Context, petId int32) error
}

type PetRepository struct {
	queries *sqlc.Queries
}

func NewPetRepository(queries *sqlc.Queries) *PetRepository {
	return &PetRepository{
		queries: queries,
	}
}

func (r *PetRepository) CreatePet(ctx context.Context, args sqlc.CreatePetParams) (sqlc.Pet, error) {
	pet, err := r.queries.CreatePet(ctx, args)
	if err != nil {
		return sqlc.Pet{}, err
	}
	return pet, nil
}

func (r *PetRepository) GetPetById(ctx context.Context, petId int32) (sqlc.Pet, error) {
	pet, err := r.queries.GetPetByID(ctx, petId)
	if err != nil {
		return sqlc.Pet{}, err
	}
	return pet, nil
}

func (r *PetRepository) UpdatePetById(ctx context.Context, params sqlc.UpdatePetParams) error {
	err := r.queries.UpdatePet(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (r *PetRepository) DeletePetById(ctx context.Context, petId int32) error {
	err := r.queries.DeletePet(ctx, petId)
	if err != nil {
		return err
	}
	return nil
}
