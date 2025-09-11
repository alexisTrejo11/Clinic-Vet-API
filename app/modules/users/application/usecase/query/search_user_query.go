package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/specification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
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
