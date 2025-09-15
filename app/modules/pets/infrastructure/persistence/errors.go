package persistence

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

// Constantes para el repositorio de mascotas
const (
	TablePets = "pets"

	// Mensajes de error específicos
	ErrMsgGetPet             = "failed to get pet"
	ErrMsgGetPetByOwnerID    = "failed to get pet by owner ID"
	ErrMsgListPets           = "failed to list pets"
	ErrMsgSearchPets         = "failed to search pets"
	ErrMsgCreatePet          = "failed to create pet"
	ErrMsgUpdatePet          = "failed to update pet"
	ErrMsgDeletePet          = "failed to delete pet"
	ErrMsgCheckPetExists     = "failed to check if pet exists"
	ErrMsgConvertPetToDomain = "failed to convert pet to domain entity"

	OpSelect = "SELECT"
	OpInsert = "INSERT"
	OpUpdate = "UPDATE"
	OpDelete = "DELETE"

	DriverSQL = "sqlc"
)

// dbError crea un error estandarizado de operación de base de datos
func (r *SqlcPetRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TablePets, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError crea un error estandarizado de entidad no encontrada
func (r *SqlcPetRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TablePets, DriverSQL)
}

// wrapConversionError envuelve errores de conversión de dominio
func (r *SqlcPetRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertPetToDomain, err)
}
