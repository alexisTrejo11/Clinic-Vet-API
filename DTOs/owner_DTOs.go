package DTOs

import "example.com/at/backend/api-vet/sqlc"

type OwnerInsertDTO struct {
	Photo string `json:"photo"`
	Name  string `json:"name" validate:"required"`
}

type OwnerReturnDTO struct {
	Id    int32  `json:"id"`
	Photo string `json:"photo"`
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone"`
}

type OwnerUpdateDTO struct {
	Id    int32  `json:"id" validate:"required"`
	Photo string `json:"photo"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (ord *OwnerReturnDTO) ModelToDTO(owner sqlc.Owner) {
	ord.Id = owner.ID
	ord.Name = owner.Name
	ord.Photo = owner.Photo.String
}
