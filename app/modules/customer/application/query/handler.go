// Package query contains all the application logic for handling customer queries
package query

import (
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
)

type CustomerQueryHandler interface {
	FindByID(qry FindCustomerByIDQuery) (CustomerResult, error)
	FindBySpecification(specialty FindCustomerBySpecificationQuery) (page.Page[CustomerResult], error)
}

type customerQueryHandler struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerQueryHandler(customerRepo repository.CustomerRepository) CustomerQueryHandler {
	return &customerQueryHandler{
		customerRepository: customerRepo,
	}
}

func (h *customerQueryHandler) FindByID(cmd FindCustomerByIDQuery) (CustomerResult, error) {
	customer, err := h.customerRepository.FindByID(cmd.CTX, cmd.ID)
	if err != nil {
		return CustomerResult{}, err
	}
	return customerToResult(customer), nil
}

func (h *customerQueryHandler) FindBySpecification(cmd FindCustomerBySpecificationQuery) (page.Page[CustomerResult], error) {
	customerPage, err := h.customerRepository.FindBySpecification(cmd.CTX, cmd.querySpect)
	if err != nil {
		return page.Page[CustomerResult]{}, err
	}

	return page.MapItems(customerPage, customerToResult), nil
}
