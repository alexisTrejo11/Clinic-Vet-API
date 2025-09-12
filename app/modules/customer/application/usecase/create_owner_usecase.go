// Package usecase contains all the implementation to handle all the business logic related to customer entity
package usecase

import (
	"context"

	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/error"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/customer/application/dto"
)

type CreateCustomerUseCase struct {
	customerRepo repository.customerRepository
}

func NewCreatecustomerUseCase(customerRepo repository.CustomerRepository) *CreateCustomerUseCase {
	return &CreatecustomerUseCase{
		customerRepo: customerRepo,
	}
}

func (uc *CreatecustomerUseCase) Execute(ctx context.Context, createData dto.customerCreate) (dto.CustomerDetail, error) {
	if exists, err := uc.customerRepo.ExistsByPhone(ctx, createData.PhoneNumber); err != nil {
		return dto.CustomerDetail{}, err
	} else if exists {
		return dto.CustomerDetail{}, domainerr.HandlePhoneConflictError()
	}

	customer, err := mapper.FromRequestCreate(createData)
	if err != nil {
		return dto.CustomerDetail{}, err
	}

	if err := uc.customerRepo.Save(ctx, customer); err != nil {
		return dto.CustomerDetail{}, err
	}

	return mapper.ToResponse(customer), nil
}
