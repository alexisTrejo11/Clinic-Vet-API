package query

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
)

type FindUserByEmailQuery struct {
	email valueobject.Email
}

func NewFindUserByEmailQuery(email string) FindUserByEmailQuery {
	return FindUserByEmailQuery{
		email: valueobject.NewEmailNoErr(email),
	}
}

type FindUserByEmailHandler struct {
	userRepository repository.UserRepository
}

func NewFindUserByEmailHandler(userRepository repository.UserRepository) FindUserByEmailHandler {
	return FindUserByEmailHandler{
		userRepository: userRepository,
	}
}

func (h FindUserByEmailHandler) Handle(ctx context.Context, query FindUserByEmailQuery) (UserResult, error) {
	user, err := h.userRepository.FindByEmail(ctx, query.email.Value())
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
