package vetUsecase

import (
	"context"

	dto "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/mappers"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
)

type ListVetUseCase struct {
	vetRepository vetRepo.VeterinarianRepository
}

func NewListVetUseCase(vetRepository vetRepo.VeterinarianRepository) *ListVetUseCase {
	return &ListVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *ListVetUseCase) Execute(ctx context.Context, searchParam map[string]interface{}) ([]dto.VetResponse, error) {
	veterinarianList, err := uc.vetRepository.List(ctx, searchParam)
	if err != nil {
		return []dto.VetResponse{}, err
	}

	vetResponseList := make([]dto.VetResponse, len(veterinarianList))
	for i, vet := range veterinarianList {
		vetResponses := vetMapper.ToResponse(vet)
		vetResponseList[i] = *vetResponses
	}

	return vetResponseList, nil
}
