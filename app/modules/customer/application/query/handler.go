// Package query contains all the application logic for handling customer queries
package query

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"context"
)

type CustomerQueryHandler interface {
	FindByID(ctx context.Context, qry FindCustomerByIDQuery) (CustomerResult, error)
	FindBySpecification(ctx context.Context, qry FindCustomerBySpecificationQuery) (page.Page[CustomerResult], error)
}

type customerQueryHandler struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerQueryHandler(customerRepo repository.CustomerRepository) CustomerQueryHandler {
	return &customerQueryHandler{
		customerRepository: customerRepo,
	}
}

func (h *customerQueryHandler) FindByID(ctx context.Context, cmd FindCustomerByIDQuery) (CustomerResult, error) {
	customer, err := h.customerRepository.FindByID(ctx, cmd.ID)
	if err != nil {
		return CustomerResult{}, err
	}
	return customerToResult(customer), nil
}

func (h *customerQueryHandler) FindBySpecification(ctx context.Context, cmd FindCustomerBySpecificationQuery) (page.Page[CustomerResult], error) {
	customerPage, err := h.customerRepository.FindBySpecification(ctx, cmd.querySpect)
	if err != nil {
		return page.Page[CustomerResult]{}, err
	}

	return page.MapItems(customerPage, customerToResult), nil
}
