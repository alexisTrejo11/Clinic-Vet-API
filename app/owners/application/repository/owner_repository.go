package ownerRepository

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner *ownerDomain.Owner) error

	List(ctx context.Context, pagination page.PageData) ([]ownerDomain.Owner, error)
	GetById(ctx context.Context, id int, includePets bool) (ownerDomain.Owner, error)
	GetByPhone(ctx context.Context, phone string) (ownerDomain.Owner, error)
	Delete(ctx context.Context, id int) error

	//GetByName(ctx context.Context, name string) ([]ownerDomain.Owner, error)
	//ListActiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	//ListInactiveOwners(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	//ListOwnersWithPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)
	//ListOwnersWithoutPets(ctx context.Context, limit, offset int) ([]ownerDomain.Owner, error)

	//CountOwners(ctx context.Context) (int64, error)
	//CountActiveOwners(ctx context.Context) (int64, error)
	//CountInactiveOwners(ctx context.Context) (int64, error)

	//ActivateOwner(ctx context.Context, id int) error
	//DeactivateOwner(ctx context.Context, id int) error

	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id int) (bool, error)
}
