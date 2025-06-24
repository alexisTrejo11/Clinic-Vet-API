package ownerUsecase

import (
	"context"

	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
)

type ListOwnersUseCase struct {
	ownerRepo ownerRepository.OwnerRepository
}

func NewListOwnersUseCase(ownerRepo ownerRepository.OwnerRepository) *ListOwnersUseCase {
	return &ListOwnersUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *ListOwnersUseCase) Execute(ctx context.Context, dto ownerDTOs.GetOwnersRequest) (ownerDTOs.OwnerListResponse, error) {
	// Set defaults
	if dto.Page.Limit <= 0 {
		dto.Page.Limit = 20
	}
	if dto.Page.Offset < 0 {
		dto.Page.Offset = 1
	}

	owners, err := uc.ownerRepo.List(ctx, "", dto.Page.Limit, dto.Page.Offset)
	if err != nil {
		return ownerDTOs.OwnerListResponse{}, err
	}

	ownerResponses := make([]ownerDTOs.OwnerResponse, len(owners))
	for i, owner := range owners {
		ownerResponses[i] = *ownerMappers.ToResponse(owner)

		if dto.WithPets {
		}
	}

	return ownerDTOs.OwnerListResponse{
		Owners: ownerResponses,
		//Total:   total,
		//Limit:   dto.Page.Limit,
		//Offset:  dto.Page.Offset,
		//HasMore: int64(dto.Page.Offset+dto.Page.Limit) < total,
	}, nil
}

type OwnerStatsResponse struct {
	TotalOwners    int64 `json:"total_owners"`
	ActiveOwners   int64 `json:"active_owners"`
	InactiveOwners int64 `json:"inactive_owners"`
	OwnersWithPets int64 `json:"owners_with_pets"`
}
