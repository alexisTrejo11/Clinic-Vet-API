package query

import (
	"context"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/repository"
	p "clinic-vet-api/app/shared/page"
)

type FindUsersByRoleQuery struct {
	role       enum.UserRole
	pagination p.PageInput
}

func NewFindUsersByRoleQuery(role string, pagination p.PageInput) (*FindUsersByRoleQuery, error) {
	roleEnum, err := enum.ParseUserRole(role)
	if err != nil {
		return nil, err
	}

	return &FindUsersByRoleQuery{
		role:       roleEnum,
		pagination: pagination,
	}, nil
}

type FindUsersByRoleHandler struct {
	userRepository repository.UserRepository
}

func NewFindUsersByRoleHandler(userRepository repository.UserRepository) *FindUsersByRoleHandler {
	return &FindUsersByRoleHandler{
		userRepository: userRepository,
	}
}

func (h *FindUsersByRoleHandler) Handle(ctx context.Context, query FindUsersByRoleQuery) (p.Page[UserResult], error) {
	users, err := h.userRepository.FindByRole(ctx, query.role.String(), query.pagination)
	if err != nil {
		return p.Page[UserResult]{}, err
	}

	return toResultPage(users), nil
}
