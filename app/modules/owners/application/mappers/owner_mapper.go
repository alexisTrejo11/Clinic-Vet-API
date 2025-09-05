package mapper

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
)

func FromRequestCreate(ownerCreate dto.OwnerCreate) *entity.Owner {
	fullName, _ := valueobject.NewPersonName(ownerCreate.FirstName, ownerCreate.LastName)

	builder := entity.NewOwnerBuilder().
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

func ToResponse(owner *entity.Owner) dto.OwnerDetail {
	response := &dto.OwnerDetail{
		ID:          owner.ID().GetValue(),
		Photo:       owner.Photo(),
		Name:        owner.FullName().FullName(),
		PhoneNumber: owner.PhoneNumber(),
		Address:     owner.Address(),
		IsActive:    owner.IsActive(),
	}
	return *response
}

func ToResponseList(owners []entity.Owner) []dto.OwnerDetail {
	responses := make([]dto.OwnerDetail, len(owners))
	for i, owner := range owners {
		responses[i] = ToResponse(&owner)
	}
	return responses
}
