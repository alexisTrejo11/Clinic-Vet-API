package userRepository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, id int, softDelete bool) error

	GetById(ctx context.Context, id int) (*user.User, error)
	GetByIdWithProfile(ctx context.Context, id int) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
	GetByPhone(ctx context.Context, phone string) (*user.User, error)
	ListByRole(ctx context.Context, role string, pageInput page.PageData) (page.Page[[]user.User], error)
	Search(ctx context.Context, filterParams map[string]interface{}, pageInput page.PageData) (page.Page[[]user.User], error)

	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	UpdateLastLogin(ctx context.Context, id int) error
	UpdateProfile(ctx context.Context, id int, profile user.Profile) error
}
