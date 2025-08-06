package userQueries

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
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

type ListUsersResponse struct {
	Users  []UserResponse `json:"users"`
	Total  int            `json:"total"`
	Count  int            `json:"count"`
	Offset int            `json:"offset"`
	Limit  int            `json:"limit"`
}
