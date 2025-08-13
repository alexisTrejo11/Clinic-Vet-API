package ownerMappers

import (
	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

func FromRequestCreate(ownerCreate ownerDTOs.OwnerCreate) *ownerDomain.Owner {
	fullName, _ := user.NewPersonName(ownerCreate.FirstName, ownerCreate.LastName)

	builder := ownerDomain.NewOwnerBuilder(0, fullName, ownerCreate.PhoneNumber).
		WithGender(ownerCreate.Gender).
		WithPhoto(ownerCreate.Photo).
		WithDateOfBirth(ownerCreate.DateOfBirth)

	if ownerCreate.Address != nil {
		builder.WithAddress(*ownerCreate.Address)
	}

	newOwner := builder.Build()

	return newOwner
}

func ToResponse(owner *ownerDomain.Owner) *ownerDTOs.OwnerResponse {
	return &ownerDTOs.OwnerResponse{
		Id:          owner.ID(),
		Photo:       owner.Photo(),
		Name:        owner.FullName().FullName(),
		PhoneNumber: owner.PhoneNumber(),
		Address:     owner.Address(),
		IsActive:    owner.IsActive(),
	}
}
