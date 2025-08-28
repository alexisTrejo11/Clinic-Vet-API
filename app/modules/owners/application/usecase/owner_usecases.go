package ownerUsecase

import (
	"context"

	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type OwnerServiceFacade interface {
	GetOwnerById(ctx context.Context, ownerId int) (ownerDTOs.OwnerResponse, error)
	ListOwners(ctx context.Context, params ownerDTOs.GetOwnersRequest) (page.Page[[]ownerDTOs.OwnerResponse], error)
	CreateOwner(ctx context.Context, data ownerDTOs.OwnerCreate) (ownerDTOs.OwnerResponse, error)
	UpdateOwner(ctx context.Context, ownerId int, data ownerDTOs.OwnerUpdate) (ownerDTOs.OwnerResponse, error)
	SoftDeleteOwner(ctx context.Context, ownerId int) error
}

type ownerUseCases struct {
	getOwnerByIdUseCase    *GetOwnerByIdUseCase
	listOwnersUseCase      *ListOwnersUseCase
	createOwnerUseCase     *CreateOwnerUseCase
	updateOwnerUseCase     *UpdateOwnerUseCase
	softDeleteOwnerUseCase *SoftDeleteOwnerUseCase
}

func NewOwnerUseCases(
	getOwnerByIdUC *GetOwnerByIdUseCase,
	listOwnersUC *ListOwnersUseCase,
	createOwnerUC *CreateOwnerUseCase,
	updateOwnerUC *UpdateOwnerUseCase,
	softDeleteOwnerUC *SoftDeleteOwnerUseCase,
) OwnerServiceFacade {
	return &ownerUseCases{
		getOwnerByIdUseCase:    getOwnerByIdUC,
		listOwnersUseCase:      listOwnersUC,
		createOwnerUseCase:     createOwnerUC,
		updateOwnerUseCase:     updateOwnerUC,
		softDeleteOwnerUseCase: softDeleteOwnerUC,
	}
}

func (uc *ownerUseCases) GetOwnerById(ctx context.Context, ownerId int) (ownerDTOs.OwnerResponse, error) {
	return uc.getOwnerByIdUseCase.Execute(ctx, ownerId)
}

func (uc *ownerUseCases) ListOwners(ctx context.Context, params ownerDTOs.GetOwnersRequest) (page.Page[[]ownerDTOs.OwnerResponse], error) {
	return uc.listOwnersUseCase.Execute(ctx, params)
}

func (uc *ownerUseCases) CreateOwner(ctx context.Context, data ownerDTOs.OwnerCreate) (ownerDTOs.OwnerResponse, error) {
	return uc.createOwnerUseCase.Execute(ctx, data)
}

func (uc *ownerUseCases) UpdateOwner(ctx context.Context, ownerId int, data ownerDTOs.OwnerUpdate) (ownerDTOs.OwnerResponse, error) {
	return uc.updateOwnerUseCase.Execute(ctx, ownerId, data)
}

func (uc *ownerUseCases) SoftDeleteOwner(ctx context.Context, ownerId int) error {
	return uc.softDeleteOwnerUseCase.Execute(ctx, ownerId)
}
