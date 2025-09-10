package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/owner"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/pet"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SqlcOwnerRepository struct {
	queries       *sqlc.Queries
	petRepository repository.PetRepository
}

func NewSqlcOwnerRepository(queries *sqlc.Queries, petRepository repository.PetRepository) repository.OwnerRepository {
	return &SqlcOwnerRepository{
		queries:       queries,
		petRepository: petRepository,
	}
}

// Save creates or updates an owner based on whether it has an ID
func (r *SqlcOwnerRepository) Save(ctx context.Context, owner *owner.Owner) error {
	if owner.ID().IsZero() {
		return r.create(ctx, owner)
	}
	return r.update(ctx, owner)
}

func (r *SqlcOwnerRepository) GetByID(ctx context.Context, id valueobject.OwnerID) (owner.Owner, error) {
	type petsResult struct {
		pets []pet.Pet
		err  error
	}

	petsChan := make(chan petsResult, ConcurrentOpTimeout)
	var wg sync.WaitGroup

	ownerRow, err := r.queries.GetOwnerByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return owner.Owner{}, r.notFoundError("id", fmt.Sprintf("%d", id.Value()))
		}
		return owner.Owner{}, r.dbError(OpSelect, fmt.Sprintf("failed to get owner with ID %d", id.Value()), err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(petsChan)

		pets, petErr := r.petRepository.ListByOwnerID(ctx, id)
		petsChan <- petsResult{pets: pets, err: petErr}
	}()

	// Wait for pets to be fetched
	wg.Wait()
	petsRes := <-petsChan

	if petsRes.err != nil {
		return owner.Owner{}, fmt.Errorf("%s with ID %d: %w", "", id.Value(), petsRes.err)
	}

	ownerEntity, err := sqlRowToOwner(ownerRow, petsRes.pets)
	if err != nil {
		return owner.Owner{}, r.wrapConversionError(err)
	}

	return ownerEntity, nil
}

func (r *SqlcOwnerRepository) GetByPhone(ctx context.Context, phone string) (owner.Owner, error) {
	sqlRow, err := r.queries.GetOwnerByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return owner.Owner{}, r.notFoundError("phone", phone)
		}
		return owner.Owner{}, r.dbError(OpSelect, fmt.Sprintf("failed to get owner with phone %s", phone), err)
	}

	ownerEntity, err := sqlRowToOwner(sqlRow, nil)
	if err != nil {
		return owner.Owner{}, r.wrapConversionError(err)
	}

	return ownerEntity, nil
}

// Search retrieves owners with pagination
func (r *SqlcOwnerRepository) Search(ctx context.Context, search any) (page.Page[[]owner.Owner], error) {
	return page.Page[[]owner.Owner]{}, errors.New("not implemented")
}

func (r *SqlcOwnerRepository) SoftDelete(ctx context.Context, id valueobject.OwnerID) error {
	if err := r.queries.SoftDeleteOwner(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("failed to soft delete owner with ID %d", id.Value()), err)
	}
	return nil
}

func (r *SqlcOwnerRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	exists, err := r.queries.ExistByPhoneNumber(ctx, phone)
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("failed to check existence by phone %s", phone), err)
	}
	return exists, nil
}

// ExistsByID checks if an owner exists with the given ID
func (r *SqlcOwnerRepository) ExistsByID(ctx context.Context, id valueobject.OwnerID) (bool, error) {
	exists, err := r.queries.ExistByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("failed to check existence by ID %d", id.Value()), err)
	}
	return exists, nil
}

func (r *SqlcOwnerRepository) ActivateOwner(ctx context.Context, id int) error {
	if err := r.queries.ActivateUser(ctx, int32(id)); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to activate owner with ID %d", id), err)
	}
	return nil
}

func (r *SqlcOwnerRepository) DeactivateOwner(ctx context.Context, id int) error {
	if err := r.queries.DeactivateUser(ctx, int32(id)); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to deactivate owner with ID %d", id), err)
	}
	return nil
}

// create inserts a new owner into the database
func (r *SqlcOwnerRepository) create(ctx context.Context, owner *owner.Owner) error {
	createParams := toCreateParams(*owner)

	_, err := r.queries.CreateOwner(ctx, *createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateOwner, err)
	}

	return nil
}

// update modifies an existing owner in the database
func (r *SqlcOwnerRepository) update(ctx context.Context, owner *owner.Owner) error {
	params := entityToUpdateParams(*owner)

	if err := r.queries.UpdateOwner(ctx, *params); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to update owner with ID %d", owner.ID().Value()), err)
	}

	return nil
}
