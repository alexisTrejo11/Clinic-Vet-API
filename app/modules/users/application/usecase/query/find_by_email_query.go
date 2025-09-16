package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type FindUserByEmailQuery struct {
	email          valueobject.Email
	includeProfile bool
	ctx            context.Context
}

func NewFindUserByEmailQuery(ctx context.Context, email string, includeProfile bool) (*FindUserByEmailQuery, error) {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &FindUserByEmailQuery{
		email:          emailVO,
		includeProfile: includeProfile,
		ctx:            ctx,
	}, nil
}

type FindUserByEmailHandler struct {
	userRepository repository.UserRepository
}

func NewFindUserByEmailHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResult] {
	return &FindUserByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *FindUserByEmailHandler) Handle(q cqrs.Query) (UserResult, error) {
	query, ok := q.(FindUserByEmailQuery)
	if !ok {
		return UserResult{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.FindByEmail(query.ctx, query.email.Value())
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
