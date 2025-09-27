// Package bus implements a simple command bus for handling commands in the customer module.
package bus

import (
	"clinic-vet-api/app/modules/customer/application/command"
	"clinic-vet-api/app/modules/customer/application/query"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/page"
	"context"
)

type CustomerBus struct {
	CommandBus *CustomerCommandBus
	QueryBus   *CustomerQueryBus
}

func NewCustomerModuleBus(commandBus *CustomerCommandBus, queryBus *CustomerQueryBus) *CustomerBus {
	return &CustomerBus{
		CommandBus: commandBus,
		QueryBus:   queryBus,
	}
}

type CustomerCommandBus struct {
	CommandBus command.CustomerCommandHandler
}

func NewCustomerCommandBus(commandBus command.CustomerCommandHandler) *CustomerCommandBus {
	return &CustomerCommandBus{
		CommandBus: commandBus,
	}
}

func (bus *CustomerCommandBus) CreateCustomer(ctx context.Context, cmd command.CreateCustomerCommand) cqrs.CommandResult {
	return bus.CommandBus.Create(ctx, cmd)
}

func (bus *CustomerCommandBus) UpdateCustomer(ctx context.Context, cmd command.UpdateCustomerCommand) cqrs.CommandResult {
	return bus.CommandBus.Update(ctx, cmd)
}

func (bus *CustomerCommandBus) DeactivateCustomer(ctx context.Context, cmd command.DeactivateCustomerCommand) cqrs.CommandResult {
	return bus.CommandBus.Deactivate(ctx, cmd)
}

type CustomerQueryBus struct {
	QueryBus query.CustomerQueryHandler
}

func NewCustomerQueryBus(queryBus query.CustomerQueryHandler) *CustomerQueryBus {
	return &CustomerQueryBus{
		QueryBus: queryBus,
	}
}

func (bus *CustomerQueryBus) GetCustomerByID(ctx context.Context, cmd query.FindCustomerByIDQuery) (query.CustomerResult, error) {
	return bus.QueryBus.FindByID(ctx, cmd)
}

func (bus *CustomerQueryBus) FindCustomerByCriteria(ctx context.Context, cmd query.FindCustomerBySpecificationQuery) (page.Page[query.CustomerResult], error) {
	return bus.QueryBus.FindBySpecification(ctx, cmd)
}
