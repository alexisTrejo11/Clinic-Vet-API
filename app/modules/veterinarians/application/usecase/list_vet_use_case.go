package usecase

import (
	"context"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/mapper"
)

type ListVetUseCase struct {
	vetRepository repository.VetRepository
}

func NewListVetUseCase(vetRepository repository.VetRepository) *ListVetUseCase {
	return &ListVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *ListVetUseCase) Execute(ctx context.Context, searchParam dto.VetSearchParams) ([]dto.VetResponse, error) {
	veterinarianList, err := uc.vetRepository.List(ctx, searchParam)
	if err != nil {
		return []dto.VetResponse{}, err
	}

	vetResponseList := make([]dto.VetResponse, len(veterinarianList))
	for i, vet := range veterinarianList {
		vetResponses := mapper.ToResponse(&vet)
		vetResponseList[i] = vetResponses
	}

	return vetResponseList, nil
}
