// Package persistence defines the persistence layer implementations.
package persistence

import (
	"context"
	"database/sql"
	"errors"
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

func (r *SqlcPetRepository) GetByIDAndOwnerID(ctx context.Context, id valueobject.PetID, ownerID valueobject.OwnerID) (pet.Pet, error) {
	sqlPet, err := r.queries.GetPetByIDAndOwnerID(ctx, sqlc.GetPetByIDAndOwnerIDParams{
		ID:      int32(id.Value()),
		OwnerID: int32(ownerID.Value()),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pet.Pet{}, r.notFoundError("id and owner_id", fmt.Sprintf("pet %d for owner %d", id.Value(), ownerID.Value()))
		}
		return pet.Pet{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d and owner ID %d", ErrMsgGetPet, id.Value(), ownerID.Value()), err)
	}

	domainPet, err := ToDomainPet(sqlPet)
	if err != nil {
		return pet.Pet{}, r.wrapConversionError(err)
	}

	return *domainPet, nil
}

func (r *SqlcPetRepository) Search(ctx context.Context, searchParams any) ([]pet.Pet, error) {
	sqlPets, err := r.queries.ListPets(ctx)
	if err != nil {
		return []pet.Pet{}, r.dbError(OpSelect, ErrMsgSearchPets, err)
	}

	pets := make([]pet.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, r.wrapConversionError(err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID) ([]pet.Pet, error) {
	sqlPets, err := r.queries.GetPetsByOwnerID(ctx, int32(ownerID.Value()))
	if err != nil {
		return []pet.Pet{}, r.dbError(OpSelect, fmt.Sprintf("%s for owner ID %d", ErrMsgGetPetByOwnerID, ownerID.Value()), err)
	}

	pets := make([]pet.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, r.wrapConversionError(err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) GetByID(ctx context.Context, petID valueobject.PetID) (pet.Pet, error) {
	sqlPet, err := r.queries.GetPetByID(ctx, int32(petID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return pet.Pet{}, r.notFoundError("id", petID.String())
		}
		return pet.Pet{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgGetPet, petID.Value()), err)
	}

	domainPet, err := ToDomainPet(sqlPet)
	if err != nil {
		return pet.Pet{}, r.wrapConversionError(err)
	}

	return *domainPet, nil
}

func (r *SqlcPetRepository) Save(ctx context.Context, pet *pet.Pet) error {
	if pet.ID().IsZero() {
		return r.create(ctx, pet)
	}
	return r.update(ctx, pet)
}

func (r *SqlcPetRepository) ExistsByID(ctx context.Context, petID valueobject.PetID) (bool, error) {
	_, err := r.queries.GetPetByID(ctx, int32(petID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgCheckPetExists, petID.Value()), err)
	}
	return true, nil
}

func (r *SqlcPetRepository) create(ctx context.Context, pet *pet.Pet) error {
	params := ToSqlCreateParam(pet)

	_, err := r.queries.CreatePet(ctx, *params)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreatePet, err)
	}

	return nil
}

func (r *SqlcPetRepository) update(ctx context.Context, pet *pet.Pet) error {
	params := ToSqlUpdateParam(pet)

	err := r.queries.UpdatePet(ctx, *params)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s with ID %d", ErrMsgUpdatePet, pet.ID().Value()), err)
	}

	return nil
}

func (r *SqlcPetRepository) Delete(ctx context.Context, petID valueobject.PetID) error {
	if err := r.queries.DeletePet(ctx, int32(petID.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgDeletePet, petID.Value()), err)
	}
	return nil
}
