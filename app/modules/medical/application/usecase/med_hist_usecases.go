package usecase

import (
	"context"
	"errors"

	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/modules/medical/application/dto"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type MedicalHistoryUseCase struct {
	medHistRepo repository.MedicalHistoryRepository
	ownerRepo   repository.OwnerRepository
	vetRepo     repository.VetRepository
	petRepo     repository.PetRepository
}

func NewMedicalHistoryUseCase(
	medHistRepo repository.MedicalHistoryRepository,
	ownerRepo repository.OwnerRepository,
	vetRepo repository.VetRepository,
	petRepo repository.PetRepository,
) *MedicalHistoryUseCase {
	return &MedicalHistoryUseCase{
		medHistRepo: medHistRepo,
		ownerRepo:   ownerRepo,
		vetRepo:     vetRepo,
		petRepo:     petRepo,
	}
}

func (uc *MedicalHistoryUseCase) Search(ctx context.Context, searchParams dto.MedHistSearchParams) (page.Page[[]dto.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.Search(ctx, searchParams)
	if err != nil {
		return page.Page[[]dto.MedHistResponse]{}, err
	}

	responsePage := page.NewPage(dto.ListToResponse(medHistPage.Data), medHistPage.Metadata)
	return responsePage, nil
}

func (uc *MedicalHistoryUseCase) GetByPet(ctx context.Context, petID int, pagintation page.PageData) (page.Page[[]dto.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.ListByPetID(ctx, petID, pagintation)
	if err != nil {
		return page.Page[[]dto.MedHistResponse]{}, err
	}

	medHistResponse := dto.ListToResponse(medHistPage.Data)
	return page.NewPage(medHistResponse, medHistPage.Metadata), nil
}

func (uc *MedicalHistoryUseCase) ListByOwner(ctx context.Context, ownerID int, pagintation page.PageData) (page.Page[[]dto.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.ListByOwnerID(ctx, ownerID, pagintation)
	if err != nil {
		return page.Page[[]dto.MedHistResponse]{}, err
	}

	medHistResponse := dto.ListToResponse(medHistPage.Data)
	reponsePage := page.NewPage(medHistResponse, medHistPage.Metadata)
	return reponsePage, nil
}

func (uc *MedicalHistoryUseCase) ListByVet(ctx context.Context, vetID int, pagintation page.PageData) (page.Page[[]dto.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.ListByVetID(ctx, vetID, pagintation)
	if err != nil {
		return page.Page[[]dto.MedHistResponse]{}, err
	}

	medHistResponse := dto.ListToResponse(medHistPage.Data)
	return page.NewPage(medHistResponse, medHistPage.Metadata), nil
}

func (uc *MedicalHistoryUseCase) GetByID(ctx context.Context, id int) (dto.MedHistResponse, error) {
	medHistory, err := uc.medHistRepo.GetByID(ctx, id)
	if err != nil {
		return dto.MedHistResponse{}, err
	}

	return dto.ToResponse(*medHistory), nil
}

func (uc *MedicalHistoryUseCase) GetByIDWithDeatils(ctx context.Context, id int) (dto.MedHistResponseDetail, error) {
	medHistory, err := uc.medHistRepo.GetByID(ctx, id)
	if err != nil {
		return dto.MedHistResponseDetail{}, err
	}

	owner, err := uc.ownerRepo.GetByID(ctx, medHistory.OwnerID())
	if err != nil {
		return dto.MedHistResponseDetail{}, err
	}

	pet, err := uc.petRepo.GetByID(ctx, medHistory.PetID().GetValue())
	if err != nil {
		return dto.MedHistResponseDetail{}, err
	}

	vet, err := uc.vetRepo.GetByID(ctx, medHistory.VetID())
	if err != nil {
		return dto.MedHistResponseDetail{}, err
	}

	return dto.ToResponseDetail(*medHistory, owner, vet, pet), nil
}

func (uc *MedicalHistoryUseCase) Create(ctx context.Context, createData dto.MedicalHistoryCreate) error {
	medHistory, err := dto.FromCreateDTO(createData)
	if err != nil {
		return err
	}

	if err := uc.validateCreation(ctx, &createData); err != nil {
		return err
	}

	if err := medHistory.ValidateBusinessRules(); err != nil {
		return err
	}

	if err := uc.medHistRepo.Create(ctx, &medHistory); err != nil {
		return err
	}

	return nil
}

func (uc *MedicalHistoryUseCase) Update(ctx context.Context, id int, dto dto.MedicalHistoryUpdate) error {
	medHistory, err := uc.medHistRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if err := uc.validateUpdate(ctx, &dto); err != nil {
		return err
	}

	medHistoryUpdated, err := dto.FromUpdateDTO(dto, *medHistory)
	if err != nil {
		return err
	}

	if err := medHistoryUpdated.ValidateBusinessRules(); err != nil {
		return err
	}

	if err := uc.medHistRepo.Update(ctx, &medHistoryUpdated); err != nil {
		return err
	}

	return nil
}

func (uc *MedicalHistoryUseCase) Delete(ctx context.Context, id int, isSoftDelete bool) error {
	_, err := uc.medHistRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return uc.medHistRepo.Delete(ctx, id, isSoftDelete)
}

func (uc *MedicalHistoryUseCase) validateCreation(ctx context.Context, createData *dto.MedicalHistoryCreate) error {
	owner, err := uc.ownerRepo.GetByID(ctx, medHistory.OwnerID)
	if err != nil {
		return err
	}

	petFound := false
	for _, pet := range owner.Pets() {
		if pet.GetID().GetValue() == medHistory.PetID {
			createData.PetID = pet.GetID().GetValue()
			petFound = true
			break
		}
	}

	if !petFound {
		return errors.New("pet not found in owner's pets")
	}

	if _, err := uc.vetRepo.GetByID(ctx, createData.VetID); err != nil {
		return err
	}

	return nil
}

func (uc *MedicalHistoryUseCase) validateUpdate(ctx context.Context, updateData *dto.MedicalHistoryUpdate) error {
	if updateData.OwnerID != nil {
		owner, err := uc.ownerRepo.GetByID(ctx, *updateData.OwnerID)
		if err != nil {
			return err
		}

		if updateData.PetID == nil {
			petFound := false
			for _, pet := range owner.Pets() {
				if pet.GetID() == *updateData.PetID {
					petFound = true
					break
				}
			}

			if !petFound {
				return errors.New("pet not found in owner's pets")
			}
		}

	}

	if updateData.VetID != nil {
		if _, err := uc.vetRepo.GetByID(ctx, *updateData.VetID); err != nil {
			return err
		}
	}

	return nil
}
