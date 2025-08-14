package sqlcPetRepository

import (
	"context"
	"fmt"

	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SqlcPetRepository struct {
	queries *sqlc.Queries
}

func NewSqlcPetRepository(queries *sqlc.Queries) petRepository.PetRepository {
	return &SqlcPetRepository{
		queries: queries,
	}
}

func (r *SqlcPetRepository) List(ctx context.Context) ([]petDomain.Pet, error) {
	sqlPets, err := r.queries.ListPets(ctx)
	if err != nil {
		return []petDomain.Pet{}, DBSelectFoundError(err.Error())
	}

	pets := make([]petDomain.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, fmt.Errorf("error while mapping pet %d: %w", sqlPet.ID, err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) ListByOwnerId(ctx context.Context, ownerId int) ([]petDomain.Pet, error) {
	sqlPets, err := r.queries.GetPetsByOwnerID(ctx, int32(ownerId))
	if err != nil {
		return []petDomain.Pet{}, DBSelectFoundError(err.Error())
	}

	pets := make([]petDomain.Pet, len(sqlPets))
	for i, sqlPet := range sqlPets {
		domainPet, err := ToDomainPet(sqlPet)
		if err != nil {
			return nil, fmt.Errorf("error while mapping pet %d: %w", sqlPet.ID, err)
		}
		pets[i] = *domainPet
	}

	return pets, nil
}

func (r *SqlcPetRepository) GetById(ctx context.Context, petId int) (petDomain.Pet, error) {
	sqlPet, err := r.queries.GetPetByID(ctx, int32(petId))
	if err != nil {
		return petDomain.Pet{}, DBSelectFoundError(err.Error())
	}

	domainPet, err := ToDomainPet(sqlPet)
	if err != nil {
		return petDomain.Pet{}, err
	}

	return *domainPet, nil
}

func (r *SqlcPetRepository) Save(ctx context.Context, pet *petDomain.Pet) error {
	if pet.GetID() == 0 {
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

func (r *SqlcPetRepository) ExistsById(ctx context.Context, petId int) (bool, error) {
	return true, nil
}

func (r *SqlcPetRepository) create(ctx context.Context, pet *petDomain.Pet) error {
	params := ToSqlCreateParam(pet)

	petCreated, err := r.queries.CreatePet(ctx, *params)
	if err != nil {
		return err
	}

	_, err = ToDomainPet(petCreated)
	if err != nil {
		return err
	}

	pet.SetID(int(petCreated.ID))

	return nil
}

func (r *SqlcPetRepository) update(ctx context.Context, pet *petDomain.Pet) error {
	params := ToSqlUpdateParam(pet)

	err := r.queries.UpdatePet(ctx, *params)
	if err != nil {
		return err
	}

	return nil
}

func (r *SqlcPetRepository) Delete(ctx context.Context, petId int) error {
	if err := r.queries.DeletePet(ctx, int32(petId)); err != nil {
		return DBDeleteError(err.Error())

	}
	return nil
}
