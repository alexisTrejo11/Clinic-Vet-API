// Package query contains the data structures and conversion functions for user-related query responses.
package query

import (
	"time"

	"clinic-vet-api/app/core/domain/entity/user"
	"clinic-vet-api/app/shared/page"
)

type UserResult struct {
	ID          uint
	PhoneNumber string
	Email       string
	Role        string
	Status      string
	JoinedAt    time.Time
	LastLoginAt *time.Time
	CustomerID  *uint
	EmployeeID  *uint
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
		ID:          user.ID().Value(),
		Email:       user.Email().String(),
		PhoneNumber: user.PhoneNumber().String(),
		Role:        user.Role().String(),
		Status:      string(user.Status()),
		JoinedAt:    user.CreatedAt(),
		LastLoginAt: user.LastLoginAt(),
	}

	if user.CustomerID() != nil {
		val := user.CustomerID().Value()
		userResult.CustomerID = &val
	}

	if user.EmployeeID() != nil {
		val := user.EmployeeID().Value()
		userResult.EmployeeID = &val
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
