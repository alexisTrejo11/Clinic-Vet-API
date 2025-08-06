package userQueries

type UserResponse struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	Status      string `json:"status"`
	JoinedAt    string `json:"joined_at"`
	LastLoginAt string `json:"last_login_at"`
}

type ProfileResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Gender      string
	Bio         string `json:"bio"`
	ProfilePic  string `json:"profile_pic"`
	Location    string `json:"location"`
	DateOfBirth string `json:"date_of_birth"`
	JoinedAt    string `json:"joined_at"`
}
