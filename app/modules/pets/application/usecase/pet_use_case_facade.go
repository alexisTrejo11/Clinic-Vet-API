package usecase

import (
	"context"

	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/pets/application/dto"
	"clinic-vet-api/app/shared/page"
)

type PetUseCases interface {
	SearchPets(ctx context.Context, spec specification.PetSpecification) (page.Page[dto.PetResponse], error)
	FindPetByID(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error)
	FindPetByIDAndCustomerID(ctx context.Context, petID valueobject.PetID, customerID valueobject.CustomerID) (dto.PetResponse, error)
	FindsPetByCustomerID(ctx context.Context, customerID uint, pageInput page.PageInput) (page.Page[dto.PetResponse], error)

	CreatePet(ctx context.Context, petCreate dto.CreatePetData) (dto.PetResponse, error)
	UpdatePet(ctx context.Context, petUpdate dto.PetUpdateData) (dto.PetResponse, error)
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

func (f *petUseCase) SearchPets(ctx context.Context, spec specification.PetSpecification) (page.Page[dto.PetResponse], error) {
	useCase := NewSearchPetsUseCase(f.petRepository)
	return useCase.Execute(ctx, spec)
}

func (f *petUseCase) FindsPetByCustomerID(ctx context.Context, customerID uint, pageInput page.PageInput) (page.Page[dto.PetResponse], error) {
	useCase := NewFindPetsByCustomerIDUseCase(f.petRepository)
	return useCase.Execute(ctx, valueobject.NewCustomerID(customerID), pageInput)
}

func (f *petUseCase) FindPetByCustomerID(ctx context.Context, customerID valueobject.CustomerID, pageInput page.PageInput) (page.Page[dto.PetResponse], error) {
	useCase := NewFindPetsByCustomerIDUseCase(f.petRepository)
	return useCase.Execute(ctx, customerID, pageInput)
}

func (f *petUseCase) FindPetByIDAndCustomerID(ctx context.Context, petID valueobject.PetID, customerID valueobject.CustomerID) (dto.PetResponse, error) {
	useCase := NewFindPetByIDAndCustomerIDUseCase(f.petRepository)
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

func (f *petUseCase) FindPetByID(ctx context.Context, petID valueobject.PetID) (dto.PetResponse, error) {
	useCase := NewFindPetByIDUseCase(f.petRepository)
	return useCase.Execute(ctx, petID)
}

func (f *petUseCase) DeletePet(ctx context.Context, petID valueobject.PetID, isSoftDelete bool) error {
	useCase := NewDeletePetUseCase(f.petRepository)
	return useCase.Execute(ctx, petID, isSoftDelete)
}
