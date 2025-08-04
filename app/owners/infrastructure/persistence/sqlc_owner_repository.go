package sqlcOwnerRepository

import (
	"context"
	"fmt"
	"sync"

	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/repositories"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	userEnums "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/enum"
	userValueObjects "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/valueobjects"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SqlcOwnerRepository struct {
	queries       *sqlc.Queries
	petRepository petRepository.PetRepository
}

func NewSlqcOwnerRepository(queries *sqlc.Queries, petRepository petRepository.PetRepository) ownerRepository.OwnerRepository {
	return &SqlcOwnerRepository{
		queries:       queries,
		petRepository: petRepository,
	}
}

func (r *SqlcOwnerRepository) Save(ctx context.Context, owner *ownerDomain.Owner) error {
	if owner.Id == 0 {
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

func (r *SqlcOwnerRepository) GetById(ctx context.Context, id int, includePets bool) (ownerDomain.Owner, error) {
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
	if includePets {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p, petErr := r.petRepository.ListByOwnerId(ctx, id)
			petsChan <- struct {
				pets []petDomain.Pet
				err  error
			}{pets: p, err: petErr}
		}()
	} else {
	}

	ownerName, err := userValueObjects.NewPersonName(ownerRow.FirstName, ownerRow.LastName)
	if err != nil {
		return ownerDomain.Owner{}, fmt.Errorf("failed to create owner name: %w", err)
	}

	if includePets {
		petsResult := <-petsChan
		if petsResult.err != nil {
			return ownerDomain.Owner{}, fmt.Errorf("failed to list pets for owner: %w", petsResult.err)
		}
		pets = petsResult.pets
	}

	owner := ownerDomain.Owner{
		Id:          int(ownerRow.ID),
		Photo:       ownerRow.Photo,
		FullName:    ownerName,
		DateOfBirth: ownerRow.DateOfBirth.Time,
		Gender:      userEnums.Gender(ownerRow.Gender),
		PhoneNumber: ownerRow.PhoneNumber,
		Address:     &ownerRow.Address.String,
		IsActive:    ownerRow.IsActive,
		Pets:        pets,
	}

	return owner, nil
}

func (r *SqlcOwnerRepository) GetByPhone(ctx context.Context, phone string) (ownerDomain.Owner, error) {
	row, err := r.queries.GetOwnerByPhone(ctx, phone)
	if err != nil {
		return ownerDomain.Owner{}, err
	}

	ownerName, err := userValueObjects.NewPersonName(row.FirstName, row.LastName)
	if err != nil {
		return ownerDomain.Owner{}, nil
	}

	owner := ownerDomain.Owner{
		Id:          int(row.ID),
		Photo:       row.Photo,
		FullName:    ownerName,
		DateOfBirth: row.DateOfBirth.Time,
		Gender:      userEnums.Gender(row.Gender),
		PhoneNumber: row.PhoneNumber,
		Address:     &row.Address.String,
		IsActive:    row.IsActive,
	}

	return owner, nil
}

// Add Seacrh
func (r *SqlcOwnerRepository) List(ctx context.Context, pagination page.PageData) ([]ownerDomain.Owner, error) {
	pageParams := sqlc.ListOwnersParams{Limit: int32(pagination.PageNumber), Offset: int32(pagination.PageNumber - 1)}
	ownerRow, err := r.queries.ListOwners(ctx, pageParams)
	if err != nil {
		return []ownerDomain.Owner{}, err
	}

	owners, err := ListRowToOwner(ownerRow)
	if err != nil {
		return []ownerDomain.Owner{}, err
	}

	return owners, nil
}

func (r *SqlcOwnerRepository) Delete(ctx context.Context, OwnerId int) error {
	if err := r.queries.DeleteOwner(ctx, int32(OwnerId)); err != nil {
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
	params := ToCreateParams(*owner)

	ownerCreated, err := r.queries.CreateOwner(ctx, *params)
	if err != nil {
		return err
	}

	owner.Id = int(ownerCreated.ID)
	return nil
}

func (r *SqlcOwnerRepository) update(ctx context.Context, owner *ownerDomain.Owner) error {
	params := ToUpdateParams(*owner)

	err := r.queries.UpdateOwner(ctx, *params)
	if err != nil {
		return err
	}

	return nil
}
