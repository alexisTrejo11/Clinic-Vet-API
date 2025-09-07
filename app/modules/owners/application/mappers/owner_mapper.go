package mapper

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/owner"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
)

func FromRequestCreate(ownerCreate dto.OwnerCreate) (*owner.Owner, error) {
	fullName, _ := valueobject.NewPersonName(ownerCreate.FirstName, ownerCreate.LastName)

	ownerEntity, err := owner.NewOwner(
		valueobject.OwnerID{},
		owner.WithFullName(fullName),
		owner.WithIsActive(true),
		owner.WithPhoneNumber(ownerCreate.PhoneNumber),
		owner.WithGender(ownerCreate.Gender),
		owner.WithPhoto(ownerCreate.Photo),
		owner.WithDateOfBirth(ownerCreate.DateOfBirth),
	)
	if err != nil {
		return nil, err
	}

	return ownerEntity, nil
}

func ToResponse(owner *owner.Owner) dto.OwnerDetail {
	response := &dto.OwnerDetail{
		ID:          owner.ID().Value(),
		Photo:       owner.Photo(),
		Name:        owner.FullName().FullName(),
		PhoneNumber: owner.PhoneNumber(),
		IsActive:    owner.IsActive(),
	}
	return *response
}

func ToResponseList(owners []owner.Owner) []dto.OwnerDetail {
	responses := make([]dto.OwnerDetail, len(owners))
	for i, owner := range owners {
		responses[i] = ToResponse(&owner)
	}
	return responses
}
