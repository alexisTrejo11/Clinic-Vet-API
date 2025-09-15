package repository

import (
	"context"
	"time"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/entity/user/profile"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

type UserRepository interface {
	FindSpecification(ctx context.Context, spec specification.UserSpecification) (page.Page[user.User], error)
	FindByID(ctx context.Context, id valueobject.UserID) (user.User, error)
	FindByEmail(ctx context.Context, email string) (user.User, error)
	FindByPhone(ctx context.Context, phone string) (user.User, error)

	FindAll(ctx context.Context, pageInput page.PageInput) (page.Page[user.User], error)
	FindByRole(ctx context.Context, role string, pageInput page.PageInput) (page.Page[user.User], error)
	FindActive(ctx context.Context, pageInput page.PageInput) (page.Page[user.User], error)
	FindInactive(ctx context.Context, pageInput page.PageInput) (page.Page[user.User], error)

	FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (user.User, error)
	FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (user.User, error)
	FindRecentlyLoggedIn(ctx context.Context, since time.Time, pageInput page.PageInput) (page.Page[user.User], error)

	ExistsByID(ctx context.Context, id valueobject.UserID) (bool, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
	ExistsByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (bool, error)
	ExistsByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (bool, error)

	Save(ctx context.Context, user *user.User) error
	Update(ctx context.Context, user *user.User) error
	SoftDelete(ctx context.Context, id valueobject.UserID) error
	HardDelete(ctx context.Context, id valueobject.UserID) error

	UpdateLastLogin(ctx context.Context, id valueobject.UserID) error
	UpdatePassword(ctx context.Context, id valueobject.UserID, hashedPassword string) error
	UpdateStatus(ctx context.Context, id valueobject.UserID, status enum.UserStatus) error

	CountAll(ctx context.Context) (int64, error)
	CountByRole(ctx context.Context, role string) (int64, error)
	CountByStatus(ctx context.Context, status enum.UserStatus) (int64, error)
	CountActive(ctx context.Context) (int64, error)
}

type ProfileRepository interface {
	GetByUserID(ctx context.Context, userID valueobject.UserID) (profile.Profile, error)
	Create(ctx context.Context, profile *profile.Profile) error
	Update(ctx context.Context, profile *profile.Profile) error
	DeleteByUserID(ctx context.Context, userID valueobject.UserID) error
}
