package query

import (
	"context"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
)

type FindUserByEmailQuery struct {
	email          valueobject.Email
	includeProfile bool
}

func NewFindUserByEmailQuery(email string, includeProfile bool) (*FindUserByEmailQuery, error) {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &FindUserByEmailQuery{
		email:          emailVO,
		includeProfile: includeProfile,
	}, nil
}

type FindUserByEmailHandler struct {
	userRepository repository.UserRepository
}

func NewFindUserByEmailHandler(userRepository repository.UserRepository) *FindUserByEmailHandler {
	return &FindUserByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *FindUserByEmailHandler) Handle(ctx context.Context, query FindUserByEmailQuery) (UserResult, error) {
	user, err := h.userRepository.FindByEmail(ctx, query.email.Value())
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
