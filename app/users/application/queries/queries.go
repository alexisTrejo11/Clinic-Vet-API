package userQueries

import (
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type UserByIdQuery struct {
	ID string `json:"id"`
}

type UserByEmailQuery struct {
	Email string `json:"email"`
}

type UserSearchQuery struct {
	EmailLike       string              `json:"email_like"`
	PhoneLike       string              `json:"phone_like"`
	Role            userDomain.UserRole `json:"role"`
	JoindedAfter    string              `json:"joined_after"`
	JoinedBefore    string              `json:"joined_before"`
	LastLoginAtter  string              `json:"last_login_after"`
	LastLoginBefore string              `json:"last_login_before"`
	FirstNameLike   string              `json:"first_name_like"`
	LastNameLike    string              `json:"last_name_like"`
	Pagination      page.PageData       `json:"pagination"`
}

func (q *UserSearchQuery) ToMap() map[string]interface{} {
	searchMap := make(map[string]interface{})

	if q.EmailLike != "" {
		searchMap["email_like"] = q.EmailLike
	}
	if q.PhoneLike != "" {
		searchMap["phone_like"] = q.PhoneLike
	}
	if q.Role != "" {
		searchMap["role"] = q.Role
	}
	if q.JoindedAfter != "" {
		searchMap["joined_after"] = q.JoindedAfter
	}
	if q.JoinedBefore != "" {
		searchMap["joined_before"] = q.JoinedBefore
	}
	if q.LastLoginAtter != "" {
		searchMap["last_login_after"] = q.LastLoginAtter
	}
	if q.LastLoginBefore != "" {
		searchMap["last_login_before"] = q.LastLoginBefore
	}
	if q.FirstNameLike != "" {
		searchMap["first_name_like"] = q.FirstNameLike
	}
	if q.LastNameLike != "" {
		searchMap["last_name_like"] = q.LastNameLike
	}

	return searchMap
}

type UserListByRoleQuery struct {
	Role       userDomain.UserRole `json:"role"`
	Pagination page.PageData       `json:"pagination"`
}
