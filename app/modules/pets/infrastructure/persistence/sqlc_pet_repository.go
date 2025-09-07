// Package persistence defines the persistence layer implementations.
package persistence

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"

	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SqlcPetRepository struct {
	queries *sqlc.Queries
}

func NewSqlcPetRepository(queries *sqlc.Queries) repository.PetRepository {
	return &SqlcPetRepository{
		queries: queries,
	}
}

func (r *SqlcPetRepository) GetByIDAndOwnerID(ctx context.Context, id valueobject.PetID, petID valueobject.OwnerID) (pet.Pet, error) {
	return pet.Pet{}, nil
}

func (r *SqlcPetRepository) Search(ctx context.Context, searchParmas any) ([]pet.Pet, error) {
	sqlPets, err := r.queries.ListPets(ctx)
	if err != nil {
		return []pet.Pet{}, DBSelectFoundError(err.Error())
	}

	pets := make([]pet.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, fmt.Errorf("error while mapping pet %d: %w", sqlPet.ID, err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID) ([]pet.Pet, error) {
	sqlPets, err := r.queries.GetPetsByOwnerID(ctx, int32(ownerID.Value()))
	if err != nil {
		return []pet.Pet{}, DBSelectFoundError(err.Error())
	}

	pets := make([]pet.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, fmt.Errorf("error while mapping pet %d: %w", sqlPet.ID, err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) GetByID(ctx context.Context, petID valueobject.PetID) (pet.Pet, error) {
	sqlPet, err := r.queries.GetPetByID(ctx, int32(petID.Value()))
	if err != nil {
		return pet.Pet{}, DBSelectFoundError(err.Error())
	}

	domainPet, err := ToDomainPet(sqlPet)
	if err != nil {
		return pet.Pet{}, err
	}

	return *domainPet, nil
}

func (r *SqlcPetRepository) Save(ctx context.Context, pet *pet.Pet) error {
	if pet.ID().IsZero() {
		if err := r.create(ctx, pet); err != nil {
			return DBCreateError(err.Error())
		}
		return nil
	}

	if err := r.update(ctx, pet); err != nil {
		return DBUpdateError(err.Error())
	}

	return nil
}

func (r *SqlcPetRepository) ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error) {
	return true, nil
}

func (r *SqlcPetRepository) create(ctx context.Context, pet *pet.Pet) error {
	params := ToSqlCreateParam(pet)

	_, err := r.queries.CreatePet(ctx, *params)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlcPetRepository) update(ctx context.Context, pet *pet.Pet) error {
	params := ToSqlUpdateParam(pet)

	err := r.queries.UpdatePet(ctx, *params)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlcPetRepository) Delete(ctx context.Context, petID valueobject.PetID) error {
	if err := r.queries.DeletePet(ctx, int32(petID.Value())); err != nil {
		return DBDeleteError(err.Error())
	}
	return nil
}
