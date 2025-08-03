package vetUsecase

import (
	"context"

	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/mappers"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
)

type GetVetByIdUseCase struct {
	vetRepository vetRepo.VeterinarianRepository
}

func NewGetVetByIdUseCase(vetRepository vetRepo.VeterinarianRepository) *GetVetByIdUseCase {
	return &GetVetByIdUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *GetVetByIdUseCase) Execute(ctx context.Context, vetId int) (dto.VetResponse, error) {
	veterinarian, err := uc.vetRepository.GetByID(ctx, vetId)
	if err != nil {
		return dto.VetResponse{}, err
	}

	vetResponse := vetMapper.ToResponse(veterinarian)
	return *vetResponse, nil
}
