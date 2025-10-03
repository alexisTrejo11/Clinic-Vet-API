package handler

import (
	"clinic-vet-api/app/modules/core/repository"
	c "clinic-vet-api/app/modules/employee/application/command"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

var (
	FailFindEmployeeMsg   = "an error occurred finding employee"
	FailSaveEmployeeMsg   = "an error occurred saving employee"
	FailDeleteEmployeeMsg = "failing deleting employee"
	FailBuissnessLogicMsg = "employee business logic validation failed"
	EmployeeNotFoundMsg   = "employee not found"

	SuccessEmployeeCreatedMsg = "Employee created successfully"
	SuccessEmployeeUpdatedMsg = "Employee updated successfully"
	SuccessEmployeeDeletedMsg = "Employee deleted successfully"
)

type EmployeeCommandHandler struct {
	employeeRepo repository.EmployeeRepository
}

func NewEmployeeCommandHandler(employeeRepo repository.EmployeeRepository) *EmployeeCommandHandler {
	return &EmployeeCommandHandler{employeeRepo: employeeRepo}
}

func (h *EmployeeCommandHandler) HandleCreate(ctx context.Context, cmd c.CreateEmployeeCommand) cqrs.CommandResult {
	employee := cmd.ToEntity()

	if err := employee.Validate(ctx); err != nil {
		return cqrs.FailureResult(FailBuissnessLogicMsg, err)
	}

	if err := h.employeeRepo.Save(ctx, &employee); err != nil {
		return cqrs.FailureResult(FailSaveEmployeeMsg, err)
	}

	return cqrs.SuccessCreateResult(employee.ID().String(), SuccessEmployeeCreatedMsg)
}

func (h *EmployeeCommandHandler) HandleUpdate(ctx context.Context, cmd c.UpdateEmployeeCommand) cqrs.CommandResult {
	existingEmployee, err := h.employeeRepo.FindByID(ctx, cmd.EmployeeID())
	if err != nil {
		return cqrs.FailureResult(FailFindEmployeeMsg, err)
	}

	employeeUpdated := cmd.UpdateEmployee(existingEmployee)
	if err := h.employeeRepo.Save(ctx, &employeeUpdated); err != nil {
		return cqrs.FailureResult(FailSaveEmployeeMsg, err)
	}

	return cqrs.SuccessResult(SuccessEmployeeUpdatedMsg)
}

func (h *EmployeeCommandHandler) HandleDelete(ctx context.Context, cmd c.DeleteEmployeeCommand) cqrs.CommandResult {
	if _, err := h.employeeRepo.FindByID(ctx, cmd.ID()); err != nil {
		return cqrs.FailureResult(FailFindEmployeeMsg, err)
	}
	if err := h.employeeRepo.Delete(ctx, cmd.ID(), false); err != nil {
		return cqrs.FailureResult(FailDeleteEmployeeMsg, err)
	}

	return cqrs.SuccessResult(SuccessEmployeeDeletedMsg)
}
