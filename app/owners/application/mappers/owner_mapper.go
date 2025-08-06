package ownerMappers

import (
	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

func FromRequestCreate(ownerCreate ownerDTOs.OwnerCreate) ownerDomain.Owner {
	fullName, _ := userDomain.NewPersonName(ownerCreate.FirstName, ownerCreate.LastName)
	return ownerDomain.Owner{
		Photo:       ownerCreate.Photo,
		FullName:    fullName,
		PhoneNumber: ownerCreate.PhoneNumber,
		Address:     ownerCreate.Address,
		Gender:      ownerCreate.Gender,
		IsActive:    true,
	}
}

func ToResponse(owner ownerDomain.Owner) *ownerDTOs.OwnerResponse {
	return &ownerDTOs.OwnerResponse{
		Id:          owner.Id,
		Photo:       owner.Photo,
		Name:        owner.FullName.FullName(),
		PhoneNumber: owner.PhoneNumber,
		Address:     owner.Address,
		IsActive:    owner.IsActive,
	}

}
