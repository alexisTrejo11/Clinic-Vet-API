package query

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
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

	return *FromEntityToResult(customer), nil
}

func (h *customerQueryHandler) FindBySpecification(cmd FindCustomerBySpecificationQuery) (page.Page[CustomerResult], error) {
	customerPage, err := h.customerRepository.FindBySpecification(cmd.CTX, cmd.querySpect)
	if err != nil {
		return page.EmptyPage[CustomerResult](), err
	}
	customerResults := FromEntityListToResultList(customerPage.Items)
	return page.NewPage(customerResults, customerPage.Metadata), nil
}
