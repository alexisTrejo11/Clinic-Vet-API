package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/owners/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type OwnerServiceFacade interface {
	GetOwnerByID(ctx context.Context, ownerID valueobject.OwnerID) (dto.OwnerDetail, error)
	ListOwners(ctx context.Context, params dto.GetOwnersRequest) (page.Page[[]dto.OwnerDetail], error)
	CreateOwner(ctx context.Context, data dto.OwnerCreate) (dto.OwnerDetail, error)
	UpdateOwner(ctx context.Context, ownerID valueobject.OwnerID, data dto.OwnerUpdate) (dto.OwnerDetail, error)
	SoftDeleteOwner(ctx context.Context, ownerID valueobject.OwnerID) error
}

type ownerUseCases struct {
	getOwnerByIDUseCase    *GetOwnerByIDUseCase
	listOwnersUseCase      *ListOwnersUseCase
	createOwnerUseCase     *CreateOwnerUseCase
	updateOwnerUseCase     *UpdateOwnerUseCase
	softDeleteOwnerUseCase *SoftDeleteOwnerUseCase
}

func NewOwnerUseCases(
	getOwnerByIDUC *GetOwnerByIDUseCase,
	listOwnersUC *ListOwnersUseCase,
	createOwnerUC *CreateOwnerUseCase,
	updateOwnerUC *UpdateOwnerUseCase,
	softDeleteOwnerUC *SoftDeleteOwnerUseCase,
) OwnerServiceFacade {
	return &ownerUseCases{
		getOwnerByIDUseCase:    getOwnerByIDUC,
		listOwnersUseCase:      listOwnersUC,
		createOwnerUseCase:     createOwnerUC,
		updateOwnerUseCase:     updateOwnerUC,
		softDeleteOwnerUseCase: softDeleteOwnerUC,
	}
}

func (uc *ownerUseCases) GetOwnerByID(ctx context.Context, ownerID valueobject.OwnerID) (dto.OwnerDetail, error) {
	return uc.getOwnerByIDUseCase.Execute(ctx, ownerID)
}

func (uc *ownerUseCases) ListOwners(ctx context.Context, params dto.GetOwnersRequest) (page.Page[[]dto.OwnerDetail], error) {
	return uc.listOwnersUseCase.Execute(ctx, params)
}

func (uc *ownerUseCases) CreateOwner(ctx context.Context, data dto.OwnerCreate) (dto.OwnerDetail, error) {
	return uc.createOwnerUseCase.Execute(ctx, data)
}

func (uc *ownerUseCases) UpdateOwner(ctx context.Context, ownerID valueobject.OwnerID, data dto.OwnerUpdate) (dto.OwnerDetail, error) {
	return uc.updateOwnerUseCase.Execute(ctx, ownerID, data)
}

func (uc *ownerUseCases) SoftDeleteOwner(ctx context.Context, ownerID valueobject.OwnerID) error {
	return uc.softDeleteOwnerUseCase.Execute(ctx, ownerID)
}
