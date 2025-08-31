package persistence

import (
	"context"
	"fmt"
	"sync"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
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

func (r *SqlcOwnerRepository) Save(ctx context.Context, owner *entity.Owner) error {
	if owner.GetID().IsZero() {
		if err := r.create(ctx, owner); err != nil {
			return DBCreateError(err.Error())
		}
		return nil
	}

	if err := r.update(ctx, owner); err != nil {
		return DBUpdateError(err.Error())
	}

	return nil
}

func (r *SqlcOwnerRepository) GetByID(ctx context.Context, id valueobject.OwnerID) (entity.Owner, error) {
	petsChan := make(chan struct {
		pets []entity.Pet
		err  error
	}, 1)

	var wg sync.WaitGroup

	ownerRow, err := r.queries.GetOwnerByID(ctx, int32(id.GetValue()))
	if err != nil {
		return entity.Owner{}, fmt.Errorf("failed to get owner by ID: %w", err)
	}

	var pets []entity.Pet
	wg.Add(1)
	go func() {
		defer wg.Done()
		p, petErr := r.petRepository.ListByOwnerID(ctx, id)
		petsChan <- struct {
			pets []entity.Pet
			err  error
		}{pets: p, err: petErr}
	}()

	petsResult := <-petsChan
	if petsResult.err != nil {
		return entity.Owner{}, fmt.Errorf("failed to list pets for owner: %w", petsResult.err)
	}
	pets = petsResult.pets

	owner := rowToOwner(ownerRow)
	owner.SetPets(pets)
	return owner, nil
}

func (r *SqlcOwnerRepository) GetByPhone(ctx context.Context, phone string) (entity.Owner, error) {
	sqlRow, err := r.queries.GetOwnerByPhone(ctx, phone)
	if err != nil {
		return entity.Owner{}, err
	}

	return rowToOwner(sqlRow), nil
}

// Add Seacrh
func (r *SqlcOwnerRepository) List(ctx context.Context, pagination page.PageData) (page.Page[[]entity.Owner], error) {
	pageParams := sqlc.ListOwnersParams{
		Limit:  int32(pagination.PageSize),
		Offset: int32((pagination.PageNumber - 1) * pagination.PageSize),
	}

	ownerRow, err := r.queries.ListOwners(ctx, pageParams)
	if err != nil {
		return page.Page[[]entity.Owner]{}, err
	}

	owners, err := ListRowToOwner(ownerRow)
	if err != nil {
		return page.Page[[]entity.Owner]{}, err
	}

	return page.NewPage(owners, *page.GetPageMetadata(len(owners), pagination)), nil
}

func (r *SqlcOwnerRepository) SoftDelete(ctx context.Context, id valueobject.OwnerID) error {
	if err := r.queries.SoftDeleteOwner(ctx, int32(id.GetValue())); err != nil {
		return DBDeleteError(err.Error())
	}
	return nil
}

func (r *SqlcOwnerRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	exists, err := r.queries.ExistByPhoneNumber(ctx, phone)
	if err != nil {
		return false, DBSelectFoundError(err.Error())
	}

	return exists, nil
}

func (r *SqlcOwnerRepository) ExistsByID(ctx context.Context, id valueobject.OwnerID) (bool, error) {
	exists, err := r.queries.ExistByID(ctx, int32(id.GetValue()))
	if err != nil {
		return false, DBSelectFoundError(err.Error())
	}
	return exists, nil
}

func (r *SqlcOwnerRepository) ActivateOwner(ctx context.Context, id int) error {
	if err := r.queries.ActivateUser(ctx, int32(id)); err != nil {
		return DBSelectFoundError(err.Error())
	}
	return nil
}

func (r *SqlcOwnerRepository) DeactivateOwner(ctx context.Context, id int) error {
	if err := r.queries.DeactivateUser(ctx, int32(id)); err != nil {
		return DBSelectFoundError(err.Error())
	}
	return nil
}

func (r *SqlcOwnerRepository) create(ctx context.Context, owner *entity.Owner) error {
	createParams := toCreateParams(*owner)

	ownerCreated, err := r.queries.CreateOwner(ctx, *createParams)
	if err != nil {
		return err
	}

	ownerID, _ := valueobject.NewOwnerID(int(ownerCreated.ID))
	owner.SetID(ownerID)
	return nil
}

func (r *SqlcOwnerRepository) update(ctx context.Context, owner *entity.Owner) error {
	params := toUpdateParams(*owner)

	if err := r.queries.UpdateOwner(ctx, *params); err != nil {
		return err
	}

	return nil
}
