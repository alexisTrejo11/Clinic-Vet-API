package domain

/*

import (
	"example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
)

type MedicalHistoryService interface {
	CreateMedicalRepository(medicalHistoryInsertDTO DTOs.MedicalHistoryInsertDTO) error
	GetMedicalRepositoryById(medicalHistoryID int32) (*DTOs.MedicalHistoryDTO, error)
	GetMedicalRepositoryByOwner(ownerDTO DTOs.OwnerDTO) ([]DTOs.MedicalHistoryDTO, error)
	GetMedicalRepositoryByPetID(petID int32) ([]DTOs.MedicalHistoryDTO, error)
	UpdateMedicalRepository(medicalHistoryUpdateDTO DTOs.MedicalHistoryUpdateDTO) error
	DeleteMedicalRepositoryById(medicalHistoryID int32) error
}

type medicalHistoryServiceImpl struct {
	medicalHistoryRepository repository.MedicalHistoryRepository
	medHistMappers           mappers.MedicalHistoryMappers
}

func NewMedicalHistoryService(medicalHistoryRepository repository.MedicalHistoryRepository) MedicalHistoryService {
	return &medicalHistoryServiceImpl{
		medicalHistoryRepository: medicalHistoryRepository,
	}
}

func (mhs medicalHistoryServiceImpl) CreateMedicalRepository(medicalHistoryInsertDTO DTOs.MedicalHistoryInsertDTO) error {
	medHistCreateParams := mhs.medHistMappers.InsertDtoCreateParams(medicalHistoryInsertDTO)

	if err := mhs.medicalHistoryRepository.CreateMedicalHistory(medHistCreateParams); err != nil {
		return err
	}

	return nil
}

func (mhs medicalHistoryServiceImpl) GetMedicalRepositoryById(medicalHistoryID int32) (*DTOs.MedicalHistoryDTO, error) {
	medicalHistory, err := mhs.medicalHistoryRepository.GetMedicalHistoryByID(medicalHistoryID)
	if err != nil {
		return nil, err
	}

	medicalHistoryDTO := mhs.medHistMappers.SqlcEntityToDTO(*medicalHistory)

	return &medicalHistoryDTO, nil
}

func (mhs medicalHistoryServiceImpl) GetMedicalRepositoryByVetID(vetID int32) ([]DTOs.MedicalHistoryDTO, error) {
	medicalHistories, err := mhs.medicalHistoryRepository.GetMedicalHistoryByVetID(vetID)
	if err != nil {
		return nil, err
	}

	var medicalHistoriesDTOs []DTOs.MedicalHistoryDTO
	for _, medicalHistory := range medicalHistories {
		medicalHistoryDTO := mhs.medHistMappers.SqlcEntityToDTO(medicalHistory)
		medicalHistoriesDTOs = append(medicalHistoriesDTOs, medicalHistoryDTO)
	}

	return medicalHistoriesDTOs, nil
}

func (mhs medicalHistoryServiceImpl) GetMedicalRepositoryByPetID(petID int32) ([]DTOs.MedicalHistoryDTO, error) {
	medicalHistories, err := mhs.medicalHistoryRepository.GetMedicalHistoryByPetID(petID)
	if err != nil {
		return nil, err
	}

	var medicalHistoriesDTOs []DTOs.MedicalHistoryDTO
	for _, medicalHistory := range medicalHistories {
		medicalHistoryDTO := mhs.medHistMappers.SqlcEntityToDTO(medicalHistory)
		medicalHistoriesDTOs = append(medicalHistoriesDTOs, medicalHistoryDTO)
	}

	return medicalHistoriesDTOs, nil
}

func (mhs medicalHistoryServiceImpl) GetMedicalRepositoryByOwner(ownerDTO DTOs.OwnerDTO) ([]DTOs.MedicalHistoryDTO, error) {
	petsIDs := ownerDTO.GetPetsIDs()

	var medicalHistoriesDTOs []DTOs.MedicalHistoryDTO

	for _, petId := range petsIDs {
		// Get Pet Medical Histories
		medicalHistories, _ := mhs.medicalHistoryRepository.GetMedicalHistoryByPetID(petId)
		for _, medicalHistory := range medicalHistories {
			// For Each Medical History Per Pet Map Into DTO and append to List
			medicalHistoryDTO := mhs.medHistMappers.SqlcEntityToDTO(medicalHistory)
			medicalHistoriesDTOs = append(medicalHistoriesDTOs, medicalHistoryDTO)
		}

	}

	return medicalHistoriesDTOs, nil
}

func (mhs medicalHistoryServiceImpl) UpdateMedicalRepository(medicalHistoryUpdateDTO DTOs.MedicalHistoryUpdateDTO) error {
	medHistUpdateParams := mhs.medHistMappers.UpdateDtoUpdateParams(medicalHistoryUpdateDTO)

	if err := mhs.medicalHistoryRepository.UpdateMedicalHistory(medHistUpdateParams); err != nil {
		return err
	}

	return nil
}

func (mhs medicalHistoryServiceImpl) DeleteMedicalRepositoryById(medicalHistoryID int32) error {
	err := mhs.medicalHistoryRepository.DeleteMedicalHistory(medicalHistoryID)
	if err != nil {
		return err
	}
	return nil
}
*/
