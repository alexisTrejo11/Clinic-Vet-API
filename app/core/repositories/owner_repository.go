package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type OwnerRepository interface {
	Save(ctx context.Context, owner *entity.Owner) error
	List(ctx context.Context, pagination page.PageData) (page.Page[[]entity.Owner], error)
	GetById(ctx context.Context, id int) (entity.Owner, error)
	GetByPhone(ctx context.Context, phone string) (entity.Owner, error)
	SoftDelete(ctx context.Context, id int) error

	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id int) (bool, error)
}
