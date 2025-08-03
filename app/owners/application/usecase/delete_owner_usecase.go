package ownerUsecase

import (
	"context"

	ownerAppErr "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
)

type DeleteOwnerUseCase struct {
	ownerRepo ownerRepository.OwnerRepository
}

func NewDeleteOwnerUseCase(ownerRepo ownerRepository.OwnerRepository) *DeleteOwnerUseCase {
	return &DeleteOwnerUseCase{
		ownerRepo: ownerRepo,
	}
}

func (uc *DeleteOwnerUseCase) Execute(ctx context.Context, id int) error {
	_, err := uc.ownerRepo.ExistsByID(ctx, id)
	if err != nil {
		return ownerAppErr.HandleGetByIdError(err, id)
	}

	if err := uc.ownerRepo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
