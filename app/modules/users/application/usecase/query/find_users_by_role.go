package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
)

type FindUsersByRoleQuery struct {
	role       enum.UserRole
	pagination p.PageInput
	ctx        context.Context
}

func NewFindUsersByRoleQuery(ctx context.Context, role string, pagination p.PageInput) (*FindUsersByRoleQuery, error) {
	roleEnum, err := enum.ParseUserRole(role)
	if err != nil {
		return nil, err
	}

	return &FindUsersByRoleQuery{
		role:       roleEnum,
		pagination: pagination,
		ctx:        ctx,
	}, nil
}

type FindUsersByRoleHandler struct {
	userRepository repository.UserRepository
}

func NewFindUsersByRoleHandler(userRepository repository.UserRepository) cqrs.QueryHandler[p.Page[UserResult]] {
	return &FindUsersByRoleHandler{
		userRepository: userRepository,
	}
}

func (h *FindUsersByRoleHandler) Handle(q cqrs.Query) (p.Page[UserResult], error) {
	query, valid := q.(FindUsersByRoleQuery)
	if !valid {
		return p.Page[UserResult]{}, errors.New("invalid query type")
	}

	users, err := h.userRepository.FindByRole(query.ctx, query.role.String(), query.pagination)
	if err != nil {
		return p.Page[UserResult]{}, err
	}

	return toResultPage(users), nil
}
