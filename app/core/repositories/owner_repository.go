package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner *entity.Owner) error
	List(ctx context.Context, pagination page.PageInput) (page.Page[[]entity.Owner], error)
	GetByID(ctx context.Context, id valueobject.OwnerID) (entity.Owner, error)
	GetByPhone(ctx context.Context, phone string) (entity.Owner, error)
	SoftDelete(ctx context.Context, id valueobject.OwnerID) error

	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id valueobject.OwnerID) (bool, error)
}
