package ownerMappers

import (
	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
)

func FromRequestCreate(ownerCreate ownerDTOs.OwnerCreate) *ownerDomain.Owner {
	fullName, _ := valueObjects.NewPersonName(ownerCreate.FirstName, ownerCreate.LastName)

	builder := ownerDomain.NewOwnerBuilder().
		WithFullName(fullName).
		WithIsActive(true).
		WithPhoneNumber(ownerCreate.PhoneNumber).
		WithGender(ownerCreate.Gender).
		WithPhoto(ownerCreate.Photo).
		WithDateOfBirth(ownerCreate.DateOfBirth)

	if ownerCreate.Address != nil {
		builder.WithAddress(*ownerCreate.Address)
	}

	newOwner := builder.Build()

	return newOwner
}

func ToResponse(owner *ownerDomain.Owner) ownerDTOs.OwnerResponse {
	response := &ownerDTOs.OwnerResponse{
		Id:          owner.Id(),
		Photo:       owner.Photo(),
		Name:        owner.FullName().FullName(),
		PhoneNumber: owner.PhoneNumber(),
		Address:     owner.Address(),
		IsActive:    owner.IsActive(),
	}
	return *response
}

func ToResponseList(owners []ownerDomain.Owner) []ownerDTOs.OwnerResponse {
	responses := make([]ownerDTOs.OwnerResponse, len(owners))
	for i, owner := range owners {
		responses[i] = ToResponse(&owner)
	}
	return responses
}
