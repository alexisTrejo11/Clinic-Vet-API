package sqlcUserRepo

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/db/models"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SQLCUserRepository struct {
	queries *sqlc.Queries
	SQLCProfileRepository
}

func NewSQLCUserRepository(queries *sqlc.Queries) userRepository.UserRepository {
	return &SQLCUserRepository{
		queries: queries,
	}
}

func (r *SQLCUserRepository) GetById(ctx context.Context, id int) (*user.User, error) {
	sqlRow, err := r.queries.GetUserByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *SQLCUserRepository) GetByIdWithProfile(ctx context.Context, id int) (*user.User, error) {
	sqlRow, err := r.queries.GetUserByID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return nil, err
	}

	profile, err := r.SQLCProfileRepository.GetByUserId(ctx, user.Id().GetValue())
	if err != nil {
		return nil, err
	}

	user.SetProfile(profile)
	return user, nil
}

func (r *SQLCUserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	sqlRow, err := r.queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		return nil, err
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *SQLCUserRepository) GetByPhone(ctx context.Context, phone string) (*user.User, error) {
	sqlRow, err := r.queries.GetUserByPhoneNumber(ctx, pgtype.Text{String: phone, Valid: true})
	if err != nil {
		return nil, err
	}

	users, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *SQLCUserRepository) ListByRole(ctx context.Context, role string, pageInput page.PageData) (page.Page[[]user.User], error) {
	sqlRows, err := r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		return page.Page[[]user.User]{}, err
	}

	users, err := MapUsersFromSQLC(sqlRows)
	if err != nil {
		return page.Page[[]user.User]{}, err
	}

	return page.Page[[]user.User]{
		Data:     users,
		Metadata: *page.GetPageMetadata(len(users), pageInput),
	}, nil

}

func (r *SQLCUserRepository) Search(ctx context.Context, filterParams map[string]interface{}, pageInput page.PageData) (page.Page[[]user.User], error) {
	sqlRows, err := r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil {
		return page.Page[[]user.User]{}, err
	}

	users, err := MapUsersFromSQLC(sqlRows)
	if err != nil {
		return page.Page[[]user.User]{}, err
	}

	return page.Page[[]user.User]{
		Data:     users,
		Metadata: *page.GetPageMetadata(len(users), pageInput),
	}, nil
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

func (r *SQLCUserRepository) Save(ctx context.Context, user *user.User) error {
	if user.Id().GetValue() == 0 {
		return r.create(ctx, user)
	}

	return r.update(ctx, user)
}

func (r *SQLCUserRepository) UpdateLastLogin(ctx context.Context, id int) error {
	if err := r.queries.UpdateUserLastLogin(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (r *SQLCUserRepository) Delete(ctx context.Context, id int, softDelete bool) error {
	if softDelete {
		return r.softDelete(ctx, id, softDelete)
	}

	return r.hardDelete(ctx, id, softDelete)
}

func (r *SQLCUserRepository) UpdateProfile(ctx context.Context, id int, profile user.Profile) error {
	return r.SQLCProfileRepository.Update(ctx, &profile)
}

func (r *SQLCUserRepository) softDelete(ctx context.Context, id int, softDelete bool) error {
	if err := r.queries.SoftDeleteUser(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (r *SQLCUserRepository) hardDelete(ctx context.Context, id int, softDelete bool) error {
	if err := r.queries.HardDeleteUser(ctx, int32(id)); err != nil {
		return err
	}
	return nil
}

func (r *SQLCUserRepository) create(ctx context.Context, user *user.User) error {
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

func (r *SQLCUserRepository) update(ctx context.Context, user *user.User) error {
	_, err := r.queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:          int32(user.Id().GetValue()),
		Email:       pgtype.Text{String: user.Email().String(), Valid: true},
		PhoneNumber: pgtype.Text{String: user.PhoneNumber().String(), Valid: true},
		Password:    pgtype.Text{String: user.Password(), Valid: true},
		Role:        models.UserRole(user.Role()),
	})
	return err
}
