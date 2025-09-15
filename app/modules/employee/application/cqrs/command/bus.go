package command

import (
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type EmployeeCommandBus struct {
	createHandler *CreateEmployeeHandler
	updateHandler *UpdateEmployeeHandler
	deleteHandler *DeleteEmployeeHandler
}

func NewEmployeeCommandBus(
	createHandler *CreateEmployeeHandler,
	updateHandler *UpdateEmployeeHandler,
	deleteHandler *DeleteEmployeeHandler,
) *EmployeeCommandBus {
	return &EmployeeCommandBus{
		createHandler: createHandler,
		updateHandler: updateHandler,
		deleteHandler: deleteHandler,
	}
}

func (b *EmployeeCommandBus) CreateEmployee(ctx context.Context, cmd CreateEmployeeCommand) cqrs.CommandResult {
	return b.createHandler.Handle(ctx, cmd)
}

func (b *EmployeeCommandBus) UpdateEmployee(ctx context.Context, cmd UpdateEmployeeCommand) cqrs.CommandResult {
	return b.updateHandler.Handle(ctx, cmd)
}

func (b *EmployeeCommandBus) DeleteEmployee(ctx context.Context, cmd DeleteEmployeeCommand) cqrs.CommandResult {
	return b.deleteHandler.Handle(ctx, cmd)
}
