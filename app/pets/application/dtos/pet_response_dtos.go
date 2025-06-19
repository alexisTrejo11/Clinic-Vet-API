package dtos

type PetResponse struct {
	Id      int32  `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Photo   string `json:"photo" validate:"required"`
	Species string `json:"species" validate:"required"`
	Breed   string `json:"breed" validate:"required"`
	Age     int32  `json:"age" validate:"required"`
	OwnerID int32  `json:"owner_id" validate:"required"`
}
