package dtos

type PetCreate struct {
	Name    string `json:"name" validate:"required"`
	Photo   string `json:"photo"`
	Species string `json:"species" validate:"required"`
	Breed   string `json:"breed" validate:"required"`
	Age     int32  `json:"age" validate:"required"`
}
type PetUpdate struct {
	Id      int32  `json:"id" validate:"required"`
	Name    string `json:"name"`
	Photo   string `json:"photo"`
	Species string `json:"species"`
	Breed   string `json:"breed"`
	Age     int32  `json:"age"`
}
