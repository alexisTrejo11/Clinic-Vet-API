package usecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
)

type VeterinarianUseCases struct {
	listVetsUseCase   ListVetUseCase
	getVetByIDUseCase GetVetByIDUseCase
	createVetUseCase  CreateVetUseCase
	updateVetUseCase  UpdateVetUseCase
	deleteVetUseCase  DeleteVetUseCase
}

func NewVetUseCase(
	listVetsUseCase ListVetUseCase,
	getVetByIDUseCase GetVetByIDUseCase,
	createVetUseCase CreateVetUseCase,
	updateVetUseCase UpdateVetUseCase,
	deleteVetUseCase DeleteVetUseCase,
) *VeterinarianUseCases {
	return &VeterinarianUseCases{
		listVetsUseCase:   listVetsUseCase,
		getVetByIDUseCase: getVetByIDUseCase,
		createVetUseCase:  createVetUseCase,
		updateVetUseCase:  updateVetUseCase,
		deleteVetUseCase:  deleteVetUseCase,
	}
}

func (uc *VeterinarianUseCases) ListVetUseCase(ctx context.Context, searchParams dto.VetSearchParams) ([]dto.VetResponse, error) {
	return uc.listVetsUseCase.Execute(ctx, searchParams)
}

func (uc *VeterinarianUseCases) GetVetByIDUseCase(ctx context.Context, vetID valueobject.VetID) (dto.VetResponse, error) {
	return uc.getVetByIDUseCase.Execute(ctx, vetID)
}

func (uc *VeterinarianUseCases) CreateVetUseCase(ctx context.Context, vetCreateData dto.VetCreate) (dto.VetResponse, error) {
	return uc.createVetUseCase.Execute(ctx, vetCreateData)
}

func (uc *VeterinarianUseCases) UpdateVetUseCase(ctx context.Context, vetID valueobject.VetID, vetCreateData dto.VetUpdate) (dto.VetResponse, error) {
	return uc.updateVetUseCase.Execute(ctx, vetID, vetCreateData)
}

func (uc *VeterinarianUseCases) DeleteVetUseCase(ctx context.Context, vetID valueobject.VetID) error {
	return uc.deleteVetUseCase.Execute(ctx, vetID)
}
