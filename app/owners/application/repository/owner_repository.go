package ownerRepository

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner *ownerDomain.Owner) error
	GetByID(ctx context.Context, id uint) (ownerDomain.Owner, error)
	Delete(ctx context.Context, id uint) error

	GetByPhone(ctx context.Context, phone string) (ownerDomain.Owner, error)

	List(ctx context.Context, query string, limit, offset int) ([]ownerDomain.Owner, error)

	//GetByName(ctx context.Context, name string) ([]ownerDomain.Owner, error)
	ListActiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	ListInactiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	ListOwnersWithPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	ListOwnersWithoutPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)

	CountOwners(ctx context.Context) (int64, error)
	CountActiveOwners(ctx context.Context) (int64, error)
	CountInactiveOwners(ctx context.Context) (int64, error)

	//ActivateOwner(id uint) error
	//DeactivateOwner(id uint) error
	//ExistsByPhone(phone string) (bool, error)
	//ExistsByID(id uint) (bool, error)
}
