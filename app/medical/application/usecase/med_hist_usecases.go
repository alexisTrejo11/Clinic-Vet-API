package medHistUsecases

import (
	"context"
	"errors"

	mhDTOs "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/application/dtos"
	medHistRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/medical/domain/repositories"
	ownerRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/owners/application/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	vetRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/application/repositories"
)

type MedicalHistoryUseCase struct {
	medHistRepo medHistRepo.MedicalHistoryRepository
	ownerRepo   ownerRepository.OwnerRepository
	vetRepo     vetRepo.VeterinarianRepository
}

func NewMedicalHistoryUseCase(medHistRepo medHistRepo.MedicalHistoryRepository, ownerRepo ownerRepository.OwnerRepository, vetRepo vetRepo.VeterinarianRepository) *MedicalHistoryUseCase {
	return &MedicalHistoryUseCase{medHistRepo: medHistRepo, ownerRepo: ownerRepo, vetRepo: vetRepo}
}

func (uc *MedicalHistoryUseCase) Search(ctx context.Context, searchParams mhDTOs.MedHistSearchParams) (page.Page[[]mhDTOs.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.Search(ctx, searchParams)
	if err != nil {
		return page.Page[[]mhDTOs.MedHistResponse]{}, err
	}

	reponsePage := page.NewPage(mhDTOs.ListToResponse(medHistPage.Data), medHistPage.Metadata)

	return *reponsePage, nil
}

func (uc *MedicalHistoryUseCase) GetByPet(ctx context.Context, petId int, pagintation page.PageData) (page.Page[[]mhDTOs.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.ListByPetId(ctx, petId, pagintation)
	if err != nil {
		return page.Page[[]mhDTOs.MedHistResponse]{}, err
	}

	reponsePage := page.NewPage(mhDTOs.ListToResponse(medHistPage.Data), medHistPage.Metadata)
	return *reponsePage, nil
}

func (uc *MedicalHistoryUseCase) ListByOwner(ctx context.Context, ownerId int, pagintation page.PageData) (page.Page[[]mhDTOs.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.ListByOwnerId(ctx, ownerId, pagintation)
	if err != nil {
		return page.Page[[]mhDTOs.MedHistResponse]{}, err
	}

	reponsePage := page.NewPage(mhDTOs.ListToResponse(medHistPage.Data), medHistPage.Metadata)
	return *reponsePage, nil
}

func (uc *MedicalHistoryUseCase) ListByVet(ctx context.Context, vetId int, pagintation page.PageData) (page.Page[[]mhDTOs.MedHistResponse], error) {
	medHistPage, err := uc.medHistRepo.ListByVetId(ctx, vetId, pagintation)
	if err != nil {
		return page.Page[[]mhDTOs.MedHistResponse]{}, err
	}

	reponsePage := page.NewPage(mhDTOs.ListToResponse(medHistPage.Data), medHistPage.Metadata)
	return *reponsePage, nil
}

func (uc *MedicalHistoryUseCase) GetById(ctx context.Context, id int) (mhDTOs.MedHistResponse, error) {
	medHistory, err := uc.medHistRepo.GetById(ctx, id)
	if err != nil {
		return mhDTOs.MedHistResponse{}, err
	}

	return mhDTOs.ToResponse(*medHistory), nil
}

func (uc *MedicalHistoryUseCase) GetByIdWithDeatils(ctx context.Context, id int) (mhDTOs.MedHistResponseDetail, error) {
	medHistory, err := uc.medHistRepo.GetById(ctx, id)
	if err != nil {
		return mhDTOs.MedHistResponseDetail{}, err
	}

	owner, err := uc.ownerRepo.GetById(ctx, medHistory.OwnerId, true)
	if err != nil {
		return mhDTOs.MedHistResponseDetail{}, err
	}

	pet, err := owner.GetPetById(medHistory.PetId.GetValue())
	if err != nil {
		return mhDTOs.MedHistResponseDetail{}, err
	}

	vet, err := uc.vetRepo.GetById(ctx, medHistory.VetId.GetValue())
	if err != nil {
		return mhDTOs.MedHistResponseDetail{}, err
	}

	return mhDTOs.ToResponseDetail(*medHistory, owner, vet, *pet), nil
}

func (uc *MedicalHistoryUseCase) Create(ctx context.Context, dto mhDTOs.MedicalHistoryCreate) error {
	medHistory, err := mhDTOs.FromCreateDTO(dto)
	if err != nil {
		return err
	}

	if err := uc.valIdateCreation(ctx, &dto); err != nil {
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

func (uc *MedicalHistoryUseCase) Update(ctx context.Context, id int, dto mhDTOs.MedicalHistoryUpdate) error {
	medHistory, err := uc.medHistRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if err := uc.valIdateUpdate(ctx, &dto); err != nil {
		return err
	}

	medHistoryUpdated, err := mhDTOs.FromUpdateDTO(dto, *medHistory)
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
	_, err := uc.medHistRepo.GetById(ctx, id)
	if err != nil {
		return err
	}

	return uc.medHistRepo.Delete(ctx, id, isSoftDelete)
}

func (uc *MedicalHistoryUseCase) valIdateCreation(ctx context.Context, medHistory *mhDTOs.MedicalHistoryCreate) error {
	owner, err := uc.ownerRepo.GetById(ctx, medHistory.OwnerId, true)
	if err != nil {
		return err
	}

	petFound := false
	for _, pet := range owner.Pets {
		if pet.Id == medHistory.PetId {
			medHistory.PetId = pet.Id
			petFound = true
			break
		}
	}

	if !petFound {
		return errors.New("pet not found in owner's pets")
	}

	if _, err := uc.vetRepo.GetById(ctx, medHistory.VetId); err != nil {
		return err
	}

	return nil
}

func (uc *MedicalHistoryUseCase) valIdateUpdate(ctx context.Context, medHistory *mhDTOs.MedicalHistoryUpdate) error {
	if medHistory.OwnerId != nil {
		owner, err := uc.ownerRepo.GetById(ctx, *medHistory.OwnerId, true)
		if err != nil {
			return err
		}

		if medHistory.PetId == nil {
			petFound := false
			for _, pet := range owner.Pets {
				if pet.Id == *medHistory.PetId {
					petFound = true
					break
				}
			}

			if !petFound {
				return errors.New("pet not found in owner's pets")
			}
		}

	}

	if medHistory.VetId != nil {
		if _, err := uc.vetRepo.GetById(ctx, *medHistory.VetId); err != nil {
			return err
		}
	}

	return nil
}
