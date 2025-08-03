package ownerUsecase

import (
	"context"

	ownerAppErr "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application"
	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	petMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/mapper"
)

type GetOwnerByIdUseCase struct {
	ownerRepo ownerRepository.OwnerRepository
}

func NewGetOwnerByIdUseCase(ownerRepo ownerRepository.OwnerRepository) *GetOwnerByIdUseCase {
	return &GetOwnerByIdUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *GetOwnerByIdUseCase) Execute(ctx context.Context, id int, includePets bool) (*ownerDTOs.OwnerResponse, error) {
	owner, err := uc.ownerRepo.GetByID(ctx, id, includePets)
	if err != nil {
		return nil, ownerAppErr.HandleGetByIdError(err, id)
	}

	ownerResponse := ownerMappers.ToResponse(owner)
	if includePets {
		ownerResponse.Pets = petMapper.ToResponseList(owner.Pets)
	}

	return ownerResponse, nil
}
