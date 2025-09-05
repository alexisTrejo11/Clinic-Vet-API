package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type UserRepository interface {
	Save(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int, softDelete bool) error

	GetByID(ctx context.Context, id int) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	GetByPhone(ctx context.Context, phone string) (entity.User, error)
	ListByRole(ctx context.Context, role string, pageInput page.PageInput) (page.Page[[]entity.User], error)
	Search(ctx context.Context, filterParams any, pageInput page.PageInput) (page.Page[[]entity.User], error)

	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id int) (bool, error)
	UpdateLastLogin(ctx context.Context, id int) error
}

type ProfileRepository interface {
	GetByUserID(ctx context.Context, userID int) (entity.Profile, error)
	Create(ctx context.Context, profile *entity.Profile) error
	Update(ctx context.Context, profile *entity.Profile) error
	Delete(ctx context.Context, userID int) error
}
