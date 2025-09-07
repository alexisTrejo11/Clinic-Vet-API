package usecase

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/mapper"
)

type CreateVetUseCase struct {
	vetRepository repository.VetRepository
}

func NewCreateVetUseCase(vetRepository repository.VetRepository) *CreateVetUseCase {
	return &CreateVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *CreateVetUseCase) Execute(ctx context.Context, vetCreateData dto.CreateVetData) (dto.VetResponse, error) {
	vet, err := mapper.FromCreateDTO(vetCreateData)
	if err != nil {
		return dto.VetResponse{}, err
	}
	if err := uc.vetRepository.Save(ctx, vet); err != nil {
		return dto.VetResponse{}, err
	}

	return mapper.ToResponse(vet), nil
}
