package handler

import (
	"clinic-vet-api/app/modules/core/repository"
	q "clinic-vet-api/app/modules/employee/application/query"
	"clinic-vet-api/app/shared/page"
	"context"
)

type EmployeeQueryHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeQueryHandler(employeeRepo repository.EmployeeRepository) *EmployeeQueryHandler {
	return &EmployeeQueryHandler{
		employeeRepo: employeeRepo,
	}
}

func (h *EmployeeQueryHandler) HandleFindActives(
	ctx context.Context, query q.FindActiveEmployeesQuery,
) (page.Page[EmployeeResult], error) {
	employeesPage, err := h.employeeRepo.FindActive(ctx, query.Pagination())
	if err != nil {
		return page.Page[EmployeeResult]{}, err
	}

	return page.MapItems(employeesPage, employeeToResult), nil
}

func (h *EmployeeQueryHandler) HandleFindByID(ctx context.Context, query q.FindEmployeeByIDQuery) (EmployeeResult, error) {
	employee, err := h.employeeRepo.FindByID(ctx, query.ID())
	if err != nil {
		return EmployeeResult{}, err
	}

	return employeeToResult(employee), nil
}
