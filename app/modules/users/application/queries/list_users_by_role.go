package userDomainQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type ListUsersByRoleQuery struct {
	Role       userDomain.UserRole `json:"role"`
	Pagination page.PageData       `json:"pagination"`
	Ctx        context.Context     `json:"-"`
}

type listUsersByRoleQueryHandler struct {
	userRepository userDomain.UserRepository
}

func NewListUsersByRoleQueryHandler(userRepository userDomain.UserRepository) *listUsersByRoleQueryHandler {
	return &listUsersByRoleQueryHandler{
		userRepository: userRepository,
	}
}

func (h *listUsersByRoleQueryHandler) Handle(query ListUsersByRoleQuery) (page.Page[[]UserResponse], error) {
	users, err := h.userRepository.ListByRole(query.Ctx, query.Role.String(), query.Pagination)
	if err != nil {
		return page.Page[[]UserResponse]{}, err
	}

	return toResponsePage(users), nil
}
