package dtos

type VetResponse struct {
	Id        int32  `json:"id"`
	UserId    *int32 `json:"user_id"`
	Name      string `json:"name"`
	Photo     string `json:"photo"`
	Email     string `json:"email"`
	Specialty string `json:"specialty"`
}
