package userQueries

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

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

func toResponse(user user.User) *UserResponse {
	return &UserResponse{
		Id:          user.Id().String(),
		Email:       user.Email().String(),
		PhoneNumber: user.PhoneNumber().String(),
		Role:        user.Role().String(),
		Status:      string(user.Status()),
		JoinedAt:    user.JoinedAt().Format("2006-01-02 15:04:05"),
		LastLoginAt: user.LastLoginAt().Format("2006-01-02 15:04:05"),
	}
}

func toResponsePage(userPage page.Page[[]user.User]) page.Page[[]UserResponse] {
	var userResponses []UserResponse

	users := userPage.Data
	for _, user := range users {
		userResponses = append(userResponses, *toResponse(user))
	}

	return *page.NewPage(userResponses, userPage.Metadata)
}
