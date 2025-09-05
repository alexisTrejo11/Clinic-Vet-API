package persistence

import (
	"fmt"

	dberr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/database"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

const (
	ErrMsgGetOwner           = "failed to get owner"
	ErrMsgGetOwnerByPhone    = "failed to get owner by phone"
	ErrMsgListOwners         = "failed to list owners"
	ErrMsgCreateOwner        = "failed to create owner"
	ErrMsgUpdateOwner        = "failed to update owner"
	ErrMsgDeleteOwner        = "failed to delete owner"
	ErrMsgActivateOwner      = "failed to activate owner"
	ErrMsgDeactivateOwner    = "failed to deactivate owner"
	ErrMsgCheckExistsByPhone = "failed to check if owner exists by phone"
	ErrMsgCheckExistsByID    = "failed to check if owner exists by ID"
	ErrMsgListPetsForOwner   = "failed to list pets for owner"
	ErrMsgConvertToDomain    = "failed to convert to domain entity"
	ErrMsgCountOwners        = "failed to count owners"
	ErrMsgCreateOwnerID      = "failed to create owner ID from result"
)

const (
	OpSelect = "select"
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
	OpCount  = "count"

	TableOwners = "owners"
	DriverSQL   = "sql"

	// Goroutine timeout for concurrent operations
	ConcurrentOpTimeout = 1
)

// calculateOffset computes the database offset for pagination
func (r *SqlcOwnerRepository) calculateOffset(pagination page.PageInput) int32 {
	return int32((pagination.PageNumber - 1) * pagination.PageSize)
}

// dbError creates a standardized database operation error
func (r *SqlcOwnerRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableOwners, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError creates a standardized entity not found error
func (r *SqlcOwnerRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableOwners, DriverSQL)
}

// wrapConversionError wraps domain conversion errors
func (r *SqlcOwnerRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}
