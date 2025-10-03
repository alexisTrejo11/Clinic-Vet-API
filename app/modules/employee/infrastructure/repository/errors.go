package repository

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

const (
	TableEmployees = "veterinarians"
	OpSelect       = "select"
	OpInsert       = "insert"
	OpUpdate       = "update"
	OpDelete       = "delete"
	OpCount        = "count"
	DriverSQL      = "sqlc"

	ErrMsgGetEmployee             = "failed to get veterinarian"
	ErrMsgGetEmployeeByUserID     = "failed to get veterinarian by user ID"
	ErrMsgListEmployees           = "failed to list veterinarians"
	ErrMsgCreateEmployee          = "failed to create veterinarian"
	ErrMsgUpdateEmployee          = "failed to update veterinarian"
	ErrMsgSoftDeleteEmployee      = "failed to soft delete veterinarian"
	ErrMsgCheckEmployeeExists     = "failed to check if veterinarian exists"
	ErrMsgConvertEmployeeToDomain = "failed to convert veterinarian to domain entity"
)

func (r *SqlcEmployeeRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableEmployees, DriverSQL, fmt.Errorf("%s: %v", message, err))
}

func (r *SqlcEmployeeRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableEmployees, DriverSQL)
}
