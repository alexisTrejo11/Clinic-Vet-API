// Package repository contains the implementation of the customer repository using SQLC.
package repository

import (
	"fmt"

	dberr "clinic-vet-api/app/shared/error/infrastructure/database"
)

const (
	ErrMsgGetCustomer         = "failed to get customer"
	ErrMsgGetCustomerByPhone  = "failed to get customer by phone"
	ErrMsgListCustomers       = "failed to list customers"
	ErrMsgCreateCustomer      = "failed to create customer"
	ErrMsgUpdateCustomer      = "failed to update customer"
	ErrMsgDeleteCustomer      = "failed to delete customer"
	ErrMsgActivateCustomer    = "failed to activate customer"
	ErrMsgDeactivateCustomer  = "failed to deactivate customer"
	ErrMsgCheckExistsByPhone  = "failed to check if customer exists by phone"
	ErrMsgCheckExistsByID     = "failed to check if customer exists by ID"
	ErrMsgListPetsForCustomer = "failed to list pets for customer"
	ErrMsgConvertToDomain     = "failed to convert to domain entity"
	ErrMsgCountCustomers      = "failed to count customers"
	ErrMsgCreateCustomerID    = "failed to create customer ID from result"
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
	return dberr.DatabaseOperationError(operation, TableCustomers, DriverSQL, fmt.Errorf("%s: %v", message, err))
}

// notFoundError creates a standardized entity not found error
func (r *SqlcCustomerRepository) notFoundError(parameterName, parameterValue string) error {
	return dberr.EntityNotFoundError(parameterName, parameterValue, OpSelect, TableCustomers, DriverSQL)
}

// wrapConversionError wraps domain conversion errors
func (r *SqlcCustomerRepository) wrapConversionError(err error) error {
	return fmt.Errorf("%s: %w", ErrMsgConvertToDomain, err)
}
