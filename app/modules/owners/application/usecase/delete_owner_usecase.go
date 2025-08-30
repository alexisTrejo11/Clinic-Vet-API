package usecase

import (
	"context"

	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/errors"
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

func (uc *SoftDeleteOwnerUseCase) Execute(ctx context.Context, id int) error {
	if exists, err := uc.ownerRepo.ExistsByID(ctx, id); err != nil {
		return err
	} else if !exists {
		return domainerr.HandleGetByIdError(err, id)
	}

	if err := uc.ownerRepo.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
