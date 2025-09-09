package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
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
	Pagination      page.PageInput  `json:"pagination"`
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

type SearchUsersHandler struct {
	repository repository.UserRepository
}

func NewSearchUsersHandler(repository repository.UserRepository) cqrs.QueryHandler[page.Page[[]UserResponse]] {
	return &SearchUsersHandler{
		repository: repository,
	}
}

func (h *SearchUsersHandler) Handle(q cqrs.Query) (page.Page[[]UserResponse], error) {
	query, valid := q.(UserSearchQuery)
	if !valid {
		return page.Page[[]UserResponse]{}, errors.New("invalid query type")
	}

	searchCriteria := query.ToMap()
	userPage, err := h.repository.Search(query.Ctx, searchCriteria, query.Pagination)
	if err != nil {
		return page.Page[[]UserResponse]{}, err
	}

	return toResponsePage(userPage), nil
}
