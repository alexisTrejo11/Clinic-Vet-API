package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/owner"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner *owner.Owner) error
	Search(ctx context.Context, search any) (page.Page[[]owner.Owner], error)
	GetByID(ctx context.Context, id valueobject.OwnerID) (owner.Owner, error)
	GetByPhone(ctx context.Context, phone string) (owner.Owner, error)
	SoftDelete(ctx context.Context, id valueobject.OwnerID) error

	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id valueobject.OwnerID) (bool, error)
}
