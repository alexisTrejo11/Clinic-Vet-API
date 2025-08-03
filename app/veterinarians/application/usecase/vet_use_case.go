package vetUsecase

import (
	"context"

	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
)

type VeterinarianUseCases struct {
	listVetsUseCase   ListVetUseCase
	getVetByIdUseCase GetVetByIdUseCase
	createVetUseCase  CreateVetUseCase
	updateVetUseCase  UpdateVetUseCase
	deleteVetUseCase  DeleteVetUseCase
}

func NewVetUseCase(
	listVetsUseCase ListVetUseCase,
	getVetByIdUseCase GetVetByIdUseCase,
	createVetUseCase CreateVetUseCase,
	updateVetUseCase UpdateVetUseCase,
	deleteVetUseCase DeleteVetUseCase,
) *VeterinarianUseCases {
	return &VeterinarianUseCases{
		listVetsUseCase:   listVetsUseCase,
		getVetByIdUseCase: getVetByIdUseCase,
		createVetUseCase:  createVetUseCase,
		updateVetUseCase:  updateVetUseCase,
		deleteVetUseCase:  deleteVetUseCase,
	}
}

func (uc *VeterinarianUseCases) ListVetUseCase(ctx context.Context, searchParams vetDtos.VetSearchParams) ([]vetDtos.VetResponse, error) {
	return uc.listVetsUseCase.Execute(ctx, searchParams)
}

func (uc *VeterinarianUseCases) GetVetByIdUseCase(ctx context.Context, vetId int) (vetDtos.VetResponse, error) {
	return uc.getVetByIdUseCase.Execute(ctx, vetId)
}

func (uc *VeterinarianUseCases) CreateVetUseCase(ctx context.Context, vetCreateData vetDtos.VetCreate) (vetDtos.VetResponse, error) {
	return uc.createVetUseCase.Execute(ctx, vetCreateData)
}

func (uc *VeterinarianUseCases) UpdateVetUseCase(ctx context.Context, vetId int, vetCreateData vetDtos.VetUpdate) (vetDtos.VetResponse, error) {
	return uc.updateVetUseCase.Execute(ctx, vetId, vetCreateData)
}

func (uc *VeterinarianUseCases) DeleteVetUseCase(ctx context.Context, vetId int) error {
	return uc.deleteVetUseCase.Execute(ctx, vetId)
}
