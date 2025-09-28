package repositoryimpl

import (
	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
	"fmt"
)

const (
	TableUsers = "users"

	ErrMsgFindUser            = "failed to find user"
	ErrMsgFindUserByEmail     = "failed to find user by email"
	ErrMsgFindUserByPhone     = "failed to find user by phone"
	ErrMsgFindUsers           = "failed to list users"
	ErrMsgSearchUsers         = "failed to search users"
	ErrMsgCreateUser          = "failed to create user"
	ErrMsgUpdateUser          = "failed to update user"
	ErrMsgSoftDeleteUser      = "failed to soft delete user"
	ErrMsgHardDeleteUser      = "failed to hard delete user"
	ErrMsgCheckUserExists     = "failed to check if user exists"
	ErrMsgUpdateLastLogin     = "failed to update last login"
	ErrMsgConvertUserToDomain = "failed to convert user to domain entity"

	DriverSQL = "SQLC"
	OpSelect  = "SELECT"
	OpInsert  = "INSERT"
	OpUpdate  = "UPDATE"
	OpDelete  = "DELETE"
	OpCount   = "COUNT"
)

func (r *SqlcUserRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableUsers, DriverSQL, fmt.Errorf("%s: %v", message, err))
}

func (r *SqlcUserRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableUsers, DriverSQL)
}

func (r *SqlcUserRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertUserToDomain, err)
}
