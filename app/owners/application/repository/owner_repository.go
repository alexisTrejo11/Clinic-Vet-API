package ownerRepository

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner *ownerDomain.Owner) error

	List(ctx context.Context, query string, limit, offset int) ([]ownerDomain.Owner, error)
	GetByID(ctx context.Context, id uint) (ownerDomain.Owner, error)
	GetByPhone(ctx context.Context, phone string) (ownerDomain.Owner, error)

	Delete(ctx context.Context, id uint) error

	//GetByName(ctx context.Context, name string) ([]ownerDomain.Owner, error)
	//ListActiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	//ListInactiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	//ListOwnersWithPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	//ListOwnersWithoutPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)

	//CountOwners(ctx context.Context) (int64, error)
	//CountActiveOwners(ctx context.Context) (int64, error)
	//CountInactiveOwners(ctx context.Context) (int64, error)

	//ActivateOwner(ctx context.Context, id uint) error
	//DeactivateOwner(ctx context.Context, id uint) error

	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id uint) (bool, error)
}
