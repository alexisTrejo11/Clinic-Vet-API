package query

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/specification"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
)

type FindUserBySpecificationQuery struct {
	ctx  context.Context
	spec specification.UserSpecification
}
type FindUserBySpecificationHandler struct {
	repository repository.UserRepository
}

func NewFindUserBySpecificationHandler(repository repository.UserRepository) FindUserBySpecificationHandler {
	return FindUserBySpecificationHandler{
		repository: repository,
	}
}

func (h FindUserBySpecificationHandler) Handle(ctx context.Context, query FindUserBySpecificationQuery) (page.Page[UserResult], error) {
	userPage, err := h.repository.FindSpecification(query.ctx, query.spec)
	if err != nil {
		return page.Page[UserResult]{}, err
	}

	return toResultPage(userPage), nil
}
