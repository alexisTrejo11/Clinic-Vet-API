package query

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"context"
)

type GetActiveEmployeesQuery struct {
	PaginationRequest page.PaginationRequest
}

type GetActiveEmployeesHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewGetActiveEmployeesHandler(employeeRepo repository.EmployeeRepository) *GetActiveEmployeesHandler {
	return &GetActiveEmployeesHandler{
		employeeRepo: employeeRepo,
	}
}

func (h *GetActiveEmployeesHandler) Handle(ctx context.Context, query GetActiveEmployeesQuery) (page.Page[EmployeeResult], error) {
	employeesPage, err := h.employeeRepo.FindActive(ctx, query.PaginationRequest)
	if err != nil {
		return page.Page[EmployeeResult]{}, err
	}

	return page.MapItems(employeesPage, employeeToResult), nil
}
