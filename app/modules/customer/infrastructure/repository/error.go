// Package repository contains the implementation of the owner repository using SQLC.
package repository

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

const (
	ErrMsgGetCustomer         = "failed to get owner"
	ErrMsgGetCustomerByPhone  = "failed to get owner by phone"
	ErrMsgListCustomers       = "failed to list owners"
	ErrMsgCreateCustomer      = "failed to create owner"
	ErrMsgUpdateCustomer      = "failed to update owner"
	ErrMsgDeleteCustomer      = "failed to delete owner"
	ErrMsgActivateCustomer    = "failed to activate owner"
	ErrMsgDeactivateCustomer  = "failed to deactivate owner"
	ErrMsgCheckExistsByPhone  = "failed to check if owner exists by phone"
	ErrMsgCheckExistsByID     = "failed to check if owner exists by ID"
	ErrMsgListPetsForCustomer = "failed to list pets for owner"
	ErrMsgConvertToDomain     = "failed to convert to domain entity"
	ErrMsgCountCustomers      = "failed to count owners"
	ErrMsgCreateCustomerID    = "failed to create owner ID from result"
)

const (
	OpSelect = "select"
	OpInsert = "insert"
	OpUpdate = "update"
	OpDelete = "delete"
	OpCount  = "count"

	TableCustomers = "customers"
	DriverSQL      = "sql"

	// Goroutine timeout for concurrent operations
	ConcurrentOpTimeout = 1
)

// dbError creates a standardized database operation error
func (r *SqlcCustomerRepository) dbError(operation, message string, err error) error {
	return dberr.DatabaseOperationError(operation, TableCustomers, DriverSQL, fmt.Sprintf("%s: %v", message, err))
}

// notFoundError creates a standardized entity not found error
func (r *SqlcCustomerRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableCustomers, DriverSQL)
}

// wrapConversionError wraps domain conversion errors
func (r *SqlcCustomerRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}
