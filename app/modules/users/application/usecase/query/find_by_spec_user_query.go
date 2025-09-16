package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/specification"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
	"clinic-vet-api/app/shared/page"
)

type UserFindBySpecificationQuery struct {
	ctx  context.Context
	spec specification.UserSpecification
}
type FindBySpecificationUsersHandler struct {
	repository repository.UserRepository
}

func NewFindBySpecificationUsersHandler(repository repository.UserRepository) cqrs.QueryHandler[page.Page[UserResult]] {
	return &FindBySpecificationUsersHandler{
		repository: repository,
	}
}

func (h *FindBySpecificationUsersHandler) Handle(q cqrs.Query) (page.Page[UserResult], error) {
	query, valid := q.(UserFindBySpecificationQuery)
	if !valid {
		return page.Page[UserResult]{}, errors.New("invalid query type")
	}

	userPage, err := h.repository.FindSpecification(query.ctx, query.spec)
	if err != nil {
		return page.Page[UserResult]{}, err
	}

	return toResultPage(userPage), nil
}
