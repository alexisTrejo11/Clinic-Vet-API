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

func (uc *VeterinarianUseCases) ListVetUseCase(ctx context.Context, limit, offset int) ([]vetDtos.VetResponse, error) {
	return uc.listVetsUseCase.Execute(ctx, limit, offset)
}

func (uc *VeterinarianUseCases) GetVetByIdUseCase(ctx context.Context, vetId uint) (vetDtos.VetResponse, error) {
	return uc.getVetByIdUseCase.Execute(ctx, vetId)

}

func (uc *VeterinarianUseCases) CreateVetUseCase(ctx context.Context, vetCreateData vetDtos.VetCreate) (vetDtos.VetResponse, error) {
	return uc.createVetUseCase.Execute(ctx, vetCreateData)
}
func (uc *VeterinarianUseCases) UpdateVetUseCase(ctx context.Context, vetId uint, vetCreateData vetDtos.VetUpdate) (vetDtos.VetResponse, error) {
	return uc.updateVetUseCase.Execute(ctx, vetId, vetCreateData)
}

func (uc *VeterinarianUseCases) DeleteVetUseCase(ctx context.Context, vetId uint) error {
	return uc.deleteVetUseCase.Execute(ctx, vetId)
}
