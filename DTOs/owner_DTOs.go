package DTOs

import (
	"example.com/at/backend/api-vet/sqlc"
)

type OwnerInsertDTO struct {
	Photo string `json:"photo"`
	Name  string `json:"name" validate:"required"`
}

type OwnerDTO struct {
	Id    int32    `json:"id"`
	Photo string   `json:"photo"`
	Name  string   `json:"name" validate:"required"`
	Phone string   `json:"phone"`
	Pets  []PetDTO `json:"pets"`
}

type OwnerUpdateDTO struct {
	Id    int32  `json:"id" validate:"required"`
	Photo string `json:"photo"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (ord *OwnerDTO) ModelToDTO(owner sqlc.Owner) {
	ord.Id = owner.ID
	ord.Name = owner.Name
	ord.Photo = owner.Photo.String
}

func (ord *OwnerDTO) GetPetsIDs() []int32 {
	var petsIDs []int32
	pets := ord.Pets

	for _, pets := range pets {
		petsIDs = append(petsIDs, pets.Id)
	}

	return petsIDs
}
