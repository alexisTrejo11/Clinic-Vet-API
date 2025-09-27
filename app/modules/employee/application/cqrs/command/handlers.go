package command

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	apperror "clinic-vet-api/app/shared/error/application"
	"context"
)

type CreateEmployeeHandler struct {
	employeeRepo repository.EmployeeRepository
	eventBus     EventBus
}

// TODO: EventBus implementation
func NewCreateEmployeeHandler(employeeRepo repository.EmployeeRepository) *CreateEmployeeHandler {
	return &CreateEmployeeHandler{
		employeeRepo: employeeRepo,
		eventBus:     nil,
	}
}

type UpdateEmployeeHandler struct {
	employeeRepo repository.EmployeeRepository
	eventBus     EventBus
}

func NewUpdateEmployeeHandler(employeeRepo repository.EmployeeRepository, eventBus EventBus) *UpdateEmployeeHandler {
	return &UpdateEmployeeHandler{
		employeeRepo: employeeRepo,
		eventBus:     eventBus,
	}
}

type EventBus interface{}

type DeleteEmployeeHandler struct {
	employeeRepo repository.EmployeeRepository
	eventBus     EventBus
}

func NewDeleteEmployeeHandler(employeeRepo repository.EmployeeRepository, eventBus EventBus) *DeleteEmployeeHandler {
	return &DeleteEmployeeHandler{
		employeeRepo: employeeRepo,
		eventBus:     eventBus,
	}
}

func (h *CreateEmployeeHandler) Handle(ctx context.Context, cmd CreateEmployeeCommand) cqrs.CommandResult {
	employee := cmd.ToEntity()
	if err := h.employeeRepo.Save(ctx, &employee); err != nil {
		return *cqrs.FailureResult("an error occurred saving employee", err)
	}

	if err := employee.Validate(ctx); err != nil {
		return *cqrs.FailureResult("validation error", err)
	}

	return *cqrs.SuccessResult(employee.ID().String(), "Employee created successfully")
}

func (h *UpdateEmployeeHandler) Handle(ctx context.Context, cmd UpdateEmployeeCommand) cqrs.CommandResult {
	employee, err := h.employeeRepo.FindByID(ctx, cmd.EmployeeID)
	if err != nil {
		return *cqrs.FailureResult("an error occurred finding employee", err)
	}

	if err := cmd.updateEmployee(ctx, &employee); err != nil {
		return *cqrs.FailureResult("an error occurred updating employee", err)
	}

	if err := h.employeeRepo.Save(ctx, &employee); err != nil {
		return *cqrs.FailureResult("an error occurred saving employee", err)
	}

	return *cqrs.SuccessResult(employee.ID().String(), "Employee updated successfully")
}

func (h *DeleteEmployeeHandler) Handle(ctx context.Context, cmd DeleteEmployeeCommand) cqrs.CommandResult {
	exists, err := h.employeeRepo.ExistsByID(ctx, cmd.EmployeeID)
	if err != nil {
		return *cqrs.FailureResult("failing finding employee", err)
	}
	if !exists {
		return *cqrs.FailureResult("employee not found", apperror.EntityNotFoundValidationError("Employee", "id", cmd.EmployeeID.String()))
	}

	if err := h.employeeRepo.SoftDelete(ctx, cmd.EmployeeID); err != nil {
		return *cqrs.FailureResult("failing deleting employee", err)
	}

	return *cqrs.SuccessResult("", "Employee deleted successfully")
}
