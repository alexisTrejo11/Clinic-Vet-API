package usecase

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	mapper "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/mappers"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListOwnersUseCase struct {
	ownerRepo repository.OwnerRepository
}

func NewListOwnersUseCase(ownerRepo repository.OwnerRepository) *ListOwnersUseCase {
	return &ListOwnersUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *ListOwnersUseCase) Execute(ctx context.Context, queryData dto.GetOwnersRequest) (page.Page[[]dto.OwnerDetail], error) {
	ownersPage, err := uc.ownerRepo.List(ctx, queryData.Page)
	if err != nil {
		return page.EmptyPage[[]dto.OwnerDetail](), err
	}

	ownerResponses := mapper.ToResponseList(ownersPage.Data)
	return page.NewPage(ownerResponses, ownersPage.Metadata), nil
}

type OwnerStatsResponse struct {
	TotalOwners    int64 `json:"total_owners"`
	ActiveOwners   int64 `json:"active_owners"`
	InactiveOwners int64 `json:"inactive_owners"`
	OwnersWithPets int64 `json:"owners_with_pets"`
}
