package authDto

type RequestSignup struct {
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=8"`
	Birthday    string `json:"birthday"`
	Photo       string `json:"photo"`
	Genre       Genre  `json:"genre" validate:"required,oneof=male female other"`
}

type RequestLogin struct {
	IdentifierField string `json:"identifier_field" validate:"required"`
	Password        string `json:"password" validate:"required"`
}

type RequestLogout struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

type Genre string

const (
	Male   Genre = "male"
	Female Genre = "female"
	Other  Genre = "other"
)
