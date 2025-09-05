package persistence

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

// SqlcOwnerRepository implements OwnerRepository using SQLC
type SqlcOwnerRepository struct {
	queries       *sqlc.Queries
	petRepository repository.PetRepository
}

// NewSqlcOwnerRepository creates a new owner repository instance
func NewSqlcOwnerRepository(queries *sqlc.Queries, petRepository repository.PetRepository) repository.OwnerRepository {
	return &SqlcOwnerRepository{
		queries:       queries,
		petRepository: petRepository,
	}
}

// Save creates or updates an owner based on whether it has an ID
func (r *SqlcOwnerRepository) Save(ctx context.Context, owner *entity.Owner) error {
	if owner.GetID().IsZero() {
		return r.create(ctx, owner)
	}
	return r.update(ctx, owner)
}

// GetByID retrieves an owner by ID along with their pets using concurrent fetching
func (r *SqlcOwnerRepository) GetByID(ctx context.Context, id valueobject.OwnerID) (entity.Owner, error) {
	type petsResult struct {
		pets []entity.Pet
		err  error
	}

	petsChan := make(chan petsResult, ConcurrentOpTimeout)
	var wg sync.WaitGroup

	// Fetch owner data
	ownerRow, err := r.queries.GetOwnerByID(ctx, int32(id.GetValue()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Owner{}, r.notFoundError("id", fmt.Sprintf("%d", id.GetValue()))
		}
		return entity.Owner{}, r.dbError(OpSelect, fmt.Sprintf("failed to get owner with ID %d", id.GetValue()), err)
	}

	// Concurrently fetch pets for the owner
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
		return entity.Owner{}, fmt.Errorf("%s with ID %d: %w", ErrMsgListPetsForOwner, id.GetValue(), petsRes.err)
	}

	// Convert SQL row to domain entity
	owner, err := sqlRowToOwner(ownerRow)
	if err != nil {
		return entity.Owner{}, r.wrapConversionError(err)
	}

	// Set the pets for the owner
	owner.SetPets(petsRes.pets)
	return owner, nil
}

// GetByPhone retrieves an owner by phone number
func (r *SqlcOwnerRepository) GetByPhone(ctx context.Context, phone string) (entity.Owner, error) {
	sqlRow, err := r.queries.GetOwnerByPhone(ctx, phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Owner{}, r.notFoundError("phone", phone)
		}
		return entity.Owner{}, r.dbError(OpSelect, fmt.Sprintf("failed to get owner with phone %s", phone), err)
	}

	owner, err := sqlRowToOwner(sqlRow)
	if err != nil {
		return entity.Owner{}, r.wrapConversionError(err)
	}

	return owner, nil
}

// List retrieves owners with pagination
func (r *SqlcOwnerRepository) List(ctx context.Context, pagination page.PageData) (page.Page[[]entity.Owner], error) {
	pageParams := sqlc.ListOwnersParams{
		Limit:  int32(pagination.PageSize),
		Offset: r.calculateOffset(pagination),
	}

	ownerRows, err := r.queries.ListOwners(ctx, pageParams)
	if err != nil {
		return page.Page[[]entity.Owner]{}, r.dbError(OpSelect, ErrMsgListOwners, err)
	}

	owners, err := ListRowToOwner(ownerRows)
	if err != nil {
		return page.Page[[]entity.Owner]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountOwners(ctx)
	if err != nil {
		return page.Page[[]entity.Owner]{}, r.dbError(OpCount, ErrMsgCountOwners, err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pagination)
	return page.NewPage(owners, *pageMetadata), nil
}

// SoftDelete marks an owner as deleted without removing the record
func (r *SqlcOwnerRepository) SoftDelete(ctx context.Context, id valueobject.OwnerID) error {
	if err := r.queries.SoftDeleteOwner(ctx, int32(id.GetValue())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("failed to soft delete owner with ID %d", id.GetValue()), err)
	}
	return nil
}

// ExistsByPhone checks if an owner exists with the given phone number
func (r *SqlcOwnerRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	exists, err := r.queries.ExistByPhoneNumber(ctx, phone)
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("failed to check existence by phone %s", phone), err)
	}
	return exists, nil
}

// ExistsByID checks if an owner exists with the given ID
func (r *SqlcOwnerRepository) ExistsByID(ctx context.Context, id valueobject.OwnerID) (bool, error) {
	exists, err := r.queries.ExistByID(ctx, int32(id.GetValue()))
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("failed to check existence by ID %d", id.GetValue()), err)
	}
	return exists, nil
}

// ActivateOwner activates an owner account
func (r *SqlcOwnerRepository) ActivateOwner(ctx context.Context, id int) error {
	if err := r.queries.ActivateUser(ctx, int32(id)); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to activate owner with ID %d", id), err)
	}
	return nil
}

// DeactivateOwner deactivates an owner account
func (r *SqlcOwnerRepository) DeactivateOwner(ctx context.Context, id int) error {
	if err := r.queries.DeactivateUser(ctx, int32(id)); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to deactivate owner with ID %d", id), err)
	}
	return nil
}

// create inserts a new owner into the database
func (r *SqlcOwnerRepository) create(ctx context.Context, owner *entity.Owner) error {
	createParams := toCreateParams(*owner)

	ownerCreated, err := r.queries.CreateOwner(ctx, *createParams)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateOwner, err)
	}

	ownerID, err := valueobject.NewOwnerID(int(ownerCreated.ID))
	if err != nil {
		return fmt.Errorf("%s: %w", ErrMsgCreateOwnerID, err)
	}

	owner.SetID(ownerID)
	return nil
}

// update modifies an existing owner in the database
func (r *SqlcOwnerRepository) update(ctx context.Context, owner *entity.Owner) error {
	params := entityToUpdateParams(*owner)

	if err := r.queries.UpdateOwner(ctx, *params); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("failed to update owner with ID %d", owner.GetID().GetValue()), err)
	}

	return nil
}
