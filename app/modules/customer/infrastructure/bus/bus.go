// Package bus implements a simple command bus for handling commands in the customer module.
package bus

import (
	"clinic-vet-api/app/modules/customer/application/command"
	"clinic-vet-api/app/modules/customer/application/query"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/page"
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

func (bus *CustomerCommandBus) CreateCustomer(cmd command.CreateCustomerCommand) cqrs.CommandResult {
	return bus.CommandBus.Create(cmd)
}

func (bus *CustomerCommandBus) UpdateCustomer(cmd command.UpdateCustomerCommand) cqrs.CommandResult {
	return bus.CommandBus.Update(cmd)
}

func (bus *CustomerCommandBus) DeactivateCustomer(cmd command.DeactivateCustomerCommand) cqrs.CommandResult {
	return bus.CommandBus.Deactivate(cmd)
}

type CustomerQueryBus struct {
	QueryBus query.CustomerQueryHandler
}

func NewCustomerQueryBus(queryBus query.CustomerQueryHandler) *CustomerQueryBus {
	return &CustomerQueryBus{
		QueryBus: queryBus,
	}
}

func (bus *CustomerQueryBus) GetCustomerByID(cmd query.FindCustomerByIDQuery) (query.CustomerResult, error) {
	return bus.QueryBus.FindByID(cmd)
}

func (bus *CustomerQueryBus) FindCustomerByCriteria(cmd query.FindCustomerBySpecificationQuery) (page.Page[query.CustomerResult], error) {
	return bus.QueryBus.FindBySpecification(cmd)
}
