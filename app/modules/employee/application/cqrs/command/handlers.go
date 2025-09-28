package command

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
	"context"
)

type CreateEmployeeHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewCreateEmployeeHandler(employeeRepo repository.EmployeeRepository) *CreateEmployeeHandler {
	return &CreateEmployeeHandler{employeeRepo: employeeRepo}
}

type UpdateEmployeeHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewUpdateEmployeeHandler(employeeRepo repository.EmployeeRepository) *UpdateEmployeeHandler {
	return &UpdateEmployeeHandler{employeeRepo: employeeRepo}
}

type DeleteEmployeeHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewDeleteEmployeeHandler(employeeRepo repository.EmployeeRepository) *DeleteEmployeeHandler {
	return &DeleteEmployeeHandler{employeeRepo: employeeRepo}
}

func (h *CreateEmployeeHandler) Handle(ctx context.Context, cmd CreateEmployeeCommand) cqrs.CommandResult {
	employee := cmd.ToEntity()
	if err := h.employeeRepo.Save(ctx, &employee); err != nil {
		return *cqrs.FailureResult("an error occurred saving employee", err)
	}

	if err := employee.Validate(ctx); err != nil {
		return *cqrs.FailureResult("validation error", err)
	}

	return *cqrs.SuccessCreateResult(employee.ID().String(), "Employee created successfully")
}

func (h *UpdateEmployeeHandler) Handle(ctx context.Context, cmd UpdateEmployeeCommand) cqrs.CommandResult {
	existingEmployee, err := h.employeeRepo.FindByID(ctx, cmd.EmployeeID)
	if err != nil {
		return *cqrs.FailureResult("an error occurred finding employee", err)
	}

	employeeUpdated := cmd.updateEmployee(existingEmployee)
	if err := h.employeeRepo.Save(ctx, &employeeUpdated); err != nil {
		return *cqrs.FailureResult("an error occurred saving employee", err)
	}

	return *cqrs.SuccessResult("Employee updated successfully")
}

func (h *DeleteEmployeeHandler) Handle(ctx context.Context, cmd DeleteEmployeeCommand) cqrs.CommandResult {
	if exists, err := h.employeeRepo.ExistsByID(ctx, cmd.EmployeeID); err != nil {
		return *cqrs.FailureResult("failing finding employee", err)
	} else if !exists {
		return *cqrs.FailureResult("employee not found", apperror.EntityNotFoundValidationError("Employee", "id", cmd.EmployeeID.String()))
	}

	if err := h.employeeRepo.Delete(ctx, cmd.EmployeeID, false); err != nil {
		return *cqrs.FailureResult("failing deleting employee", err)
	}

	return *cqrs.SuccessResult("Employee deleted successfully")
}
