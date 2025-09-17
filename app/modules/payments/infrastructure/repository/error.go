package repository

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

const (
	OpSelect = "select"
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
	OpCount  = "count"

	TablePayments = "payments"
	DriverSQL     = "sql"
)

// Error messages
const (
	ErrMsgGetPayment              = "failed to get payment"
	ErrMsgGetPaymentByTransaction = "failed to get payment by transaction ID"
	ErrMsgListPayments            = "failed to list payments"
	ErrMsgListPaymentsByUser      = "failed to list payments by user ID"
	ErrMsgListPaymentsByStatus    = "failed to list payments by status"
	ErrMsgListOverduePayments     = "failed to list overdue payments"
	ErrMsgListPaymentsByDateRange = "failed to list payments by date range"
	ErrMsgCreatePayment           = "failed to create payment"
	ErrMsgUpdatePayment           = "failed to update payment"
	ErrMsgSoftDeletePayment       = "failed to soft delete payment"
	ErrMsgCountPayments           = "failed to count payments"
	ErrMsgConvertToDomain         = "failed to convert to domain entity"
	ErrMsgCreatePaymentID         = "failed to create payment ID from result"
	ErrMsgSearchPayments          = "failed to search payments"
)

// dbError creates a standardized database operation error
func (r *SQLCPaymentRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TablePayments, DriverSQL, fmt.Errorf("%s: %v", message, err))
}

// notFoundError creates a standardized entity not found error
func (r *SQLCPaymentRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TablePayments, DriverSQL)
}

// wrapConversionError wraps domain conversion errors
func (r *SQLCPaymentRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}
