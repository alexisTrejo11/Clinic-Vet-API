package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/pets/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type PetUseCases interface {
	SearchPets(ctx context.Context, spec specification.PetSpecification) (page.Page[[]dto.PetResponse], error)
	GetPetByID(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error)
	GetPetByIDAndCustomerID(ctx context.Context, petID valueobject.PetID, customerID valueobject.CustomerID) (dto.PetResponse, error)
	ListsPetByCustomerID(ctx context.Context, customerID uint) (page.Page[[]dto.PetResponse], error)

	CreatePet(ctx context.Context, petCreate dto.CreatePetData) error
	UpdatePet(ctx context.Context, petUpdate dto.PetUpdateData) error
	DeletePet(ctx context.Context, petID valueobject.PetID, isSoftDelete bool) error
}

type petUseCase struct {
	petRepository   repository.PetRepository
	ownerRepository repository.CustomerRepository
}

func NewPetUseCases(petRepo repository.PetRepository, ownerRepo repository.CustomerRepository) PetUseCases {
	return &petUseCase{
		petRepository:   petRepo,
		ownerRepository: ownerRepo,
	}
}

func (f *petUseCase) SearchPets(ctx context.Context, spec specification.PetSpecification) (page.Page[[]dto.PetResponse], error) {
	useCase := NewSearchPetsUseCase(f.petRepository)
	return useCase.Execute(ctx, spec)
}

func (f *petUseCase) ListPetByCustomerID(ctx context.Context, customerID valueobject.CustomerID) (page.Page[[]dto.PetResponse], error) {
	useCase := NewListPetsByCustomerIDUseCase(f.petRepository)
	return useCase.Execute(ctx, customerID)
}

func (f *petUseCase) GetPetByIDAndCustomerID(ctx context.Context, petID valueobject.PetID, customerID valueobject.CustomerID) (dto.PetResponse, error) {
	useCase := NewGetPetByIDAndCustomerIDUseCase(f.petRepository)
	return useCase.Execute(ctx, petID, customerID)
}

func (f *petUseCase) CreatePet(ctx context.Context, petCreate dto.CreatePetData) (dto.PetResponse, error) {
	useCase := NewCreatePetUseCase(f.petRepository, f.ownerRepository)
	return useCase.Execute(ctx, petCreate)
}

func (f *petUseCase) UpdatePet(ctx context.Context, petUpdate dto.PetUpdateData) (dto.PetResponse, error) {
	useCase := NewUpdatePetUseCase(f.petRepository, f.ownerRepository)
	return useCase.Execute(ctx, petUpdate)
}

func (f *petUseCase) GetPetByID(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error) {
	useCase := NewGetPetByIDUseCase(f.petRepository)
	return useCase.Execute(ctx, petID)
}

func (f *petUseCase) DeletePet(ctx context.Context, petID valueobject.PetID, isSoftDelete bool) error {
	useCase := NewDeletePetUseCase(f.petRepository)
	return useCase.Execute(ctx, petID, isSoftDelete)
}
