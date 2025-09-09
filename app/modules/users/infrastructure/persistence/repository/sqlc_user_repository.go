// Package repositoryimpl implements the UserRepository interface using SQLC for database operations.
package repositoryimpl

import (
	"context"

	u "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"

	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCUserRepository struct {
	queries *sqlc.Queries
}

func NewSQLCUserRepository(queries *sqlc.Queries) repository.UserRepository {
	return &SQLCUserRepository{
		queries: queries,
	}
}

func (r *SQLCUserRepository) GetByID(ctx context.Context, id valueobject.UserID) (u.User, error) {
	sqlRow, err := r.queries.GetUserByID(ctx, int32(id.Value()))
	if err != nil {
		return u.User{}, err
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return u.User{}, err
	}

	return *user, nil
}

func (r *SQLCUserRepository) GetByEmail(ctx context.Context, email string) (u.User, error) {
	sqlRow, err := r.queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		return u.User{}, err
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return u.User{}, err
	}

	return *user, nil
}

func (r *SQLCUserRepository) GetByPhone(ctx context.Context, phone string) (u.User, error) {
	sqlRow, err := r.queries.GetUserByPhoneNumber(ctx, pgtype.Text{String: phone, Valid: true})
	if err != nil {
		return u.User{}, err
	}

	users, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return u.User{}, err
	}

	return *users, nil
}

func (r *SQLCUserRepository) ListByRole(ctx context.Context, role string, pageInput page.PageInput) (page.Page[[]u.User], error) {
	sqlRows, err := r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		return page.Page[[]u.User]{}, err
	}

	users, err := MapUsersFromSQLC(sqlRows)
	if err != nil {
		return page.Page[[]u.User]{}, err
	}

	return page.Page[[]u.User]{
		Data:     users,
		Metadata: *page.GetPageMetadata(len(users), pageInput),
	}, nil
}

func (r *SQLCUserRepository) Search(ctx context.Context, filterParams any, pageInput page.PageInput) (page.Page[[]u.User], error) {
	sqlRows, err := r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		return page.Page[[]u.User]{}, err
	}

	users, err := MapUsersFromSQLC(sqlRows)
	if err != nil {
		return page.Page[[]u.User]{}, err
	}

	return page.Page[[]u.User]{
		Data:     users,
		Metadata: *page.GetPageMetadata(len(users), pageInput),
	}, nil
}

func (r *SQLCUserRepository) ExistsByID(ctx context.Context, id valueobject.UserID) (bool, error) {
	exists, err := r.queries.ExistsUserByID(ctx, int32(id.Value()))
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *SQLCUserRepository) UpdateLastLogin(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.UpdateUserLastLogin(ctx, int32(id.Value())); err != nil {
		return err
	}
	return nil
}

func (r *SQLCUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := r.queries.ExistsUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (r *SQLCUserRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	exist, err := r.queries.ExistsUserByPhoneNumber(ctx, pgtype.Text{String: phone, Valid: true})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func (r *SQLCUserRepository) Save(ctx context.Context, user *u.User) error {
	if user.ID().IsZero() {
		return r.create(ctx, user)
	}

	return r.update(ctx, user)
}

func (r *SQLCUserRepository) Delete(ctx context.Context, id valueobject.UserID, softDelete bool) error {
	if softDelete {
		return r.softDelete(ctx, id.Value())
	}
	return r.hardDelete(ctx, id.Value())
}

func (r *SQLCUserRepository) softDelete(ctx context.Context, id int) error {
	if err := r.queries.SoftDeleteUser(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (r *SQLCUserRepository) hardDelete(ctx context.Context, id int) error {
	if err := r.queries.HardDeleteUser(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (r *SQLCUserRepository) create(ctx context.Context, user *u.User) error {
	_, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email:       pgtype.Text{String: user.Email().String(), Valid: true},
		PhoneNumber: pgtype.Text{String: user.PhoneNumber().String(), Valid: true},
		Password:    pgtype.Text{String: user.Password(), Valid: true},
		Role:        models.UserRole(user.Role()),
		ProfileID:   pgtype.Int4{Valid: false},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLCUserRepository) update(ctx context.Context, user *u.User) error {
	_, err := r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:          int32(user.ID().Value()),
		Email:       pgtype.Text{String: user.Email().String(), Valid: true},
		PhoneNumber: pgtype.Text{String: user.PhoneNumber().String(), Valid: true},
		Password:    pgtype.Text{String: user.Password(), Valid: true},
		Role:        models.UserRole(user.Role()),
	})
	return err
}
