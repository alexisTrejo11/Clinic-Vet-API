package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/mapper"
)

type UpdateVetUseCase struct {
	vetRepository repository.VetRepository
}

func NewUpdateVetUseCase(vetRepository repository.VetRepository) *UpdateVetUseCase {
	return &UpdateVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *UpdateVetUseCase) Execute(ctx context.Context, vetId valueobject.VetID, vetUpdateData dto.UpdateVetData) (dto.VetResponse, error) {
	veterinarian, err := uc.vetRepository.GetByID(ctx, vetId)
	if err != nil {
		return dto.VetResponse{}, err
	}

	mapper.UpdateFromDTO(&veterinarian, vetUpdateData)
	if err := uc.vetRepository.Save(ctx, &veterinarian); err != nil {
		return dto.VetResponse{}, err
	}

	return mapper.ToResponse(&veterinarian), nil
}
