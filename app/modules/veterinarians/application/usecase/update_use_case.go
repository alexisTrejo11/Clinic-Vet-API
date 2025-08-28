package vetUsecase

import (
	"context"
	"strconv"

	vetApplication "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application"
	vetDtos "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/dtos"
	vetMapper "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/mappers"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
)

type UpdateVetUseCase struct {
	vetRepository vetRepo.VeterinarianRepository
}

func NewUpdateVetUseCase(vetRepository vetRepo.VeterinarianRepository) *UpdateVetUseCase {
	return &UpdateVetUseCase{
		vetRepository: vetRepository,
	}
}

func (uc *UpdateVetUseCase) Execute(ctx context.Context, vetId int, vetUpdateData vetDtos.VetUpdate) (vetDtos.VetResponse, error) {
	veterinarian, err := uc.vetRepository.GetById(ctx, vetId)
	if err != nil {
		return vetDtos.VetResponse{}, vetApplication.VetNotFoundErr("id", strconv.Itoa(vetId))
	}

	vetMapper.UpdateFromDTO(&veterinarian, vetUpdateData)
	if err := uc.vetRepository.Save(ctx, &veterinarian); err != nil {
		return vetDtos.VetResponse{}, vetApplication.VetDBErr("update", err)

	}

	return vetDtos.VetResponse{}, nil
}
