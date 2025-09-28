package repositoryimpl

import (
	u "clinic-vet-api/app/modules/core/domain/entity/user"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"context"
	"fmt"
)

func (r *SqlcUserRepository) Save(ctx context.Context, user *u.User) error {
	if user.ID().IsZero() {
		return r.create(ctx, user)
	}
	return r.update(ctx, user)
}

func (r *SqlcUserRepository) Delete(ctx context.Context, id valueobject.UserID, isHardDelete bool) error {
	if isHardDelete {
		return r.hardDelete(ctx, id)
	}
	return r.softDelete(ctx, id)
}

func (r *SqlcUserRepository) UpdatePassword(ctx context.Context, id valueobject.UserID, hashedPassword string) error {
	params := sqlc.UpdateUserPasswordParams{
		ID:       id.Int32(),
		Password: r.pgMap.PgText.FromString(hashedPassword),
	}

	err := r.queries.UpdateUserPassword(ctx, params)
	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update password for user ID %d", id.Value()), err)
	}

	return nil
}

func (r *SqlcUserRepository) UpdateStatus(ctx context.Context, id valueobject.UserID, status enum.UserStatus) error {
	err := r.queries.UpdateUserStatus(ctx, sqlc.UpdateUserStatusParams{
		ID:     int32(id.Value()),
		Status: models.UserStatus(status.String()),
	})

	if err != nil {
		return r.dbError("update", fmt.Sprintf("failed to update status for user ID %d", id.Value()), err)
	}

	return nil
}

func (r *SqlcUserRepository) UpdateLastLogin(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.UpdateUserLastLogin(ctx, id.Int32()); err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s for user ID %d", ErrMsgUpdateLastLogin, id.Value()), err)
	}
	return nil
}

func (r *SqlcUserRepository) softDelete(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.SoftDeleteUser(ctx, id.Int32()); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgSoftDeleteUser, id.Value()), err)
	}
	return nil
}

func (r *SqlcUserRepository) hardDelete(ctx context.Context, id valueobject.UserID) error {
	if err := r.queries.HardDeleteUser(ctx, id.Int32()); err != nil {
		return r.dbError(OpDelete, fmt.Sprintf("%s with ID %d", ErrMsgHardDeleteUser, id.Value()), err)
	}
	return nil
}

func (r *SqlcUserRepository) create(ctx context.Context, user *u.User) error {
	params := r.toCreateParams(user)
	userCreated, err := r.queries.CreateUser(ctx, *params)
	if err != nil {
		return r.dbError(OpInsert, ErrMsgCreateUser, err)
	}

	user.SetID(valueobject.NewUserID(uint(userCreated.ID)))

	return nil
}

func (r *SqlcUserRepository) update(ctx context.Context, user *u.User) error {
	params := r.toUpdateParams(user)

	_, err := r.queries.UpdateUser(ctx, *params)
	if err != nil {
		return r.dbError(OpUpdate, fmt.Sprintf("%s with ID %d", ErrMsgUpdateUser, user.ID().Value()), err)
	}

	return nil
}
