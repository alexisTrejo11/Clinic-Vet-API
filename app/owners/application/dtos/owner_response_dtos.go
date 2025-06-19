package dtos

type OwnerResponse struct {
	Id    int32    `json:"id"`
	Photo string   `json:"photo"`
	Name  string   `json:"name" validate:"required"`
	Phone string   `json:"phone"`
	Pets  []PetDTO `json:"pets"`
}
