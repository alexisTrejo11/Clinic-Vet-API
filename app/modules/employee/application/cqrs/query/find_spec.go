package query

import (
	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"context"
)

type SearchEmployeesQuery struct {
	Specification specification.EmployeeSearchSpecification
}

type SearchEmployeesHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewSearchEmployeesHandler(employeeRepo repository.EmployeeRepository) *SearchEmployeesHandler {
	return &SearchEmployeesHandler{
		employeeRepo: employeeRepo,
	}
}

func (h *SearchEmployeesHandler) Handle(ctx context.Context, query SearchEmployeesQuery) (page.Page[EmployeeResult], error) {
	employeesPage, err := h.employeeRepo.FindBySpecification(ctx, query.Specification)
	if err != nil {
		return page.Page[EmployeeResult]{}, err
	}

	return page.MapItems(employeesPage, employeeToResult), nil
}
