package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user/profile"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type UserRepository interface {
	Save(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, id valueobject.UserID, softDelete bool) error

	GetByID(ctx context.Context, id valueobject.UserID) (user.User, error)
	GetByEmail(ctx context.Context, email string) (user.User, error)
	GetByPhone(ctx context.Context, phone string) (user.User, error)
	ListByRole(ctx context.Context, role string, pageInput page.PageInput) (page.Page[[]user.User], error)
	Search(ctx context.Context, filterParams any, pageInput page.PageInput) (page.Page[[]user.User], error)

	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByID(ctx context.Context, id valueobject.UserID) (bool, error)
	ExistsByEmployeeID(ctx context.Context, id valueobject.VetID) (bool, error)
	UpdateLastLogin(ctx context.Context, id valueobject.UserID) error
}

type ProfileRepository interface {
	GetByUserID(ctx context.Context, userID valueobject.UserID) (profile.Profile, error)
	Create(ctx context.Context, profile *profile.Profile) error
	Update(ctx context.Context, profile *profile.Profile) error
	DeleteByUserID(ctx context.Context, userID valueobject.UserID) error
}
