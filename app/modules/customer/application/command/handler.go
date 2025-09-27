package command

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

type CustomerCommandHandler interface {
	Create(ctx context.Context, cmd CreateCustomerCommand) cqrs.CommandResult
	Update(ctx context.Context, cmd UpdateCustomerCommand) cqrs.CommandResult
	Deactivate(ctx context.Context, cmd DeactivateCustomerCommand) cqrs.CommandResult
}

type customerCommandHandler struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerCommandHandler(customerRepository repository.CustomerRepository) CustomerCommandHandler {
	return &customerCommandHandler{
		customerRepository: customerRepository,
	}
}

func (h *customerCommandHandler) Create(ctx context.Context, cmd CreateCustomerCommand) cqrs.CommandResult {
	customer := cmd.ToEntity()

	if err := customer.Validate(ctx); err != nil {
		return *cqrs.FailureResult("an error ocurred validating customer", err)
	}

	if err := h.customerRepository.Save(ctx, &customer); err != nil {
		return *cqrs.FailureResult("an error ocurred saving customer", err)
	}

	return *cqrs.SuccessResult(customer.ID().String(), "customer successfully created")
}

func (h *customerCommandHandler) Update(ctx context.Context, cmd UpdateCustomerCommand) cqrs.CommandResult {
	customer, err := h.customerRepository.FindByID(ctx, cmd.ID)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred finding customer", err)
	}

	err = cmd.UpdateEntity(&customer)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred updating customer entity", err)
	}

	err = h.customerRepository.Save(ctx, &customer)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred updating customer", err)
	}

	return *cqrs.SuccessResult(customer.ID().String(), "customer successfully updated")
}

func (h *customerCommandHandler) Deactivate(ctx context.Context, cmd DeactivateCustomerCommand) cqrs.CommandResult {
	customer, err := h.customerRepository.FindByID(ctx, cmd.ID)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred finding customer", err)
	}

	err = customer.Deactivate(ctx)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred deactivating customer", err)
	}

	err = h.customerRepository.Save(ctx, &customer)
	if err != nil {
		return *cqrs.FailureResult("an error ocurred saving customer", err)
	}

	return *cqrs.SuccessResult(customer.ID().String(), "customer successfully deactivated")
}
