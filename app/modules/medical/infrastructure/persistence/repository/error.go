package repositoryimpl

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
	"clinic-vet-api/app/shared/page"

	"github.com/jackc/pgx/v5/pgtype"
)

// Database operation constants
const (
	OpSelect = "select"
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
	OpCount  = "count"
	OpSearch = "search"

	TableMedicalHistory = "medical_history"
	DriverSQL           = "sql"
)

// Error messages
const (
	ErrMsgGetMedicalHistory    = "failed to get medical history"
	ErrMsgSearchMedicalHistory = "failed to search medical history"
	ErrMsgListByVet            = "failed to list medical history by veterinarian"
	ErrMsgListByPet            = "failed to list medical history by pet"
	ErrMsgListByOwner          = "failed to list medical history by owner"
	ErrMsgCreateMedicalHistory = "failed to create medical history"
	ErrMsgUpdateMedicalHistory = "failed to update medical history"
	ErrMsgDeleteMedicalHistory = "failed to delete medical history"
	ErrMsgConvertToDomain      = "failed to convert to domain entity"
	ErrMsgInvalidSearchParams  = "invalid search parameters type"
	ErrMsgCountMedicalHistory  = "failed to count medical history records"
)

func (r *SQLCMedHistRepository) calculateOffset(pageData page.PageInput) int32 {
	return int32(pageData.PageSize * (pageData.Page - 1))
}

func (r *SQLCMedHistRepository) buildNotesParam(notes *string) pgtype.Text {
	if notes != nil {
		return pgtype.Text{String: *notes, Valid: true}
	}
	return pgtype.Text{Valid: false}
}

func (r *SQLCMedHistRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableMedicalHistory, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

func (r *SQLCMedHistRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableMedicalHistory, DriverSQL)
}

func (r *SQLCMedHistRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}
