package usecase

import (
	"context"
	"strconv"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
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

func (uc *VeterinarianUseCases) GetVetByIDUseCase(ctx context.Context, vetIDInt int) (dto.VetResponse, error) {
	vetID, err := valueobject.NewVetID(vetIDInt)
	if err != nil {
		return dto.VetResponse{}, apperror.FieldValidationError("id", strconv.Itoa(vetIDInt), err.Error())
	}
	return uc.getVetByIDUseCase.Execute(ctx, vetID)
}

func (uc *VeterinarianUseCases) CreateVetUseCase(ctx context.Context, vetCreateData dto.CreateVetData) (dto.VetResponse, error) {
	return uc.createVetUseCase.Execute(ctx, vetCreateData)
}

func (uc *VeterinarianUseCases) UpdateVetUseCase(ctx context.Context, vetIDInt int, vetCreateData dto.UpdateVetData) (dto.VetResponse, error) {
	vetID, err := valueobject.NewVetID(vetIDInt)
	if err != nil {
		return dto.VetResponse{}, apperror.FieldValidationError("id", strconv.Itoa(vetIDInt), err.Error())
	}

	return uc.updateVetUseCase.Execute(ctx, vetID, vetCreateData)
}

func (uc *VeterinarianUseCases) DeleteVetUseCase(ctx context.Context, vetIDInt int) error {
	vetID, err := valueobject.NewVetID(vetIDInt)
	if err != nil {
		return apperror.FieldValidationError("id", strconv.Itoa(vetIDInt), err.Error())
	}
	return uc.deleteVetUseCase.Execute(ctx, vetID)
}
