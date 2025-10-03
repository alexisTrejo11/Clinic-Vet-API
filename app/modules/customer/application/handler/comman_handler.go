package handler

import (
	repo "clinic-vet-api/app/modules/core/repository"
	c "clinic-vet-api/app/modules/customer/application/command"
	"clinic-vet-api/app/shared/cqrs"
	"context"
)

var (
	FailToCheckCustomerExistenceMsg = "an error ocurred checking customer existence"
	FailBuisnessValidationMsg       = "an error ocurred validating business rules"
	FailToSaveCustomerMsg           = "an error ocurred saving customer"
	FailSavingCustomerMsg           = "an error ocurred saving customer"
	FailToDeactivateCustomerMsg     = "an error ocurred deactivating customer"

	SuccessCustomerCreatedMsg     = "customer successfully created"
	SuccessCustomerUpdatedMsg     = "customer successfully updated"
	SuccessCustomerDeactivatedMsg = "customer successfully deactivated"
)

type CustomerCommandHandler struct {
	customerRepo repo.CustomerRepository
}

func NewCustomerCommandHandler(customerRepo repo.CustomerRepository) *CustomerCommandHandler {
	return &CustomerCommandHandler{customerRepo: customerRepo}
}

func (h *CustomerCommandHandler) HandleCreate(ctx context.Context, cmd c.CreateCustomerCommand) cqrs.CommandResult {
	customer := cmd.ToEntity()
	if err := customer.Validate(ctx); err != nil {
		return cqrs.FailureResult(FailBuisnessValidationMsg, err)
	}

	if err := h.customerRepo.Save(ctx, &customer); err != nil {
		return cqrs.FailureResult(FailToSaveCustomerMsg, err)
	}

	return cqrs.SuccessCreateResult(customer.ID().String(), SuccessCustomerCreatedMsg)
}

func (h *CustomerCommandHandler) HandleUpdate(ctx context.Context, cmd c.UpdateCustomerCommand) cqrs.CommandResult {
	customer, err := h.customerRepo.FindByID(ctx, cmd.ID())
	if err != nil {
		return cqrs.FailureResult(FailToCheckCustomerExistenceMsg, err)
	}

	customerUpdated := cmd.UpdateEntity(customer)
	err = h.customerRepo.Save(ctx, &customerUpdated)
	if err != nil {
		return cqrs.FailureResult(FailSavingCustomerMsg, err)
	}

	return cqrs.SuccessResult(SuccessCustomerUpdatedMsg)
}

func (h *CustomerCommandHandler) HandleDeactivate(ctx context.Context, cmd c.DeactivateCustomerCommand) cqrs.CommandResult {
	customer, err := h.customerRepo.FindByID(ctx, cmd.ID())
	if err != nil {
		return cqrs.FailureResult(FailToCheckCustomerExistenceMsg, err)
	}

	err = customer.Deactivate(ctx)
	if err != nil {
		return cqrs.FailureResult(FailToDeactivateCustomerMsg, err)
	}

	err = h.customerRepo.Save(ctx, &customer)
	if err != nil {
		return cqrs.FailureResult(FailSavingCustomerMsg, err)
	}

	return cqrs.SuccessResult(SuccessCustomerDeactivatedMsg)
}
