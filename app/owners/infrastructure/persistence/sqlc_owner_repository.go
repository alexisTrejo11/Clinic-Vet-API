package ownerRepository

import (
	"context"
	"fmt"
	"sync"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SqlcOwnerRepository struct {
	queries       *sqlc.Queries
	petRepository petRepository.PetRepository
}

func NewSqlcOwnerRepository(queries *sqlc.Queries, petRepository petRepository.PetRepository) ownerDomain.OwnerRepository {
	return &SqlcOwnerRepository{
		queries:       queries,
		petRepository: petRepository,
	}
}

func (r *SqlcOwnerRepository) Save(ctx context.Context, owner *ownerDomain.Owner) error {
	if owner.Id() == 0 {
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

func (r *SqlcOwnerRepository) GetById(ctx context.Context, id int) (ownerDomain.Owner, error) {
	petsChan := make(chan struct {
		pets []petDomain.Pet
		err  error
	}, 1)

	var wg sync.WaitGroup

	ownerRow, err := r.queries.GetOwnerByID(ctx, int32(id))
	if err != nil {
		return ownerDomain.Owner{}, fmt.Errorf("failed to get owner by ID: %w", err)
	}

	var pets []petDomain.Pet
	wg.Add(1)
	go func() {
		defer wg.Done()
		p, petErr := r.petRepository.ListByOwnerId(ctx, id)
		petsChan <- struct {
			pets []petDomain.Pet
			err  error
		}{pets: p, err: petErr}
	}()

	petsResult := <-petsChan
	if petsResult.err != nil {
		return ownerDomain.Owner{}, fmt.Errorf("failed to list pets for owner: %w", petsResult.err)
	}
	pets = petsResult.pets

	owner := rowToOwner(ownerRow)
	owner.SetPets(pets)
	return owner, nil
}

func (r *SqlcOwnerRepository) GetByPhone(ctx context.Context, phone string) (ownerDomain.Owner, error) {
	sqlRow, err := r.queries.GetOwnerByPhone(ctx, phone)
	if err != nil {
		return ownerDomain.Owner{}, err
	}

	return rowToOwner(sqlRow), nil
}

// Add Seacrh
func (r *SqlcOwnerRepository) List(ctx context.Context, pagination page.PageData) (page.Page[[]ownerDomain.Owner], error) {
	pageParams := sqlc.ListOwnersParams{
		Limit:  int32(pagination.PageSize),
		Offset: int32((pagination.PageNumber - 1) * pagination.PageSize),
	}

	ownerRow, err := r.queries.ListOwners(ctx, pageParams)
	if err != nil {
		return page.Page[[]ownerDomain.Owner]{}, err
	}

	owners, err := ListRowToOwner(ownerRow)
	if err != nil {
		return page.Page[[]ownerDomain.Owner]{}, err
	}

	return page.NewPage(owners, *page.GetPageMetadata(len(owners), pagination)), nil
}

func (r *SqlcOwnerRepository) SoftDelete(ctx context.Context, OwnerId int) error {
	if err := r.queries.SoftDeleteOwner(ctx, int32(OwnerId)); err != nil {
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

func (r *SqlcOwnerRepository) ExistsByID(ctx context.Context, id int) (bool, error) {
	exists, err := r.queries.ExistByID(ctx, int32(id))
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

func (r *SqlcOwnerRepository) create(ctx context.Context, owner *ownerDomain.Owner) error {
	createParams := toCreateParams(*owner)

	ownerCreated, err := r.queries.CreateOwner(ctx, *createParams)
	if err != nil {
		return err
	}

	owner.SetId(int(ownerCreated.ID))
	return nil
}

func (r *SqlcOwnerRepository) update(ctx context.Context, owner *ownerDomain.Owner) error {
	params := toUpdateParams(*owner)

	if err := r.queries.UpdateOwner(ctx, *params); err != nil {
		return err
	}

	return nil
}
