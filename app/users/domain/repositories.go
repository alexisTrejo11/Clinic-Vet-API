package userDomain

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int, softDelete bool) error

	GetById(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByPhone(ctx context.Context, phone string) (*User, error)
	ListByRole(ctx context.Context, role string, pageInput page.PageData) (page.Page[[]User], error)
	Search(ctx context.Context, filterParams map[string]interface{}, pageInput page.PageData) (page.Page[[]User], error)

	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	UpdateLastLogin(ctx context.Context, id int) error
}

type ProfileRepository interface {
	GetByUserId(ctx context.Context, userId int) (Profile, error)
	Create(ctx context.Context, profile *Profile) error
	Update(ctx context.Context, profile *Profile) error
	Delete(ctx context.Context, userId int) error
}
