package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListUsersByRoleQuery struct {
	Role       enum.UserRole   `json:"role"`
	Pagination page.PageInput  `json:"pagination"`
	Ctx        context.Context `json:"-"`
}

type ListUsersByRoleHandler struct {
	userRepository repository.UserRepository
}

func NewListUsersByRoleHandler(userRepository repository.UserRepository) cqrs.QueryHandler[page.Page[[]UserResponse]] {
	return &ListUsersByRoleHandler{
		userRepository: userRepository,
	}
}

func (h *ListUsersByRoleHandler) Handle(q cqrs.Query) (page.Page[[]UserResponse], error) {
	query, valid := q.(ListUsersByRoleQuery)
	if !valid {
		return page.Page[[]UserResponse]{}, errors.New("invalid query type")
	}

	users, err := h.userRepository.ListByRole(query.Ctx, query.Role.String(), query.Pagination)
	if err != nil {
		return page.Page[[]UserResponse]{}, err
	}

	return toResponsePage(users), nil
}
