package bus

import (
	c "clinic-vet-api/app/modules/employee/application/command"
	h "clinic-vet-api/app/modules/employee/application/handler"
	q "clinic-vet-api/app/modules/employee/application/query"
	"clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type EmployeeCqrsBus interface {
	CreateEmployee(ctx context.Context, cmd c.CreateEmployeeCommand) cqrs.CommandResult
	UpdateEmployee(ctx context.Context, cmd c.UpdateEmployeeCommand) cqrs.CommandResult
	DeleteEmployee(ctx context.Context, cmd c.DeleteEmployeeCommand) cqrs.CommandResult

	FindEmployeeByID(ctx context.Context, qry q.FindEmployeeByIDQuery) (h.EmployeeResult, error)
	FindActiveEmployees(ctx context.Context, qry q.FindActiveEmployeesQuery) (p.Page[h.EmployeeResult], error)
}

type employeeQueryBus struct {
	queryHandler   h.EmployeeQueryHandler
	commandHandler h.EmployeeCommandHandler
}

func NewEmployeeCqrsBus(
	queryHandler h.EmployeeQueryHandler,
	commandHandler h.EmployeeCommandHandler,
) EmployeeCqrsBus {
	return &employeeQueryBus{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}

func (b *employeeQueryBus) CreateEmployee(ctx context.Context, cmd c.CreateEmployeeCommand) cqrs.CommandResult {
	return b.commandHandler.HandleCreate(ctx, cmd)
}

func (b *employeeQueryBus) UpdateEmployee(ctx context.Context, cmd c.UpdateEmployeeCommand) cqrs.CommandResult {
	return b.commandHandler.HandleUpdate(ctx, cmd)
}

func (b *employeeQueryBus) DeleteEmployee(ctx context.Context, cmd c.DeleteEmployeeCommand) cqrs.CommandResult {
	return b.commandHandler.HandleDelete(ctx, cmd)
}

func (b *employeeQueryBus) FindEmployeeByID(ctx context.Context, qry q.FindEmployeeByIDQuery) (h.EmployeeResult, error) {
	return b.queryHandler.HandleFindByID(ctx, qry)
}

func (b *employeeQueryBus) FindActiveEmployees(ctx context.Context, qry q.FindActiveEmployeesQuery) (p.Page[h.EmployeeResult], error) {
	return b.queryHandler.HandleFindActives(ctx, qry)
}
