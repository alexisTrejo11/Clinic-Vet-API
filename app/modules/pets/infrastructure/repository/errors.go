package repository

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

// Constantes para el repositorio de mascotas
const (
	TablePets = "pets"

	// Mensajes de error específicos
	ErrMsgFindPet             = "failed to get pet"
	ErrMsgFindPetByCustomerID = "failed to get pet by customer ID"
	ErrMsgFindPets            = "failed to list pets"
	ErrMsgSearchPets          = "failed to search pets"
	ErrMsgCreatePet           = "failed to create pet"
	ErrMsgUpdatePet           = "failed to update pet"
	ErrMsgDeletePet           = "failed to delete pet"
	ErrMsgCheckPetExists      = "failed to check if pet exists"
	ErrMsgConvertPetToDomain  = "failed to convert pet to domain entity"

	OpSelect = "SELECT"
	OpInsert = "INSERT"
	OpUpdate = "UPDATE"
	OpDelete = "DELETE"

	DriverSQL = "sqlc"
)

func (r *SqlcPetRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TablePets, DriverSQL, fmt.Errorf("%s: %v", message, err))
}

func (r *SqlcPetRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TablePets, DriverSQL)
}

func (r *SqlcPetRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertPetToDomain, err)
}
