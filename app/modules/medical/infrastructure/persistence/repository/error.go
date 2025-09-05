package persistence

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/dto"
	dberr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/database"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
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

// calculateOffset computes the database offset for pagination
func (r *SQLCMedHistRepository) calculateOffset(pageData page.PageInput) int32 {
	return int32(pageData.PageSize * (pageData.PageNumber - 1))
}

// applyManualPagination applies pagination logic in memory (temporary solution)
func (r *SQLCMedHistRepository) applyManualPagination(items []entity.MedicalHistory, pagination page.PageInput) ([]entity.MedicalHistory, int) {
	totalCount := len(items)
	startIndex := (pagination.PageNumber - 1) * pagination.PageSize
	endIndex := startIndex + pagination.PageSize

	if startIndex >= totalCount {
		return []entity.MedicalHistory{}, totalCount
	}

	if endIndex > totalCount {
		endIndex = totalCount
	}

	return items[startIndex:endIndex], totalCount
}

// buildNotesParam creates a pgtype.Text parameter for notes field
func (r *SQLCMedHistRepository) buildNotesParam(notes *string) pgtype.Text {
	if notes != nil {
		return pgtype.Text{String: *notes, Valid: true}
	}
	return pgtype.Text{Valid: false}
}

// getSearchResultCount gets the total count for search results
// TODO: Implement with actual count query
func (r *SQLCMedHistRepository) getSearchResultCount(ctx context.Context, searchParam dto.MedHistSearchParams) (int, error) {
	// Placeholder - should use actual count query with search parameters
	// For now, return 0 to avoid errors
	return 0, nil
}

// dbError creates a standardized database operation error
func (r *SQLCMedHistRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableMedicalHistory, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError creates a standardized entity not found error
func (r *SQLCMedHistRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableMedicalHistory, DriverSQL)
}

// wrapConversionError wraps domain conversion errors
func (r *SQLCMedHistRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}

// invalidParamsError creates an error for invalid parameters
func (r *SQLCMedHistRepository) invalidParamsError(message string) error {
	return fmt.Errorf("%s: expected dto.MedHistSearchParams", message)
}
