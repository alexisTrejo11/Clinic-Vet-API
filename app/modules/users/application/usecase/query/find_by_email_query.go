package query

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
)

type FindUserByEmailQuery struct {
	email          valueobject.Email
	includeProfile bool
}

func NewFindUserByEmailQuery(email string, includeProfile bool) *FindUserByEmailQuery {
	emailVO := valueobject.NewEmailNoErr(email)
	return &FindUserByEmailQuery{
		email:          emailVO,
		includeProfile: includeProfile,
	}
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
