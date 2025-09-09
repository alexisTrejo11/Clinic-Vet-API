package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
)

type PetUseCasesFacade interface {
	SearchPets(ctx context.Context, params any) ([]dto.PetResponse, error)
	GetPetByID(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error)
	CreatePet(ctx context.Context, petCreate dto.CreatePetData) (dto.PetResponse, error)
	UpdatePet(ctx context.Context, petUpdate dto.PetUpdate) (dto.PetResponse, error)
	DeletePet(ctx context.Context, petID valueobject.PetID, isSoftDelete bool) error
}

type petUseCaseFacade struct {
	petRepository   repository.PetRepository
	ownerRepository repository.OwnerRepository
}

func NewPetUseCasesFacade(petRepo repository.PetRepository, ownerRepo repository.OwnerRepository) PetUseCasesFacade {
	return &petUseCaseFacade{
		petRepository:   petRepo,
		ownerRepository: ownerRepo,
	}
}

func (f *petUseCaseFacade) SearchPets(ctx context.Context, params any) ([]dto.PetResponse, error) {
	return []dto.PetResponse{}, nil
}

func (f *petUseCaseFacade) CreatePet(ctx context.Context, petCreate dto.CreatePetData) (dto.PetResponse, error) {
	useCase := NewCreatePetUseCase(f.petRepository, f.ownerRepository)
	return useCase.Execute(ctx, petCreate)
}

func (f *petUseCaseFacade) UpdatePet(ctx context.Context, petUpdate dto.PetUpdate) (dto.PetResponse, error) {
	useCase := NewUpdatePetUseCase(f.petRepository, f.ownerRepository)
	return useCase.Execute(ctx, petUpdate)
}

func (f *petUseCaseFacade) GetPetByID(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error) {
	useCase := NewGetPetByIDUseCase(f.petRepository)
	return useCase.Execute(ctx, petID)
}

func (f *petUseCaseFacade) DeletePet(ctx context.Context, petID valueobject.PetID, isSoftDelete bool) error {
	useCase := NewDeletePetUseCase(f.petRepository)
	return useCase.Execute(ctx, petID, isSoftDelete)
}
