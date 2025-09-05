package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type SoftDeleteOwnerUseCase struct {
	ownerRepo repository.OwnerRepository
}

func NewSoftDeleteOwnerUseCase(ownerRepo repository.OwnerRepository) *SoftDeleteOwnerUseCase {
	return &SoftDeleteOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *SoftDeleteOwnerUseCase) Execute(ctx context.Context, id valueobject.OwnerID) error {
	if _, err := uc.ownerRepo.GetByID(ctx, id); err != nil {
		return err
	}

	if err := uc.ownerRepo.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
