package vetUsecase

import (
	"context"

	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/mappers"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
)

type CreateVetUseCase struct {
	vetRepository vetRepo.VeterinarianRepository
}

func NewCreateVetUseCase(vetRepository vetRepo.VeterinarianRepository) *CreateVetUseCase {
	return &CreateVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *CreateVetUseCase) Execute(ctx context.Context, vetCreateData dto.VetCreate) (dto.VetResponse, error) {
	newVet, err := vetMapper.FromCreateDTO(vetCreateData)
	if err != nil {
		return dto.VetResponse{}, err
	}

	if err := newVet.ValidateBusinessLogic(); err != nil {
		return dto.VetResponse{}, err
	}

	if err := uc.vetRepository.Save(ctx, newVet); err != nil {
		return dto.VetResponse{}, err
	}

	vetResponse := vetMapper.ToResponse(newVet)
	return *vetResponse, nil
}
