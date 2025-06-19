package dtos

type CustomerSignup struct {
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=8"`
	Birthday    string `json:"birthday"`
	Photo       string `json:"photo"`
	Genre       Genre  `json:"genre" validate:"required,oneof=male female other"`
}

type LoginDTO struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password" validate:"required"`
}

type Genre string

const (
	Male   Genre = "male"
	Female Genre = "female"
	Other  Genre = "other"
)

type EmployeeSignup struct {
	Email          string `json:"email" validate:"required,email"`
	PhoneNumber    string `json:"phone_number" validate:"required,e164"`
	Password       string `json:"password" validate:"required,min=8"`
	VeterinarianId int32  `json:"veterinarian_id" validate:"required"`
}

type EmployeeLogin struct {
	VeterinarianId *int32 `json:"veterinarian_id"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Password       string `json:"password" validate:"required"`
}
