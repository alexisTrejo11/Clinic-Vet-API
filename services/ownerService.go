package services

import (
	"context"

	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/mappers"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type OwnerService interface {
	CreateOwner(ownerInsertDTO *dtos.OwnerInsertDTO) error
	GetOwnerById(ownerID int32) (*dtos.OwnerReturnDTO, error)
	ValidateExistingOwner(ownerId int32) bool
	UpdateOwner(ownerUpdateDTO *dtos.OwnerUpdateDTO) error
	DeleteOwner(ownerID int32) error
}

type OwnerServiceImpl struct {
	ownerRepository repository.OwnerRepository
}

func NewOwnerService(ownerRepository repository.OwnerRepository) OwnerService {
	return &OwnerServiceImpl{
		ownerRepository: ownerRepository,
	}
}

func (os OwnerServiceImpl) CreateOwner(ownerInsertDTO *dtos.OwnerInsertDTO) error {
	phone := pgtype.Text{}
	err := phone.Scan(ownerInsertDTO.Phone)
	if err != nil {
		return err
	}

	_, err = os.ownerRepository.CreateOwner(context.Background(), sqlc.CreateOwnerParams{
		Name:  ownerInsertDTO.Name,
		Phone: phone,
	})
	if err != nil {
		return err
	}

	return nil
}

func (os OwnerServiceImpl) GetOwnerById(ownerID int32) (*dtos.OwnerReturnDTO, error) {
	owner, err := os.ownerRepository.GetOwnerByID(context.Background(), ownerID)
	if err != nil {
		return nil, err
	}

	var owneReturnDTO dtos.OwnerReturnDTO
	owneReturnDTO.ModelToDTO(owner)

	return &owneReturnDTO, nil
}

func (os OwnerServiceImpl) UpdateOwner(ownerUpdateDTO *dtos.OwnerUpdateDTO) error {
	owner, _ := os.ownerRepository.GetOwnerByID(context.Background(), ownerUpdateDTO.Id)

	params := mappers.MapOwnerUpdateDtoToEntity(ownerUpdateDTO, owner)

	if err := os.ownerRepository.UpdateOwner(context.Background(), params); err != nil {
		return err
	}

	return nil
}

func (os OwnerServiceImpl) DeleteOwner(ownerID int32) error {
	if err := os.ownerRepository.DeleteOwner(context.Background(), ownerID); err != nil {
		return err
	}
	return nil
}

func (os OwnerServiceImpl) ValidateExistingOwner(ownerId int32) bool {
	isExistingOwner := os.ownerRepository.ValidateExistingOwner(context.Background(), ownerId)
	return isExistingOwner
}
