package ownerUsecase

import (
	"context"

	DTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	ownerMappers "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/mappers"
	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListOwnersUseCase struct {
	ownerRepo ownerDomain.OwnerRepository
}

func NewListOwnersUseCase(ownerRepo ownerDomain.OwnerRepository) *ListOwnersUseCase {
	return &ListOwnersUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *ListOwnersUseCase) Execute(ctx context.Context, dto DTOs.GetOwnersRequest) (page.Page[[]DTOs.OwnerResponse], error) {
	ownersPage, err := uc.ownerRepo.List(ctx, dto.Page)
	if err != nil {
		return page.Page[[]DTOs.OwnerResponse]{}, err
	}

	ownerResponses := ownerMappers.ToResponseList(ownersPage.Data)

	pageResponse := page.NewPage(ownerResponses, ownersPage.Metadata)
	return pageResponse, nil
}

type OwnerStatsResponse struct {
	TotalOwners    int64 `json:"total_owners"`
	ActiveOwners   int64 `json:"active_owners"`
	InactiveOwners int64 `json:"inactive_owners"`
	OwnersWithPets int64 `json:"owners_with_pets"`
}
