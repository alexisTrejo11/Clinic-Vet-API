package userDomainQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type UserSearchQuery struct {
	EmailLike       string          `json:"email_like"`
	PhoneLike       string          `json:"phone_like"`
	Role            enum.UserRole   `json:"role"`
	JoindedAfter    string          `json:"joined_after"`
	JoinedBefore    string          `json:"joined_before"`
	LastLoginAtter  string          `json:"last_login_after"`
	LastLoginBefore string          `json:"last_login_before"`
	FirstNameLike   string          `json:"first_name_like"`
	LastNameLike    string          `json:"last_name_like"`
	Pagination      page.PageData   `json:"pagination"`
	Ctx             context.Context `json:"-"`
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

type SearchUserHandler interface {
	Handle(query UserSearchQuery) (page.Page[[]UserResponse], error)
}

type searchUserHander struct {
	repository repository.UserRepository
}

func NewSearchUserHandler(repository repository.UserRepository) SearchUserHandler {
	return &searchUserHander{
		repository: repository,
	}
}

func (h *searchUserHander) Handle(query UserSearchQuery) (page.Page[[]UserResponse], error) {
	searchCriteria := query.ToMap()
	userPage, err := h.repository.Search(query.Ctx, searchCriteria, query.Pagination)
	if err != nil {
		return page.Page[[]UserResponse]{}, err
	}

	return toResponsePage(userPage), nil
}
