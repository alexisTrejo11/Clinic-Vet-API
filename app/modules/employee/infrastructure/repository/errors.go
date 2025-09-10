package repositoryimpl

import (
	"fmt"

	dberr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/database"
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

// dbError crea un error estandarizado de operación de base de datos
func (r *SqlcEmployeeRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableEmployees, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError crea un error estandarizado de entidad no encontrada
func (r *SqlcEmployeeRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableEmployees, DriverSQL)
}

// wrapConversionError envuelve errores de conversión de dominio
func (r *SqlcEmployeeRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertEmployeeToDomain, err)
}
