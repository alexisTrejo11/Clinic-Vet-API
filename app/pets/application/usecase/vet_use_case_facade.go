package petUsecase

import (
	"context"

	ownerDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/domain"
	petDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/application/usecase/dtos"
	petDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/pets/domain"
)

type PetUseCasesFacade interface {
	GetPetById(ctx context.Context, petId int) (petDTOs.PetResponse, error)
	ListPets(ctx context.Context) ([]petDTOs.PetResponse, error)
	CreatePet(ctx context.Context, petCreate petDTOs.PetCreate) (petDTOs.PetResponse, error)
	UpdatePet(ctx context.Context, petUpdate petDTOs.PetUpdate) (petDTOs.PetResponse, error)
	DeletePet(ctx context.Context, petId int, isSoftDelete bool) error
}

type petUseCaseFacade struct {
	petRepository   petDomain.PetRepository
	ownerRepository ownerDomain.OwnerRepository
}

func NewPetUseCasesFacade(petRepo petDomain.PetRepository, ownerRepo ownerDomain.OwnerRepository) PetUseCasesFacade {
	return &petUseCaseFacade{
		petRepository:   petRepo,
		ownerRepository: ownerRepo,
	}
}

func (f *petUseCaseFacade) CreatePet(ctx context.Context, petCreate petDTOs.PetCreate) (petDTOs.PetResponse, error) {
	useCase := NewCreatePetUseCase(f.petRepository, f.ownerRepository)
	return useCase.Execute(ctx, petCreate)
}

func (f *petUseCaseFacade) UpdatePet(ctx context.Context, petUpdate petDTOs.PetUpdate) (petDTOs.PetResponse, error) {
	useCase := NewUpdatePetUseCase(f.petRepository, f.ownerRepository)
	return useCase.Execute(ctx, petUpdate)
}

func (f *petUseCaseFacade) GetPetById(ctx context.Context, petId int) (petDTOs.PetResponse, error) {
	useCase := NewGetPetByIdUseCase(f.petRepository)
	return useCase.Execute(ctx, petId)
}

func (f *petUseCaseFacade) DeletePet(ctx context.Context, petId int, isSoftDelete bool) error {
	useCase := NewDeletePetUseCase(f.petRepository)
	return useCase.Execute(ctx, petId, isSoftDelete)
}

func (f *petUseCaseFacade) ListPets(ctx context.Context) ([]petDTOs.PetResponse, error) {
	useCase := NewListPetsUseCase(f.petRepository)
	return useCase.Execute(ctx)
}
