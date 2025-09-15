// Package repositoryimpl implements the UserRepository interface using SQLC for database operations.
package repositoryimpl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	u "clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/core/domain/valueobject"
	repository "clinic-vet-api/app/core/repositories"
	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
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

func (r *SQLCUserRepository) GetByID(ctx context.Context, id valueobject.UserID) (u.User, error) {
	sqlRow, err := r.queries.GetUserByID(ctx, int32(id.Value()))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("id", id.String())
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgGetUser, id.Value()), err)
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return u.User{}, r.wrapConversionError(err)
	}

	return *user, nil
}

func (r *SQLCUserRepository) GetByEmail(ctx context.Context, email string) (u.User, error) {
	sqlRow, err := r.queries.GetUserByEmail(ctx, pgtype.Text{String: email, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("email", email)
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with email %s", ErrMsgGetUserByEmail, email), err)
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return u.User{}, r.wrapConversionError(err)
	}

	return *user, nil
}

func (r *SQLCUserRepository) GetByPhone(ctx context.Context, phone string) (u.User, error) {
	sqlRow, err := r.queries.GetUserByPhoneNumber(ctx, pgtype.Text{String: phone, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u.User{}, r.notFoundError("phone", phone)
		}
		return u.User{}, r.dbError(OpSelect, fmt.Sprintf("%s with phone %s", ErrMsgGetUserByPhone, phone), err)
	}

	user, err := MapUserFromSQLC(sqlRow)
	if err != nil {
		return u.User{}, r.wrapConversionError(err)
	}

	return *user, nil
}

func (r *SQLCUserRepository) ListByRole(ctx context.Context, role string, pageInput page.PageInput) (page.Page[[]u.User], error) {
	// Implementar paginación correcta
	limit := int32(pageInput.PageSize)
	offset := int32((pageInput.PageNumber - 1) * pageInput.PageSize)

	sqlRows, err := r.queries.ListUsersByRole(ctx, sqlc.ListUsersByRoleParams{
		Role:   models.UserRole(role),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return page.Page[[]u.User]{}, r.dbError(OpSelect, fmt.Sprintf("%s with role %s", ErrMsgListUsers, role), err)
	}

	users, err := MapUsersFromSQLC(sqlRows)
	if err != nil {
		return page.Page[[]u.User]{}, r.wrapConversionError(err)
	}

	totalCount, err := r.queries.CountUsersByRole(ctx, models.UserRole(role))
	if err != nil {
		return page.Page[[]u.User]{}, r.dbError(OpCount, fmt.Sprintf("failed to count users with role %s", role), err)
	}

	pageMetadata := page.GetPageMetadata(int(totalCount), pageInput)
	return page.NewPage(users, *pageMetadata), nil
}

func (r *SQLCUserRepository) Search(ctx context.Context, filterParams any, pageInput page.PageInput) (page.Page[[]u.User], error) {
	// Implementar búsqueda con filtros reales
	// Por ahora, listamos todos los usuarios con paginación
	limit := int32(pageInput.PageSize)
	offset := int32((pageInput.PageNumber - 1) * pageInput.PageSize)

	sqlRows, err := r.queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return page.Page[[]u.User]{}, r.dbError(OpSelect, ErrMsgSearchUsers, err)
	}

	users, err := MapUsersFromSQLC(sqlRows)
	if err != nil {
		return page.Page[[]u.User]{}, r.wrapConversionError(err)
	}

	pageMetadata := page.GetPageMetadata(1, pageInput)
	return page.NewPage(users, *pageMetadata), nil
}

func (r *SQLCUserRepository) ExistsByID(ctx context.Context, id valueobject.UserID) (bool, error) {
	exists, err := r.queries.ExistsUserByID(ctx, int32(id.Value()))
	if err != nil {
		return false, r.dbError(OpSelect, fmt.Sprintf("%s with ID %d", ErrMsgCheckUserExists, id.Value()), err)
	}
	return exists, nil
}

func (r *SQLCUserRepository) UpdateLastLogin(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.UpdateUserLastLogin(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s for user ID %d", ErrMsgUpdateLastLogin, id.Value()), err)
	}
	return nil
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

func (r *SQLCUserRepository) Save(ctx context.Context, user *u.User) error {
	if user.ID().IsZero() {
		return r.create(ctx, user)
	}
	return r.update(ctx, user)
}

func (r *SQLCUserRepository) Delete(ctx context.Context, id valueobject.UserID, softDelete bool) error {
	if softDelete {
		return r.softDelete(ctx, id)
	}
	return r.hardDelete(ctx, id)
}

func (r *SQLCUserRepository) softDelete(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.SoftDeleteUser(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgSoftDeleteUser, id.Value()), err)
	}
	return nil
}

func (r *SQLCUserRepository) hardDelete(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.HardDeleteUser(ctx, int32(id.Value())); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgHardDeleteUser, id.Value()), err)
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
		return r.dbError(OpInsert, ErrMsgCreateUser, err)
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
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s with ID %d", ErrMsgUpdateUser, user.ID().Value()), err)
	}
	return nil
}

// dbError crea un error estandarizado de operación de base de datos
func (r *SQLCUserRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableUsers, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError crea un error estandarizado de entidad no encontrada
func (r *SQLCUserRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableUsers, DriverSQL)
}

// wrapConversionError envuelve errores de conversión de dominio
func (r *SQLCUserRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertUserToDomain, err)
}
