package ownerUsecase

import (
	"context"

	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
)

type GetOwnerByIdUseCase struct {
	ownerRepo ownerDomain.OwnerRepository
}

func NewGetOwnerByIdUseCase(ownerRepo ownerDomain.OwnerRepository) *GetOwnerByIdUseCase {
	return &GetOwnerByIdUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *GetOwnerByIdUseCase) Execute(ctx context.Context, id int) (ownerDTOs.OwnerResponse, error) {
	owner, err := uc.ownerRepo.GetById(ctx, id)
	if err != nil {
		return ownerDTOs.OwnerResponse{}, ownerDomain.HandleGetByIdError(err, id)
	}

	ownerResponse := ownerMappers.ToResponse(&owner)
	ownerResponse.Pets = petMapper.ToResponseList(owner.Pets())

	return ownerResponse, nil
}
