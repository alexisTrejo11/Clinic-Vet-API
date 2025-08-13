package ownerUsecase

import (
	"context"

	ownerDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/dtos"
)

type OwnerFacade interface {
	GetOwnerById(ctx context.Context, ownerId int) (*ownerDTOs.OwnerResponse, error)
	ListOwners(ctx context.Context, params ownerDTOs.GetOwnersRequest) (ownerDTOs.OwnerListResponse, error)
	CreateOwner(ctx context.Context, data ownerDTOs.OwnerCreate) (*ownerDTOs.OwnerResponse, error)
	UpdateOwner(ctx context.Context, ownerId int, data ownerDTOs.OwnerUpdate) (*ownerDTOs.OwnerResponse, error)
	DeleteOwner(ctx context.Context, ownerId int) error
}

type ownerUseCases struct {
	getOwnerByIdUseCase *GetOwnerByIdUseCase
	listOwnersUseCase   *ListOwnersUseCase
	createOwnerUseCase  *CreateOwnerUseCase
	updateOwnerUseCase  *UpdateOwnerUseCase
	deleteOwnerUseCase  *DeleteOwnerUseCase
}

func NewOwnerUseCases(
	getOwnerByIdUC *GetOwnerByIdUseCase,
	listOwnersUC *ListOwnersUseCase,
	createOwnerUC *CreateOwnerUseCase,
	updateOwnerUC *UpdateOwnerUseCase,
	deleteOwnerUC *DeleteOwnerUseCase,
) OwnerFacade {
	return &ownerUseCases{
		getOwnerByIdUseCase: getOwnerByIdUC,
		listOwnersUseCase:   listOwnersUC,
		createOwnerUseCase:  createOwnerUC,
		updateOwnerUseCase:  updateOwnerUC,
		deleteOwnerUseCase:  deleteOwnerUC,
	}
}

func (uc *ownerUseCases) GetOwnerById(ctx context.Context, ownerId int) (*ownerDTOs.OwnerResponse, error) {
	return uc.getOwnerByIdUseCase.Execute(ctx, ownerId)
}

func (uc *ownerUseCases) ListOwners(ctx context.Context, params ownerDTOs.GetOwnersRequest) (ownerDTOs.OwnerListResponse, error) {
	return uc.listOwnersUseCase.Execute(ctx, params)
}

func (uc *ownerUseCases) CreateOwner(ctx context.Context, data ownerDTOs.OwnerCreate) (*ownerDTOs.OwnerResponse, error) {
	return uc.createOwnerUseCase.Execute(ctx, data)
}

func (uc *ownerUseCases) UpdateOwner(ctx context.Context, ownerId int, data ownerDTOs.OwnerUpdate) (*ownerDTOs.OwnerResponse, error) {
	return uc.updateOwnerUseCase.Execute(ctx, ownerId, data)
}

func (uc *ownerUseCases) DeleteOwner(ctx context.Context, ownerId int) error {
	return uc.deleteOwnerUseCase.Execute(ctx, ownerId)
}
