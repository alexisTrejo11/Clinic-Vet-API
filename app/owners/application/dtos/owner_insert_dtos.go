package dtos

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

type OwnerCreate struct {
	Photo string `json:"photo"`
	Name  string `json:"name" validate:"required"`
}

type OwnerUpdate struct {
	Id    int32  `json:"id" validate:"required"`
	Photo string `json:"photo"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (ord *OwnerUpdate) ModelToDTO(owner sqlc.Owner) {
	ord.Id = owner.ID
	ord.Name = owner.Name
	ord.Photo = owner.Photo.String
}

func (ord *OwnerUpdate) GetPetsIDs() []int32 {
	var petsIDs []int32

	return petsIDs
}
