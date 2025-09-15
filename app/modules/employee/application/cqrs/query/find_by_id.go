package query

import (
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"context"
)

type GetEmployeeByIDQuery struct {
	EmployeeID valueobject.EmployeeID
}

type GetEmployeeByIDHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewGetEmployeeByIDHandler(employeeRepo repository.EmployeeRepository) *GetEmployeeByIDHandler {
	return &GetEmployeeByIDHandler{
		employeeRepo: employeeRepo,
	}
}

func (h *GetEmployeeByIDHandler) Handle(ctx context.Context, query GetEmployeeByIDQuery) (EmployeeResult, error) {
	employee, err := h.employeeRepo.FindByID(ctx, query.EmployeeID)
	if err != nil {
		return EmployeeResult{}, err
	}

	return ToResult(&employee), nil
}
