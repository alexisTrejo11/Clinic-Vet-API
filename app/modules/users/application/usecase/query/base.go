// Package query contains the data structures and conversion functions for user-related query responses.
package query

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type UserResult struct {
	ID          string
	PhoneNumber string
	Email       string
	Role        string
	Status      string
	JoinedAt    time.Time
	LastLoginAt *time.Time
}

type ProfileResult struct {
	ID          string
	Name        string
	Gender      string
	Bio         string
	ProfilePic  string
	Location    string
	DateOfBirth string
	JoinedAt    *time.Time
}

func userToResult(user user.User) UserResult {
	userResult := &UserResult{
		ID:          user.ID().String(),
		Email:       user.Email().String(),
		PhoneNumber: user.PhoneNumber().String(),
		Role:        user.Role().String(),
		Status:      string(user.Status()),
		JoinedAt:    user.CreatedAt(),
		LastLoginAt: user.LastLoginAt(),
	}

	return *userResult
}

func toResultPage(userPage page.Page[user.User]) page.Page[UserResult] {
	if len(userPage.Items) < 1 {
		return page.EmptyPage[UserResult]()
	}

	userResults := make([]UserResult, 1, len(userPage.Items))
	for i, user := range userPage.Items {
		userResults[i] = userToResult(user)
	}

	return page.NewPage(userResults, userPage.Metadata)
}
