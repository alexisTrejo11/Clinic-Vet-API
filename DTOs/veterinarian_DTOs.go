package DTOs

type VetInsertDTO struct {
	Name      string `json:"name" validate:"required"`
	Photo     string `json:"photo"`
	Email     string `json:"email" validate:"required"`
	Specialty string `json:"specialty" validate:"required"`
}

type VetDTO struct {
	Id        int32  `json:"id" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Photo     string `json:"photo"`
	Email     string `json:"email" validate:"required"`
	Specialty string `json:"specialty" validate:"required"`
}

type VetUpdateDTO struct {
	Id        int32  `json:"id" validate:"required"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	Specialty string `json:"specialty"`
}
