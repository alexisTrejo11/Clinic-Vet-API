package services

import (
	"context"

	dtos "example.com/at/backend/api-vet/DTOs"
	"example.com/at/backend/api-vet/repository"
	"example.com/at/backend/api-vet/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type OwnerService struct {
	ownerRepository *repository.OwnerRepository
}

func NewOwnerRepository(ownerRepository *repository.OwnerRepository) *OwnerService {
	return &OwnerService{
		ownerRepository: ownerRepository,
	}
}

func (os *OwnerService) CreateOwner(ownerInsertDTO *dtos.OwnerInsertDTO) error {
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

func (os *OwnerService) GetOwnerById(ownerID int32) (*dtos.OwnerReturnDTO, error) {
	// Find Owner
	owner, err := os.ownerRepository.GetOwnerById(context.Background(), ownerID)
	if err != nil {
		return nil, err
	}

	// Convert Into DTO
	var owneReturnDTO dtos.OwnerReturnDTO
	owneReturnDTO.ModelToDTO(owner)

	// Return DTO
	return &owneReturnDTO, nil
}
