package ownerUsecase

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
)

type SoftDeleteOwnerUseCase struct {
	ownerRepo ownerDomain.OwnerRepository
}

func NewSoftDeleteOwnerUseCase(ownerRepo ownerDomain.OwnerRepository) *SoftDeleteOwnerUseCase {
	return &SoftDeleteOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *SoftDeleteOwnerUseCase) Execute(ctx context.Context, id int) error {
	if exists, err := uc.ownerRepo.ExistsByID(ctx, id); err != nil {
		return err
	} else if !exists {
		return ownerDomain.HandleGetByIdError(err, id)
	}

	if err := uc.ownerRepo.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
