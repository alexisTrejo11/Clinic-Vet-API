package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/page"
)

type UserSearchQuery struct {
	ctx  context.Context
	spec specification.UserSpecification
}
type SearchUsersHandler struct {
	repository repository.UserRepository
}

func NewSearchUsersHandler(repository repository.UserRepository) cqrs.QueryHandler[page.Page[UserResult]] {
	return &SearchUsersHandler{
		repository: repository,
	}
}

func (h *SearchUsersHandler) Handle(q cqrs.Query) (page.Page[UserResult], error) {
	query, valid := q.(UserSearchQuery)
	if !valid {
		return page.Page[UserResult]{}, errors.New("invalid query type")
	}

	userPage, err := h.repository.FindSpecification(query.ctx, query.spec)
	if err != nil {
		return page.Page[UserResult]{}, err
	}

	return toResultPage(userPage), nil
}
