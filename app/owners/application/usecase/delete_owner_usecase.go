package ownerUsecase

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
)

type DeleteOwnerUseCase struct {
	ownerRepo ownerDomain.OwnerRepository
}

func NewDeleteOwnerUseCase(ownerRepo ownerDomain.OwnerRepository) *DeleteOwnerUseCase {
	return &DeleteOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *DeleteOwnerUseCase) Execute(ctx context.Context, id int) error {
	_, err := uc.ownerRepo.ExistsByID(ctx, id)
	if err != nil {
		return ownerDomain.HandleGetByIdError(err, id)
	}

	if err := uc.ownerRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
