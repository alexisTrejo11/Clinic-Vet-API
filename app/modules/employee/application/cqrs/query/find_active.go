package query

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"context"
)

type GetActiveEmployeesQuery struct {
	PageInput page.PageInput
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
	employeesPage, err := h.employeeRepo.FindActive(ctx, query.PageInput)
	if err != nil {
		return page.Page[EmployeeResult]{}, err
	}

	responses := make([]EmployeeResult, len(employeesPage.Items))
	for i, emp := range employeesPage.Items {
		responses[i] = ToResult(&emp)
	}

	return page.NewPage(responses, employeesPage.Metadata), nil
}
