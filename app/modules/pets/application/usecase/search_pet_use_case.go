package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/mapper"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type SearchPetUseCase struct {
	petRepo repository.PetRepository
}

func NewSearchPetsUseCase(petRepo repository.PetRepository) *SearchPetUseCase {
	return &SearchPetUseCase{petRepo: petRepo}
}

func (uc *SearchPetUseCase) Execute(ctx context.Context, spec specification.PetSpecification) (page.Page[[]dto.PetResponse], error) {
	petPage, err := uc.petRepo.Search(ctx, spec)
	if err != nil {
		return page.Page[[]dto.PetResponse]{}, err
	}
	return mapper.ToResponsesPage(petPage), nil
}
