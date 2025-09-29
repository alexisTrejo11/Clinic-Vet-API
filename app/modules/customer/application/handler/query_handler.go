// Package query contains all the application logic for handling customer queries
package handler

import (
	"clinic-vet-api/app/modules/core/repository"
	q "clinic-vet-api/app/modules/customer/application/query"
	"clinic-vet-api/app/shared/page"
	"context"
)

type CustomerQueryHandler struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerQueryHandler(customerRepo repository.CustomerRepository) *CustomerQueryHandler {
	return &CustomerQueryHandler{
		customerRepository: customerRepo,
	}
}

func (h *CustomerQueryHandler) HandleFindByID(ctx context.Context, query q.FindCustomerByIDQuery) (CustomerResult, error) {
	customer, err := h.customerRepository.FindByID(ctx, query.ID())
	if err != nil {
		return CustomerResult{}, err
	}
	return customerToResult(customer), nil
}

func (h *CustomerQueryHandler) HandleFindBySpecification(ctx context.Context, query q.FindCustomerBySpecificationQuery) (page.Page[CustomerResult], error) {
	customerPage, err := h.customerRepository.FindBySpecification(ctx, query.QuerySpecification())
	if err != nil {
		return page.Page[CustomerResult]{}, err
	}

	return page.MapItems(customerPage, customerToResult), nil
}
