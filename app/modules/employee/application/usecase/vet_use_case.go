package usecase

import (
	"context"
	"strconv"

	s "github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/veterinarians/application/dto"
	apperror "github.com/alexisTrejo11/Clinic-Vet-API/app/shared/error/application"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type VeterinarianUseCases struct {
	searchVetUseCase  SearchVetUseCase
	getVetByIDUseCase GetVetByIDUseCase
	createVetUseCase  CreateVetUseCase
	updateVetUseCase  UpdateVetUseCase
	deleteVetUseCase  DeleteVetUseCase
}

func NewVetUseCase(
	searchVetUseCase SearchVetUseCase,
	getVetByIDUseCase GetVetByIDUseCase,
	createVetUseCase CreateVetUseCase,
	updateVetUseCase UpdateVetUseCase,
	deleteVetUseCase DeleteVetUseCase,
) *VeterinarianUseCases {
	return &VeterinarianUseCases{
		searchVetUseCase:  searchVetUseCase,
		getVetByIDUseCase: getVetByIDUseCase,
		createVetUseCase:  createVetUseCase,
		updateVetUseCase:  updateVetUseCase,
		deleteVetUseCase:  deleteVetUseCase,
	}
}

func (uc *VeterinarianUseCases) SearchVeterinan(
	ctx context.Context,
	specification s.VetSearchSpecification,
) (page.Page[[]dto.VetResponse], error) {
	return uc.searchVetUseCase.Execute(ctx, specification)
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

func (uc *VeterinarianUseCases) UpdateVetUseCase(ctx context.Context, vetCreateData dto.UpdateVetData) (dto.VetResponse, error) {
	return uc.updateVetUseCase.Execute(ctx, vetCreateData)
}

func (uc *VeterinarianUseCases) DeleteVet(ctx context.Context, vetIDInt int) error {
	vetID, err := valueobject.NewVetID(vetIDInt)
	if err != nil {
		return apperror.FieldValidationError("id", strconv.Itoa(vetIDInt), err.Error())
	}
	return uc.deleteVetUseCase.Execute(ctx, vetID)
}
