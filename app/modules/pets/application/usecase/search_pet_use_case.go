package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/modules/pets/application/mapper"
	"clinic-vet-api/app/shared/page"
)

type SearchPetUseCase struct {
	petRepo repository.PetRepository
}

func NewSearchPetsUseCase(petRepo repository.PetRepository) *SearchPetUseCase {
	return &SearchPetUseCase{petRepo: petRepo}
}

func (uc *SearchPetUseCase) Execute(ctx context.Context, spec specification.PetSpecification) (page.Page[dto.PetResponse], error) {
	petPage, err := uc.petRepo.FindBySpecification(ctx, spec)
	if err != nil {
		return page.Page[dto.PetResponse]{}, err
	}
	return mapper.ToResponsesPage(petPage), nil
}
