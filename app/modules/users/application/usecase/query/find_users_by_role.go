package query

import (
	"context"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/repository"
	apperror "clinic-vet-api/app/shared/error/application"
	p "clinic-vet-api/app/shared/page"
)

type FindUsersByRoleQuery struct {
	role       enum.UserRole
	pagination p.PageInput
}

func NewFindUsersByRoleQuery(role string, pagination p.PageInput) *FindUsersByRoleQuery {
	roleEnum := enum.UserRole(role)

	return &FindUsersByRoleQuery{
		role:       roleEnum,
		pagination: pagination,
	}
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
	if !query.role.IsValid() {
		return p.Page[UserResult]{}, apperror.FieldValidationError("role", query.role.String(), "invalid role")
	}

	users, err := h.userRepository.FindByRole(ctx, query.role.String(), query.pagination)
	if err != nil {
		return p.Page[UserResult]{}, err
	}

	return toResultPage(users), nil
}
