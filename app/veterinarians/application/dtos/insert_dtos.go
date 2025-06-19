package dtos

type VetCreate struct {
	Name      string `json:"name" validate:"required"`
	Photo     string `json:"photo"`
	Specialty string `json:"specialty" validate:"required"`
}

type VetUpdate struct {
	Id        int32  `json:"id" validate:"required"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	Specialty string `json:"specialty"`
}
