package persistence

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
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

func (r *SqlcPetRepository) List(ctx context.Context) ([]entity.Pet, error) {
	sqlPets, err := r.queries.ListPets(ctx)
	if err != nil {
		return []entity.Pet{}, DBSelectFoundError(err.Error())
	}

	pets := make([]entity.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, fmt.Errorf("error while mapping pet %d: %w", sqlPet.ID, err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) ListByOwnerID(ctx context.Context, ownerID valueobject.OwnerID) ([]entity.Pet, error) {
	sqlPets, err := r.queries.GetPetsByOwnerID(ctx, int32(ownerID.GetValue()))
	if err != nil {
		return []entity.Pet{}, DBSelectFoundError(err.Error())
	}

	pets := make([]entity.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, fmt.Errorf("error while mapping pet %d: %w", sqlPet.ID, err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) GetByID(ctx context.Context, petID valueobject.PetID) (entity.Pet, error) {
	sqlPet, err := r.queries.GetPetByID(ctx, int32(petID.GetValue()))
	if err != nil {
		return entity.Pet{}, DBSelectFoundError(err.Error())
	}

	domainPet, err := ToDomainPet(sqlPet)
	if err != nil {
		return entity.Pet{}, err
	}

	return *domainPet, nil
}

func (r *SqlcPetRepository) Save(ctx context.Context, pet *entity.Pet) error {
	if pet.GetID().IsZero() {
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

func (r *SqlcPetRepository) create(ctx context.Context, pet *entity.Pet) error {
	params := ToSqlCreateParam(pet)

	petCreated, err := r.queries.CreatePet(ctx, *params)
	if err != nil {
		return err
	}

	_, err = ToDomainPet(petCreated)
	if err != nil {
		return err
	}

	petID, err := valueobject.NewPetID(int(petCreated.ID))
	if err != nil {
		return err
	}
	pet.SetID(petID)

	return nil
}

func (r *SqlcPetRepository) update(ctx context.Context, pet *entity.Pet) error {
	params := ToSqlUpdateParam(pet)

	err := r.queries.UpdatePet(ctx, *params)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlcPetRepository) Delete(ctx context.Context, petID valueobject.PetID) error {
	if err := r.queries.DeletePet(ctx, int32(petID.GetValue())); err != nil {
		return DBDeleteError(err.Error())
	}
	return nil
}
