// Package query contains query bus interfaces and implementations for handling read operations.
package query

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"context"
)

type EmployeeQueryBus interface {
	GetEmployeeByID(ctx context.Context, query GetEmployeeByIDQuery) (EmployeeResult, error)
	GetActiveEmployees(ctx context.Context, query GetActiveEmployeesQuery) (page.Page[EmployeeResult], error)
}

type employeeQueryBus struct {
	getByIDHandler *GetEmployeeByIDHandler

	getActiveHandler *GetActiveEmployeesHandler
}

func NewEmployeeQueryBus(employeeRepo repository.EmployeeRepository) EmployeeQueryBus {
	getByIDHandler := NewGetEmployeeByIDHandler(employeeRepo)

	getActiveHandler := NewGetActiveEmployeesHandler(employeeRepo)
	return &employeeQueryBus{
		getByIDHandler: getByIDHandler,

		getActiveHandler: getActiveHandler,
	}
}

func (b *employeeQueryBus) GetEmployeeByID(ctx context.Context, query GetEmployeeByIDQuery) (EmployeeResult, error) {
	return b.getByIDHandler.Handle(ctx, query)
}

func (b *employeeQueryBus) GetActiveEmployees(ctx context.Context, query GetActiveEmployeesQuery) (page.Page[EmployeeResult], error) {
	return b.getActiveHandler.Handle(ctx, query)
}
