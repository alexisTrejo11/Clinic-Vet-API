package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type FindUserByIDQuery struct {
	id             valueobject.UserID
	includeProfile bool
	ctx            context.Context
}

func NewFindUserByIDQuery(ctx context.Context, id uint, includeProfile bool) *FindUserByIDQuery {
	userID := valueobject.NewUserID(id)
	return &FindUserByIDQuery{
		id:             userID,
		includeProfile: includeProfile,
		ctx:            ctx,
	}
}

type FindUserByIDHandler struct {
	userRepository repository.UserRepository
}

func NewFindUserByIDHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResult] {
	return &FindUserByIDHandler{
		userRepository: userRepository,
	}
}

func (h *FindUserByIDHandler) Handle(q cqrs.Query) (UserResult, error) {
	query, valid := q.(FindUserByIDQuery)
	if !valid {
		return UserResult{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.FindByID(query.ctx, query.id)
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
