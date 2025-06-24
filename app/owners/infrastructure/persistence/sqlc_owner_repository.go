package sqlcOwnerRepository

import (
	"context"

	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type SqlcOwnerRepository struct {
	queries *sqlc.Queries
}

func NewSlqcOwnerRepository(queries *sqlc.Queries) ownerRepository.OwnerRepository {
	return &SqlcOwnerRepository{queries: queries}
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

func (r *SqlcOwnerRepository) create(ctx context.Context, Owner *ownerDomain.Owner) error {
	params := ToCreateParams(*Owner)

	ownerCreated, err := r.queries.CreateOwner(ctx, *params)
	if err != nil {
		return err
	}

	Owner.Id = uint(ownerCreated.ID)

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

func (r *SqlcOwnerRepository) GetByID(ctx context.Context, id uint) (ownerDomain.Owner, error) {
	ownerRow, err := r.queries.GetOwnerByID(ctx, int32(id))
	if err != nil {
		return ownerDomain.Owner{}, nil
	}

	owner, err := RowToOwner(ownerRow)
	if err != nil {
		return ownerDomain.Owner{}, nil
	}

	return owner, nil
}

// Add Seacrh
func (r *SqlcOwnerRepository) List(ctx context.Context, query string, limit, offset int) ([]ownerDomain.Owner, error) {
	ownerRow, err := r.queries.ListOwners(ctx)
	if err != nil {
		return []ownerDomain.Owner{}, nil
	}

	owners, err := ListRowToOwner(ownerRow)
	if err != nil {
		return []ownerDomain.Owner{}, nil
	}

	return owners, nil
}

func (r *SqlcOwnerRepository) Delete(ctx context.Context, OwnerId uint) error {
	if err := r.queries.DeleteOwner(ctx, int32(OwnerId)); err != nil {
		return DBDeleteError(err.Error())
	}
	return nil
}

func (r *SqlcOwnerRepository) GetByPhone(ctx context.Context, phone string) (ownerDomain.Owner, error) {
	return ownerDomain.Owner{}, nil
}
func (r *SqlcOwnerRepository) ListByName(ctx context.Context, name string) ([]ownerDomain.Owner, error) {
	return nil, nil
}
func (r *SqlcOwnerRepository) ListActiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error) {
	return nil, nil
}
func (r *SqlcOwnerRepository) ListInactiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error) {
	return nil, nil
}
func (r *SqlcOwnerRepository) ListOwnersWithOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error) {
	return nil, nil
}
func (r *SqlcOwnerRepository) ListOwnersWithoutOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error) {
	return nil, nil
}

func (r *SqlcOwnerRepository) ActivateOwner(ctx context.Context, id uint) error {
	return nil
}
func (r *SqlcOwnerRepository) DeactivateOwner(ctx context.Context, id uint) error {
	return nil
}

func (r *SqlcOwnerRepository) CountOwners(ctx context.Context) (int64, error) {
	return 0, nil
}
func (r *SqlcOwnerRepository) CountActiveOwners(ctx context.Context) (int64, error) {
	return 0, nil
}
func (r *SqlcOwnerRepository) CountInactiveOwners(ctx context.Context) (int64, error) {
	return 0, nil
}

func (r *SqlcOwnerRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	return false, nil
}
func (r *SqlcOwnerRepository) ExistsByID(ctx context.Context, id uint) (bool, error) {
	return false, nil
}

func (r *SqlcOwnerRepository) ListOwnersWithPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error) {
	return []ownerDomain.Owner{}, nil
}
func (r *SqlcOwnerRepository) ListOwnersWithoutPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error) {
	return []ownerDomain.Owner{}, nil
}
