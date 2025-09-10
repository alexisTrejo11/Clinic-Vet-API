package usecase

import (
	"context"

	v "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/mapper"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchVetUseCase struct {
	vetRepository repository.VetRepository
}

func NewListVetUseCase(vetRepository repository.VetRepository) *SearchVetUseCase {
	return &SearchVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *SearchVetUseCase) Execute(ctx context.Context, specification v.VetSearchSpecification) (page.Page[[]dto.VetResponse], error) {
	vetPage, err := uc.vetRepository.Search(ctx, specification)
	if err != nil {
		return page.Page[[]dto.VetResponse]{}, err
	}

	if len(vetPage.Data) == 0 {
		return page.EmptyPage[[]dto.VetResponse](), nil
	}

	vetResponses := make([]dto.VetResponse, len(vetPage.Data))
	for i, vet := range vetPage.Data {
		vetResponse := mapper.ToResponse(&vet)
		vetResponses[i] = vetResponse
	}

	return page.NewPage(vetResponses, vetPage.Metadata), nil
}
