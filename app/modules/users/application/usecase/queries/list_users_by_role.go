package userDomainQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type ListUsersByRoleQuery struct {
	Role       enum.UserRole   `json:"role"`
	Pagination page.PageInput  `json:"pagination"`
	Ctx        context.Context `json:"-"`
}

type listUsersByRoleQueryHandler struct {
	userRepository repository.UserRepository
}

func NewListUsersByRoleQueryHandler(userRepository repository.UserRepository) *listUsersByRoleQueryHandler {
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
