package repositoryimpl

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

// Database operation constants
const (
	OpSelect = "select"
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
	OpCount  = "count"
	OpSearch = "search"

	TableMedicalSession = "medical_sessions"
	DriverSQL           = "sql"
)

// Error messages
const (
	ErrMsgGetMedicalSession    = "failed to get medical history"
	ErrMsgSearchMedicalSession = "failed to search medical history"
	ErrMsgListByVet            = "failed to list medical history by veterinarian"
	ErrMsgListByPet            = "failed to list medical history by pet"
	ErrMsgListBycustomer       = "failed to list medical history by customer"
	ErrMsgCreateMedicalSession = "failed to create medical history"
	ErrMsgUpdateMedicalSession = "failed to update medical history"
	ErrMsgDeleteMedicalSession = "failed to delete medical history"
	ErrMsgConvertToDomain      = "failed to convert to domain entity"
	ErrMsgInvalidSearchParams  = "invalid search parameters type"
	ErrMsgCountMedicalSession  = "failed to count medical history records"
)

func (r *SQLCMedSessionRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableMedicalSession, DriverSQL, fmt.Errorf("%s: %v", message, err))
}

func (r *SQLCMedSessionRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableMedicalSession, DriverSQL)
}
