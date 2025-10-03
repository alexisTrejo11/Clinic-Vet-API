// Package repositoryimpl implements the UserRepository interface using SQLC for database operations.
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"

	"clinic-vet-api/app/shared/mapper"
	"clinic-vet-api/app/shared/page"

	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcUserRepository struct {
	queries *sqlc.Queries
	pgMap   *mapper.SqlcFieldMapper
}

func NewSqlcUserRepository(queries *sqlc.Queries, pgMap *mapper.SqlcFieldMapper) repository.UserRepository {
	return &SqlcUserRepository{
		queries: queries,
		pgMap:   pgMap,
	}
}

func (r *SqlcUserRepository) FindByOAuthProvider(ctx context.Context, provider string, providerID string) (u.User, error) {
	panic("unimplemented")
}

func (r *SqlcUserRepository) FindByID(ctx context.Context, id valueobject.UserID) (u.User, error) {
	sqlRow, err := r.queries.FindUserByID(ctx, id.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("id", id.String())
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgFindUser, id.Value()), err)
	}

	return r.WithJoinToEntity(sqlRow), nil
}

func (r *SqlcUserRepository) FindByEmail(ctx context.Context, email string) (u.User, error) {
	sqlRow, err := r.queries.FindUserByEmail(ctx, r.pgMap.StringToPgText(email))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("email", email)
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with email %s", ErrMsgFindUserByEmail, email), err)
	}

	return r.WithJoinToEntity(sqlRow), nil
}

func (r *SqlcUserRepository) FindByPhone(ctx context.Context, phone string) (u.User, error) {
	sqlRow, err := r.queries.FindUserByPhoneNumber(ctx, pgtype.Text{String: phone, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("phone", phone)
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with phone %s", ErrMsgFindUserByPhone, phone), err)
	}

	return r.WithJoinToEntity(sqlRow), nil
}

func (r *SqlcUserRepository) FindByRole(ctx context.Context, role string, pagination page.PaginationRequest) (page.Page[u.User], error) {
	sqlRows, err := r.queries.FindUsersByRole(ctx, sqlc.FindUsersByRoleParams{
		Role:   models.UserRole(role),
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError(OpSelect, fmt.Sprintf("%s with role %s", ErrMsgFindUsers, role), err)
	}

	totalCount, err := r.queries.CountUsersByRole(ctx, models.UserRole(role))
	if err != nil {
		return page.Page[u.User]{}, r.dbError(OpSelect, fmt.Sprintf("failed to count users with role %s", role), err)
	}

	users := r.ToEntities(sqlRows)
	return page.NewPage(users, totalCount, pagination), nil
}

func (r *SqlcUserRepository) FindActive(ctx context.Context, pagination page.PaginationRequest) (page.Page[u.User], error) {
	userRows, err := r.queries.FindActiveUsers(ctx, sqlc.FindActiveUsersParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to find active users", err)
	}

	total, err := r.queries.CountActiveUsers(ctx)
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to count active users", err)
	}

	users := r.ToEntities(userRows)
	return page.NewPage(users, total, pagination), nil
}

func (r *SqlcUserRepository) FindAll(ctx context.Context, pagination page.PaginationRequest) (page.Page[u.User], error) {
	userRows, err := r.queries.FindAllUsers(ctx, sqlc.FindAllUsersParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
	})
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to find all users", err)
	}

	total, err := r.queries.CountAllUsers(ctx)
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to count all users", err)
	}

	users := r.ToEntities(userRows)
	return page.NewPage(users, total, pagination), nil
}

func (r *SqlcUserRepository) FindByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (u.User, error) {
	userRow, err := r.queries.FindUserByCustomerID(ctx, customerID.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("customer_id", customerID.String())
		}
		return u.User{}, r.dbError("select", fmt.Sprintf("failed to find user by customer ID %d", customerID.Value()), err)
	}

	return r.ToEntity(userRow), nil
}

func (r *SqlcUserRepository) FindByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (u.User, error) {
	userRow, err := r.queries.FindUserByEmployeeID(ctx, employeeID.Int32())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("employee_id", employeeID.String())
		}
		return u.User{}, r.dbError("select", fmt.Sprintf("failed to find user by employee ID %d", employeeID.Value()), err)
	}

	return r.ToEntity(userRow), nil
}

func (r *SqlcUserRepository) FindInactive(ctx context.Context, pagination page.PaginationRequest) (page.Page[u.User], error) {
	userRows, err := r.queries.FindInactiveUsers(ctx, sqlc.FindInactiveUsersParams{
		Limit:  pagination.Limit(),
		Offset: pagination.Offset(),
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

	users := r.ToEntities(userRows)
	return page.NewPage(users, total, pagination), nil
}

func (r *SqlcUserRepository) FindRecentlyLoggedIn(ctx context.Context, since time.Time, pagination page.PaginationRequest) (page.Page[u.User], error) {
	params := sqlc.FindRecentlyLoggedInUsersParams{
		LastLogin: r.pgMap.PgTimestamptz.FromTime(since),
		Limit:     pagination.Limit(),
		Offset:    pagination.Offset(),
	}

	userRows, err := r.queries.FindRecentlyLoggedInUsers(ctx, params)
	if err != nil {
		return page.Page[u.User]{}, r.dbError("select", "failed to find recently logged in users", err)
	}

	users := r.ToEntities(userRows)
	return page.NewPage(users, 0, pagination), nil
}
func (r *SqlcUserRepository) FindSpecification(ctx context.Context, spec specification.UserSpecification) (page.Page[u.User], error) {
	return page.Page[u.User]{}, fmt.Errorf("method FindSpecification not implemented yet")
}

func (r *SqlcUserRepository) ExistsByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (bool, error) {
	exists, err := r.queries.ExistsUserByCustomerID(ctx, customerID.Int32())
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check user existence by customer ID %d", customerID.Value()), err)
	}
	return exists, nil
}

func (r *SqlcUserRepository) ExistsByEmployeeID(ctx context.Context, employeeID valueobject.EmployeeID) (bool, error) {
	exists, err := r.queries.ExistsUserByEmployeeID(ctx, employeeID.Int32())
	if err != nil {
		return false, r.dbError("select", fmt.Sprintf("failed to check user existence by employee ID %d", employeeID.Value()), err)
	}
	return exists, nil
}

func (r *SqlcUserRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := r.queries.ExistsUserByEmail(ctx, r.pgMap.StringToPgText(email))
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with email %s", ErrMsgCheckUserExists, email), err)
	}
	return exists, nil
}

func (r *SqlcUserRepository) ExistsByPhone(ctx context.Context, phone string) (bool, error) {
	exist, err := r.queries.ExistsUserByPhoneNumber(ctx, r.pgMap.StringToPgText(phone))
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with phone %s", ErrMsgCheckUserExists, phone), err)
	}

	return exist, nil
}

func (r *SqlcUserRepository) ExistsByID(ctx context.Context, id valueobject.UserID) (bool, error) {
	exists, err := r.queries.ExistsUserByID(ctx, id.Int32())
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgCheckUserExists, id.Value()), err)
	}
	return exists, nil
}
