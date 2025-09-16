package command

import (
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type CustomerCommandHandler interface {
	Create(command CreateCustomerCommand) cqrs.CommandResult
	Update(command UpdateCustomerCommand) cqrs.CommandResult
	Deactivate(command DeactivateCustomerCommand) cqrs.CommandResult
}

type customerCommandHandler struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerCommandHandler(customerRepository repository.CustomerRepository) CustomerCommandHandler {
	return &customerCommandHandler{
		customerRepository: customerRepository,
	}
}

func (h *customerCommandHandler) Create(command CreateCustomerCommand) cqrs.CommandResult {
	customer, err := command.ToEntity()
	if err != nil {
		return *cqrs.FailureResult("an error ocurred creating customer entity", err)
	}

	err = h.customerRepository.Save(command.CTX, customer)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred saving customer", err)
	}

	return *cqrs.SuccessResult(customer.ID().String(), "customer successfully created")
}

func (h *customerCommandHandler) Update(command UpdateCustomerCommand) cqrs.CommandResult {
	customer, err := h.customerRepository.FindByID(command.CTX, command.ID)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred finding customer", err)
	}

	err = command.UpdateEntity(&customer)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred updating customer entity", err)
	}

	err = h.customerRepository.Save(command.CTX, &customer)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred updating customer", err)
	}

	return *cqrs.SuccessResult(customer.ID().String(), "customer successfully updated")
}

func (h *customerCommandHandler) Deactivate(command DeactivateCustomerCommand) cqrs.CommandResult {
	customer, err := h.customerRepository.FindByID(command.CTX, command.ID)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred finding customer", err)
	}

	err = customer.Deactivate()
	if err != nil {
		return *cqrs.FailureResult("an error ocurred deactivating customer", err)
	}

	err = h.customerRepository.Save(command.CTX, &customer)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred saving customer", err)
	}

	return *cqrs.SuccessResult(customer.ID().String(), "customer successfully deactivated")
}
