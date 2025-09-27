// Package repositoryimpl implements the UserRepository interface using SQLC for database operations.
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"

	"clinic-vet-api/app/shared/page"

	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

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

func (r *SQLCUserRepository) FindByOAuthProvider(ctx context.Context, provider string, providerID string) (u.User, error) {
	panic("unimplemented")
}

func (r *SQLCUserRepository) FindByID(ctx context.Context, id valueobject.UserID) (u.User, error) {
	sqlRow, err := r.queries.FindUserByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("id", id.String())
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgFindUser, id.Value()), err)
	}

	return sqlcRowWithJoinToEntity(sqlRow), nil
}

func (r *SQLCUserRepository) FindByEmail(ctx context.Context, email string) (u.User, error) {
	sqlRow, err := r.queries.FindUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("email", email)
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with email %s", ErrMsgFindUserByEmail, email), err)
	}

	return sqlcRowWithJoinToEntity(sqlRow), nil
}

func (r *SQLCUserRepository) FindByPhone(ctx context.Context, phone string) (u.User, error) {
	sqlRow, err := r.queries.FindUserByPhoneNumber(ctx, pgtype.Text{String: phone, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("phone", phone)
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with phone %s", ErrMsgFindUserByPhone, phone), err)
	}

	return sqlcRowWithJoinToEntity(sqlRow), nil
}

func (r *SQLCUserRepository) FindByRole(ctx context.Context, role string, pagination page.PaginationRequest) (page.Page[u.User], error) {
	sqlRows, err := r.queries.FindUsersByRole(ctx, sqlc.FindUsersByRoleParams{
		Role:   models.UserRole(role),
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError(OpSelect, fmt.Sprintf("%s with role %s", ErrMsgFindUsers, role), err)
	}

	users, err := sqlcRowsToEntities(sqlRows)
	if err != nil {
		return page.Page[u.User]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountUsersByRole(ctx, models.UserRole(role))
	if err != nil {
		return page.Page[u.User]{}, r.dbError(OpCount, fmt.Sprintf("failed to count users with role %s", role), err)
	}

	return page.NewPage(users, int(totalCount), pagination), nil
}

func (r *SQLCUserRepository) FindActive(ctx context.Context, pagination page.PaginationRequest) (page.Page[u.User], error) {
	userRows, err := r.queries.FindActiveUsers(ctx, sqlc.FindActiveUsersParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to find active users", err)
	}

	total, err := r.queries.CountActiveUsers(ctx)
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to count active users", err)
	}

	users, err := sqlcRowsToEntities(userRows)
	if err != nil {
		return page.Page[u.User]{}, r.wrapConversionError(err)
	}

	return page.NewPage(users, int(total), pagination), nil
}

func (r *SQLCUserRepository) FindAll(ctx context.Context, pagination page.PaginationRequest) (page.Page[u.User], error) {
	userRows, err := r.queries.FindAllUsers(ctx, sqlc.FindAllUsersParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to find all users", err)
	}

	total, err := r.queries.CountAllUsers(ctx)
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to count all users", err)
	}

	users, err := sqlcRowsToEntities(userRows)
	if err != nil {
		return page.Page[u.User]{}, r.wrapConversionError(err)
	}

	return page.NewPage(users, int(total), pagination), nil
}

func (r *SQLCUserRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (u.User, error) {
	userRow, err := r.queries.FindUserByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("customer_id", customerID.String())
		}
		return u.User{}, r.dbError("select", fmt.Sprintf("failed to find user by customer ID %d", customerID.Value()), err)
	}

	return sqlcRowToEntity(userRow), nil
}

func (r *SQLCUserRepository) FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (u.User, error) {
	userRow, err := r.queries.FindUserByEmployeeID(ctx, int32(employeeID.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("employee_id", employeeID.String())
		}
		return u.User{}, r.dbError("select", fmt.Sprintf("failed to find user by employee ID %d", employeeID.Value()), err)
	}

	return sqlcRowToEntity(userRow), nil
}

func (r *SQLCUserRepository) FindInactive(ctx context.Context, pagination page.PaginationRequest) (page.Page[u.User], error) {
	userRows, err := r.queries.FindInactiveUsers(ctx, sqlc.FindInactiveUsersParams{
		Limit:  int32(pagination.Limit()),
		Offset: int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to find inactive users", err)
	}

	activeCount, err := r.queries.CountActiveUsers(ctx)
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to count active users", err)
	}
	totalCount, err := r.queries.CountAllUsers(ctx)
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to count all users", err)
	}
	total := totalCount - activeCount

	users, err := sqlcRowsToEntities(userRows)
	if err != nil {
		return page.Page[u.User]{}, r.wrapConversionError(err)
	}

	return page.NewPage(users, int(total), pagination), nil
}

