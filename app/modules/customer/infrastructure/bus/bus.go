// Package bus implements a simple command bus for handling commands in the customer module.
package bus

import (
	c "clinic-vet-api/app/modules/customer/application/command"
	h "clinic-vet-api/app/modules/customer/application/handler"
	q "clinic-vet-api/app/modules/customer/application/query"
	"clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type CustomerBus interface {
	// Command
	CreateCustomer(ctx context.Context, cmd c.CreateCustomerCommand) cqrs.CommandResult
	UpdateCustomer(ctx context.Context, cmd c.UpdateCustomerCommand) cqrs.CommandResult
	DeactivateCustomer(ctx context.Context, cmd c.DeactivateCustomerCommand) cqrs.CommandResult

	// Query
	FindCustomerByID(ctx context.Context, query q.FindCustomerByIDQuery) (h.CustomerResult, error)
	FindCustomerByCriteria(ctx context.Context, query q.FindCustomerBySpecificationQuery) (p.Page[h.CustomerResult], error)
}

type customerBus struct {
	CommandHandler h.CustomerCommandHandler
	QueryHandler   h.CustomerQueryHandler
}

func NewCustomerBus(commandHandler h.CustomerCommandHandler, queryHandler h.CustomerQueryHandler) CustomerBus {
	return &customerBus{
		CommandHandler: commandHandler,
		QueryHandler:   queryHandler,
	}
}

func (bus *customerBus) CreateCustomer(ctx context.Context, cmd c.CreateCustomerCommand) cqrs.CommandResult {
	return bus.CommandHandler.HandleCreate(ctx, cmd)
}

func (bus *customerBus) UpdateCustomer(ctx context.Context, cmd c.UpdateCustomerCommand) cqrs.CommandResult {
	return bus.CommandHandler.HandleUpdate(ctx, cmd)
}

func (bus *customerBus) DeactivateCustomer(ctx context.Context, cmd c.DeactivateCustomerCommand) cqrs.CommandResult {
	return bus.CommandHandler.HandleDeactivate(ctx, cmd)
}

func (bus *customerBus) FindCustomerByID(ctx context.Context, query q.FindCustomerByIDQuery) (h.CustomerResult, error) {
	return bus.QueryHandler.HandleFindByID(ctx, query)
}

func (bus *customerBus) FindCustomerByCriteria(ctx context.Context, query q.FindCustomerBySpecificationQuery) (p.Page[h.CustomerResult], error) {
	return bus.QueryHandler.HandleFindBySpecification(ctx, query)
}
