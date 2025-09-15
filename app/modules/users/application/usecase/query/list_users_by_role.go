package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
)

type ListUsersByRoleQuery struct {
	role       enum.UserRole
	pagination p.PageInput
	ctx        context.Context
}

func NewListUsersByRoleQuery(ctx context.Context, role string, pagination p.PageInput) (*ListUsersByRoleQuery, error) {
	roleEnum, err := enum.ParseUserRole(role)
	if err != nil {
		return nil, err
	}

	return &ListUsersByRoleQuery{
		role:       roleEnum,
		pagination: pagination,
		ctx:        ctx,
	}, nil
}

type ListUsersByRoleHandler struct {
	userRepository repository.UserRepository
}

func NewListUsersByRoleHandler(userRepository repository.UserRepository) cqrs.QueryHandler[p.Page[UserResult]] {
	return &ListUsersByRoleHandler{
		userRepository: userRepository,
	}
}

func (h *ListUsersByRoleHandler) Handle(q cqrs.Query) (p.Page[UserResult], error) {
	query, valid := q.(ListUsersByRoleQuery)
	if !valid {
		return p.Page[UserResult]{}, errors.New("invalid query type")
	}

	users, err := h.userRepository.FindByRole(query.ctx, query.role.String(), query.pagination)
	if err != nil {
		return p.Page[UserResult]{}, err
	}

	return toResultPage(users), nil
}
