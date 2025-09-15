package repository

import (
	"clinic-vet-api/app/shared/page"
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

const (
	OpSelect = "select"
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
	OpCount  = "count"
	OpSearch = "search"

	TableAppts = "appointments"
	DriverSQL  = "sql"
)

const (
	ErrMsgListAppts       = "failed to list appointments"
	ErrMsgGetAppt         = "failed to get appointment"
	ErrMsgSearchAppts     = "failed to search appointments"
	ErrMsgCountAppts      = "failed to count appointments"
	ErrMsgCreateAppt      = "failed to create appointment"
	ErrMsgUpdateAppt      = "failed to update appointment"
	ErrMsgDeleteAppt      = "failed to delete appointment"
	ErrMsgConvertToDomain = "failed to convert to domain entity"
	ErrMsgNotFound        = "appointment not found"
)

// calculateOffset computes the database offset for pagination
func (r *SqlcAppointmentRepository) calculateOffset(pageInput page.PageInput) int32 {
	return int32(pageInput.Page-1) * int32(pageInput.PageSize)
}

// dbError creates a standardized database operation error
func (r *SqlcAppointmentRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableAppts, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError creates a standardized entity not found error
func (r *SqlcAppointmentRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableAppts, DriverSQL)
}

// wrapConversionError wraps domain conversion errors
func (r *SqlcAppointmentRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}
