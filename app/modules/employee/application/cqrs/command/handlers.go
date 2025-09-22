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
	employee, err := cmd.createEmployee(ctx)
	if err != nil {
		return *cqrs.FailureResult("an error occurred creating employee", err)
	}

	if err := h.employeeRepo.Save(ctx, employee); err != nil {
		return *cqrs.FailureResult("an error occurred saving employee", err)
	}

	/*
		if err := h.eventBus.PublishEvents(ctx, employee.DomainEvents()); err != nil {
			// Log error but don't fail the operation
			// logger.Error("Failed to publish events", err)
		}
	*/

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

	// Save updated employee
	if err := h.employeeRepo.Update(ctx, &employee); err != nil {
		return *cqrs.FailureResult("an error occurred saving employee", err)
	}

	/*
		if err := h.eventBus.PublishEvents(ctx, employee.DomainEvents()); err != nil {
			// Log error but don't fail the operation
		}
	*/

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

	// Publish domain events (employee deleted event)
	// events := []DomainEvent{NewEmployeeDeletedEvent(cmd.EmployeeID)}
	// h.eventBus.PublishEvents(ctx, events)

	return *cqrs.SuccessResult("", "Employee deleted successfully")
}
