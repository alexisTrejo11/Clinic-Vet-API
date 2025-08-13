package ownerDomain

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner Owner) error
	List(ctx context.Context, pagination page.PageData) ([]Owner, error)
	GetById(ctx context.Context, id int) (Owner, error)
	GetByPhone(ctx context.Context, phone string) (Owner, error)
	Delete(ctx context.Context, id int) error

	//ListByName(ctx context.Context, name string) ([]Owner, error)
	//ListActiveOwners(ctx context.Context, limit, offset int) ([]Owner, error)
	//ListInactiveOwners(ctx context.Context, limit, offset int) ([]Owner, error)
	//ListOwnersWithPets(ctx context.Context, limit, offset int) ([]Owner, error)
	//ListOwnersWithoutPets(ctx context.Context, limit, offset int) ([]Owner, error)

	//ActivateOwner(ctx context.Context, id int) error
	//DeactivateOwner(ctx context.Context, id int) error

	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id int) (bool, error)
}
