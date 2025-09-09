package persistence

import (
	"fmt"

	dberr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/database"
)

const (
	TableVeterinarians = "veterinarians"
	OpSelect           = "select"
	OpInsert           = "insert"
	OpUpdate           = "update"
	OpDelete           = "delete"
	OpCount            = "count"

	ErrMsgGetVet             = "failed to get veterinarian"
	ErrMsgGetVetByUserID     = "failed to get veterinarian by user ID"
	ErrMsgListVets           = "failed to list veterinarians"
	ErrMsgCreateVet          = "failed to create veterinarian"
	ErrMsgUpdateVet          = "failed to update veterinarian"
	ErrMsgSoftDeleteVet      = "failed to soft delete veterinarian"
	ErrMsgCheckVetExists     = "failed to check if veterinarian exists"
	ErrMsgConvertVetToDomain = "failed to convert veterinarian to domain entity"
)

// dbError crea un error estandarizado de operación de base de datos
func (r *SqlcVetRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableVeterinarians, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError crea un error estandarizado de entidad no encontrada
func (r *SqlcVetRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableVeterinarians, DriverSQL)
}

// wrapConversionError envuelve errores de conversión de dominio
func (r *SqlcVetRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertVetToDomain, err)
}