func (r *SQLCUserRepository) FindRecentlyLoggedIn(ctx context.Context, since time.Time, pagination page.PaginationRequest) (page.Page[u.User], error) {
	userRows, err := r.queries.FindRecentlyLoggedInUsers(ctx, sqlc.FindRecentlyLoggedInUsersParams{
		LastLogin: pgtype.Timestamptz{Time: since, Valid: true},
		Limit:     int32(pagination.Limit()),
		Offset:    int32(pagination.Offset()),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to find recently logged in users", err)
	}

	users, err := sqlcRowsToEntities(userRows)
	if err != nil {
		return page.Page[u.User]{}, r.wrapConversionError(err)
	}

	return page.NewPage(users, 0, pagination), nil
}
func (r *SQLCUserRepository) FindSpecification(ctx context.Context, spec specification.UserSpecification) (page.Page[u.User], error) {
	return page.Page[u.User]{}, fmt.Errorf("method FindSpecification not implemented yet")
}

func (r *SQLCUserRepository) UpdatePassword(ctx context.Context, id valueobject.UserID, hashedPassword string) error {
	err := r.queries.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		ID:       int32(id.Value()),
		Password: pgtype.Text{String: hashedPassword, Valid: true},
	})
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update password for user ID %d", id.Value()), err)
	}
	return nil
}

func (r *SQLCUserRepository) UpdateStatus(ctx context.Context, id valueobject.UserID, status enum.UserStatus) error {
	err := r.queries.UpdateUserStatus(ctx, sqlc.UpdateUserStatusParams{
		ID:     int32(id.Value()),
		Status: models.UserStatus(status.String()),
	})
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update status for user ID %d", id.Value()), err)
	}
	return nil
}

func (r *SQLCUserRepository) UpdateLastLogin(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.UpdateUserLastLogin(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s for user ID %d", ErrMsgUpdateLastLogin, id.Value()), err)
	}
	return nil
}

func (r *SQLCUserRepository) Save(ctx context.Context, user *u.User) error {
	if user.ID().IsZero() {
		return r.create(ctx, user)
	}
	return r.update(ctx, user)
}

func (r *SQLCUserRepository) SoftDelete(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.SoftDeleteUser(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgSoftDeleteUser, id.Value()), err)
	}
	return nil
}

func (r *SQLCUserRepository) HardDelete(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.HardDeleteUser(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgHardDeleteUser, id.Value()), err)
	}
	return nil
}

func (r *SQLCUserRepository) create(ctx context.Context, user *u.User) error {
	params := entityToCreateParams(user)
	userCreated, err := r.queries.CreateUser(ctx, *params)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateUser, err)
	}

	user.SetID(valueobject.NewUserID(uint(userCreated.ID)))

	return nil
}

func (r *SQLCUserRepository) update(ctx context.Context, user *u.User) error {
	params := entityToUpdateParams(user)
	_, err := r.queries.UpdateUser(ctx, *params)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s with ID %d", ErrMsgUpdateUser, user.ID().Value()), err)
	}
	return nil
}

func (r *SQLCUserRepository) ExistsByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (bool, error) {
	exists, err := r.queries.ExistsUserByCustomerID(ctx, int32(customerID.Value()))
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check user existence by customer ID %d", customerID.Value()), err)
	}
	return exists, nil
}

func (r *SQLCUserRepository) ExistsByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (bool, error) {
	exists, err := r.queries.ExistsUserByEmployeeID(ctx, int32(employeeID.Value()))
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check user existence by employee ID %d", employeeID.Value()), err)
	}
	return exists, nil
}

func (r *SQLCUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := r.queries.ExistsUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with email %s", ErrMsgCheckUserExists, email), err)
	}
	return exists, nil
}

func (r *SQLCUserRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	exist, err := r.queries.ExistsUserByPhoneNumber(ctx, pgtype.Text{String: phone, Valid: true})
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with phone %s", ErrMsgCheckUserExists, phone), err)
	}

	return exist, nil
}

func (r *SQLCUserRepository) ExistsByID(ctx context.Context, id valueobject.UserID) (bool, error) {
	exists, err := r.queries.ExistsUserByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgCheckUserExists, id.Value()), err)
	}
	return exists, nil
}

func (r *SQLCUserRepository) CountActive(ctx context.Context) (int64, error) {
	count, err := r.queries.CountActiveUsers(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count active users", err)
	}
	return count, nil
}

func (r *SQLCUserRepository) CountAll(ctx context.Context) (int64, error) {
	count, err := r.queries.CountAllUsers(ctx)
	if err != nil {
		return 0, r.dbError("select", "failed to count all users", err)
	}
	return count, nil
}

func (r *SQLCUserRepository) CountByRole(ctx context.Context, role string) (int64, error) {
	count, err := r.queries.CountUsersByRole(ctx, models.UserRole(role))
	if err != nil {
		return 0, r.dbError("select", fmt.Sprintf("failed to count users by role %s", role), err)
	}
	return count, nil
}

func (r *SQLCUserRepository) CountByStatus(ctx context.Context, status enum.UserStatus) (int64, error) {
	count, err := r.queries.CountUsersByStatus(ctx, models.UserStatus(status.String()))
	if err != nil {
		return 0, r.dbError("select", fmt.Sprintf("failed to count users by status %s", status), err)
	}
	return count, nil
}
