package mapper

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/customer"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/customer/application/dto"
)

func FromRequestCreate(ownerCreate dto.CustomerCreate) (*customer.Customer, error) {
	fullName, _ := valueobject.NewPersonName(ownerCreate.FirstName, ownerCreate.LastName)

	ownerEntity, err := customer.NewCustomer(
		valueobject.CustomerID{},
		customer.WithFullName(fullName),
		customer.WithIsActive(true),
		customer.WithPhoneNumber(ownerCreate.PhoneNumber),
		customer.WithGender(ownerCreate.Gender),
		customer.WithPhoto(ownerCreate.Photo),
		customer.WithDateOfBirth(ownerCreate.DateOfBirth),
	)
	if err != nil {
		return nil, err
	}

	return ownerEntity, nil
}

func ToResponse(customer *customer.Customer) dto.CustomerDetail {
	response := &dto.CustomerDetail{
		ID:          customer.ID().Value(),
		Photo:       customer.Photo(),
		Name:        customer.FullName().FullName(),
		PhoneNumber: customer.PhoneNumber(),
		IsActive:    customer.IsActive(),
	}
	return *response
}

func ToResponseList(owners []customer.Customer) []dto.CustomerDetail {
	responses := make([]dto.CustomerDetail, len(owners))
	for i, customer := range owners {
		responses[i] = ToResponse(&customer)
	}
	return responses
}
