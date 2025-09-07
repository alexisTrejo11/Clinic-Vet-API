package repositoryimpl

import (
	"fmt"

	dberr "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/infrastructure/database"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

const (
	OpSelect = "select"
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
	OpCount  = "count"
	OpSearch = "search"

	TableAppointments = "appointments"
	DriverSQL         = "sql"
)

const (
	ErrMsgListAppointments   = "failed to list appointments"
	ErrMsgGetAppointment     = "failed to get appointment"
	ErrMsgSearchAppointments = "failed to search appointments"
	ErrMsgCountAppointments  = "failed to count appointments"
	ErrMsgCreateAppointment  = "failed to create appointment"
	ErrMsgUpdateAppointment  = "failed to update appointment"
	ErrMsgDeleteAppointment  = "failed to delete appointment"
	ErrMsgConvertToDomain    = "failed to convert to domain entity"
	ErrMsgNotFound           = "appointment not found"
)

// calculateOffset computes the database offset for pagination
func (r *SQLCAppointmentRepository) calculateOffset(pageInput page.PageInput) int32 {
	return int32(pageInput.PageNumber-1) * int32(pageInput.PageSize)
}

// dbError creates a standardized database operation error
func (r *SQLCAppointmentRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableAppointments, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError creates a standardized entity not found error
func (r *SQLCAppointmentRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableAppointments, DriverSQL)
}

// wrapConversionError wraps domain conversion errors
func (r *SQLCAppointmentRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}
